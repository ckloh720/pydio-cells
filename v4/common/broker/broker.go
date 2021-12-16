/*
 * Copyright (c) 2019-2021. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package broker

import (
	"context"
	"fmt"
	"gocloud.dev/pubsub"

	"github.com/pydio/cells/v4/common/service/context/metadata"
	"google.golang.org/protobuf/proto"

	_ "gocloud.dev/pubsub/mempubsub"
)

var (
	std = NewBroker("memory")
)

type Broker interface {
	Publish(context.Context, string, proto.Message, ...PublishOption) error
	Subscribe(context.Context, string, SubscriberHandler, ...SubscribeOption) (UnSubscriber, error)
}

type UnSubscriber func() error

type SubscriberHandler func(Message) error

// NewBroker wraps a standard broker but prevents it from disconnecting while there still is a service running
func NewBroker(s string, opts ...Option) Broker {
	return &broker{
		publishOpener: func(ctx context.Context, topic string) (*pubsub.Topic, error) {
			return pubsub.OpenTopic(ctx, "mem://" + topic)
		},
		subscribeOpener: func(ctx context.Context, topic string) (*pubsub.Subscription, error) {
			return pubsub.OpenSubscription(ctx, "mem://" + topic)
		},
		publishers: make(map[string]*pubsub.Topic),
		Options: newOptions(opts...),
	}
}

// Publish sends a message to standard broker. For the moment, forward message to client.Publish
func Publish(ctx context.Context, topic string, message proto.Message, opts ...PublishOption) error {
	return std.Publish(ctx, topic, message, opts...)
}

// MustPublish publishes a message ignoring the error
func MustPublish(ctx context.Context, topic string, message proto.Message, opts ...PublishOption) {
	err := Publish(ctx, topic, message)
	if err != nil {
		fmt.Printf("[Message Publication Error] Topic: %s, Error: %v\n", topic, err)
	}
}

func SubscribeCancellable(ctx context.Context, topic string, handler SubscriberHandler, opts ...SubscribeOption) error {
	unsub, e := std.Subscribe(ctx, topic, handler, opts...)
	if e != nil {
		return e
	}
	go func() {
		<-ctx.Done()
		_ = unsub()
	}()

	return nil
}

func Subscribe(ctx context.Context, topic string, handler SubscriberHandler, opts ...SubscribeOption) (UnSubscriber, error) {
	return std.Subscribe(ctx,topic, handler, opts...)
}

type broker struct {
	publishOpener TopicOpener
	subscribeOpener SubscribeOpener
	publishers map[string]*pubsub.Topic
	Options
}

type TopicOpener func(context.Context, string) (*pubsub.Topic, error)
type SubscribeOpener func(context.Context, string) (*pubsub.Subscription, error)

func (b *broker) openTopic(ctx context.Context, topic string) (*pubsub.Topic, error) {
	publisher, ok := b.publishers[topic]
	if !ok {
		var err error
		publisher, err = b.publishOpener(ctx, topic)
		if err != nil{
			return nil, err
		}
		b.publishers[topic] = publisher
	}

	return publisher, nil
}

// Publish sends a message to standard broker. For the moment, forward message to client.Publish
func (b *broker) Publish(ctx context.Context, topic string, message proto.Message, opts ...PublishOption) error {
	body, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	header := make(map[string]string)
	if hh, ok := metadata.FromContext(ctx); ok {
		for k, v := range hh {
			header[k] = v
		}
	}

	publisher, err := b.openTopic(ctx, topic)
	if err != nil {
		return err
	}

	if err := publisher.Send(ctx, &pubsub.Message{
		Body: body,
		Metadata: header,
	}); err != nil {
		return err
	}

	return nil
}

func (b *broker) Subscribe(ctx context.Context, topic string, handler SubscriberHandler, opts ...SubscribeOption) (UnSubscriber, error) {
	so := &SubscribeOptions{}
	for _, o := range opts {
		o(so)
	}
	var mopts []SubscribeOption
	if so.Context != nil {
		mopts = append(mopts, SubscribeContext(so.Context))
	}
	if so.Queue != "" {
		mopts = append(mopts, Queue(so.Queue))
	}
	if so.ErrorHandler != nil {
		mopts = append(mopts, HandleError(func(err error) {
			so.ErrorHandler(err)
		}))
	}

	// Making sure topic is opened
	_, err := b.openTopic(ctx, topic)
	if err != nil {
		return nil, err
	}

	sub, err := b.subscribeOpener(ctx, topic)
	if err != nil {
		return nil, err
	}

	go func() {
		defer sub.Shutdown(ctx)

		for {
			msg, err := sub.Receive(ctx)
			if err != nil {
				break
			}

			if err := handler(&message{
				header: msg.Metadata,
				body: msg.Body,
			}); err != nil {
				fmt.Println("Could not handle message ? ", msg)
			}
		}
	}()

	return func() error {
		return sub.Shutdown(ctx)
	}, nil
}