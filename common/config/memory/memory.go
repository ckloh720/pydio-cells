package memory

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/r3labs/diff/v3"

	"github.com/pydio/cells/v4/common/config"
	"github.com/pydio/cells/v4/common/utils/configx"
)

var (
	scheme           = "mem"
	errClosedChannel = errors.New("channel is closed")
)

type URLOpener struct{}

func init() {
	o := &URLOpener{}
	config.DefaultURLMux().Register(scheme, o)
}

func (o *URLOpener) OpenURL(ctx context.Context, u *url.URL) (config.Store, error) {

	var opts []configx.Option

	encode := u.Query().Get("encode")
	switch encode {
	case "string":
		opts = append(opts, configx.WithString())
	case "yaml":
		opts = append(opts, configx.WithYAML())
	case "json":
		opts = append(opts, configx.WithJSON())
	default:
		opts = append(opts, configx.WithJSON())
	}

	store := New(opts...)

	return store, nil
}

type memory struct {
	v    configx.Values
	snap configx.Values

	opts      []configx.Option
	receivers []*receiver

	reset chan bool
	timer *time.Timer

	*sync.RWMutex
}

func New(opts ...configx.Option) config.Store {
	m := &memory{
		v:       configx.New(opts...),
		opts:    opts,
		RWMutex: &sync.RWMutex{},
		reset:   make(chan bool),
		timer:   time.NewTimer(2 * time.Second),
	}

	// TODO - change this
	// We're considering that non comparable items cannot be watched
	// fmt.Println(len(opts))
	// if len(opts) > 0 {
	go m.flush()
	//}

	return m
}

func (m *memory) flush() {
	for {
		select {
		case <-m.reset:
			m.timer.Reset(2 * time.Second)
		case <-m.timer.C:
			patch, err := diff.Diff(m.snap, m.v)
			if err != nil {
				continue
			}

			for _, op := range patch {
				var updated []*receiver

				for _, r := range m.receivers {
					if err := r.call(op); err == nil {
						updated = append(updated, r)
					}
				}

				m.receivers = updated
			}

			snap := configx.New(m.opts...)
			if err := snap.Set(m.v.Val().Get()); err != nil {
				continue
			}

			m.snap = snap
		}
	}
}

func (m *memory) update() {
	select {
	case m.reset <- true:
	default:
	}
}

func (m *memory) Get() configx.Value {
	return m.v
}

func (m *memory) Set(data interface{}) error {
	if err := m.v.Set(data); err != nil {
		return err
	}

	m.update()

	return nil
}

func (m *memory) Val(path ...string) configx.Values {
	return &values{Values: m.v.Val(path...), m: m}
}

func (m *memory) Del() error {
	return fmt.Errorf("not implemented")
}

func (m *memory) Save(string, string) error {
	// do nothing
	return nil
}

func (m *memory) Watch(opts ...configx.WatchOption) (configx.Receiver, error) {
	o := &configx.WatchOptions{}
	for _, opt := range opts {
		opt(o)
	}

	r := &receiver{
		closed:      false,
		ch:          make(chan diff.Change),
		path:       o.Path,
		m:           m,
		timer:       time.NewTimer(2 * time.Second),
		changesOnly: o.ChangesOnly,
	}

	m.receivers = append(m.receivers, r)

	return r, nil
}

type receiver struct {
	closed bool
	ch     chan diff.Change

	path       []string
	changesOnly bool

	timer *time.Timer

	m *memory
}

func (r *receiver) call(op diff.Change) error {
	if r.closed {
		return errClosedChannel
	}

	if len(r.path) == 0 {
		r.ch <- op
	}

	if strings.HasPrefix(strings.Join(op.Path, "/"), strings.Join(r.path, "/")) {
		r.ch <- op
	}
	return nil
}

func (r *receiver) Next() (interface{}, error) {
	changes := []diff.Change{}

	for {
		select {
		case op := <-r.ch:
			changes = append(changes, op)

			r.timer.Reset(2 * time.Second)
		case <-r.timer.C:
			c := configx.New()

			if r.changesOnly {
				for _, op := range changes {
					if err := c.Val(op.Type).Val(op.Path[1:]...).Set(op.To); err != nil {
						return nil, err
					}
				}

				return c, nil
			}

			for _, op := range changes {
				if err := c.Val(op.Path[1:]...).Set(op.To); err != nil {
					return nil, err
				}
			}

			return c, nil
		}
	}

	return r.Next()
}

func (r *receiver) Stop() {
	r.closed = true
	close(r.ch)
}

type values struct {
	configx.Values

	m *memory
}

func (v *values) Set(data interface{}) error {
	if err := v.Values.Set(data); err != nil {
		return err
	}

	v.m.update()

	return nil
}

func (v *values) Del() error {
	if err := v.Values.Del(); err != nil {
		return err
	}

	v.m.update()

	return nil
}

func (v *values) Val(path ...string) configx.Values {
	return &values{Values: v.Values.Val(path...), m: v.m}
}
