package service

import (
	"context"
	"fmt"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/pydio/cells/v4/common/proto/tree"
	"google.golang.org/protobuf/proto"

	"github.com/pydio/cells/v4/common"
	"github.com/pydio/cells/v4/common/client/grpc"
	"github.com/pydio/cells/v4/common/server/stubs/discoverytest"
	"gocloud.dev/pubsub"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	grpc.RegisterMock(common.ServiceBroker, discoverytest.NewBrokerService())
}

func TestServiceBroker(t *testing.T) {
	Convey("Test Service Broker", t, func() {
		numMessagesToSend := 1000
		numMessagesReceived := 0

		subscription, _ := NewSubscription("test")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		go func() {
			defer cancel()

			defer subscription.Shutdown(context.Background())

			for {
				msg, err := subscription.Receive(context.Background())
				if err == io.EOF {
					return
				}

				numMessagesReceived++

				msg.Ack()

				ev := &tree.NodeChangeEvent{}
				if err := proto.Unmarshal(msg.Body, ev); err != nil {
					return
				}
			}
		}()

		topic, _ := NewTopic("test")
		defer topic.Shutdown(ctx)

		msg := &tree.NodeChangeEvent{Source: &tree.Node{Path: "source"}, Target: &tree.Node{Path: "target"}}
		b, err := proto.Marshal(msg)
		if err != nil {
			fmt.Println("Error ", err)
			return
		}

		for i := 0; i < numMessagesToSend; i++ {
			go topic.Send(context.Background(), &pubsub.Message{
				Body: b,
			})
		}

		select {
		case <-ctx.Done():
		}

		So(numMessagesReceived, ShouldEqual, numMessagesToSend)

	})
}

func TestConcurrentReceivesGetAllTheMessages(t *testing.T) {
	howManyToSend := int(1e3)
	ctx, cancel := context.WithCancel(context.Background())

	// wg is used to wait until all messages are received.
	var wg sync.WaitGroup
	wg.Add(howManyToSend)

	// Make a subscription.
	s, _ := NewSubscription("test")
	defer s.Shutdown(ctx)

	// Start 10 goroutines to receive from it.
	var mu sync.Mutex
	receivedMsgs := make(map[string]bool)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				m, err := s.Receive(ctx)
				if err != nil {
					// Permanent error; ctx cancelled or subscription closed is
					// expected once we've received all the messages.
					mu.Lock()
					n := len(receivedMsgs)
					mu.Unlock()
					if n != howManyToSend {
						t.Errorf("Worker's Receive failed before all messages were received (%d)", n)
					}
					return
				}
				m.Ack()
				mu.Lock()
				receivedMsgs[string(m.Body)] = true
				mu.Unlock()
				wg.Done()
			}
		}()
	}

	// Send messages. Each message has a unique body used as a key to receivedMsgs.
	topic, _ := NewTopic("test")
	defer topic.Shutdown(ctx)
	for i := 0; i < howManyToSend; i++ {
		key := fmt.Sprintf("message #%d", i)
		m := &pubsub.Message{Body: []byte(key)}
		if err := topic.Send(ctx, m); err != nil {
			t.Fatal(err)
		}
	}

	// Wait for the goroutines to receive all of the messages, then cancel the
	// ctx so they all exit.
	wg.Wait()
	defer cancel()

	// Check that all the messages were received.
	for i := 0; i < howManyToSend; i++ {
		key := fmt.Sprintf("message #%d", i)
		if !receivedMsgs[key] {
			t.Errorf("message %q was not received", key)
		}
	}
}
