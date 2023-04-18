// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-client-stub v1.1.0
// - protoc             (unknown)
// source: cells-registry.proto

package registry

import (
	context "context"
	fmt "fmt"
	stubs "github.com/pydio/cells/v4/common/server/stubs"
	grpc "google.golang.org/grpc"
	io "io"
	time "time"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

type RegistryStub struct {
	RegistryServer
}

func (s *RegistryStub) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	fmt.Println("Serving", method, args, reply, opts)
	var e error
	switch method {
	case "/registry.Registry/Start":
		resp, er := s.RegistryServer.Start(ctx, args.(*Item))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/registry.Registry/Stop":
		resp, er := s.RegistryServer.Stop(ctx, args.(*Item))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/registry.Registry/Get":
		resp, er := s.RegistryServer.Get(ctx, args.(*GetRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/registry.Registry/Register":
		resp, er := s.RegistryServer.Register(ctx, args.(*Item))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/registry.Registry/Deregister":
		resp, er := s.RegistryServer.Deregister(ctx, args.(*Item))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/registry.Registry/List":
		resp, er := s.RegistryServer.List(ctx, args.(*ListRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	default:
		e = fmt.Errorf(method + " not implemented")
	}
	return e
}
func (s *RegistryStub) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Println("Serving", method)
	switch method {
	case "/registry.Registry/Session":
		st := &RegistryStub_SessionStreamer{}
		st.Init(ctx)
		go s.RegistryServer.Session(st)
		return st, nil
	case "/registry.Registry/Watch":
		st := &RegistryStub_WatchStreamer{}
		st.Init(ctx, func(i interface{}) error {
			go func() {
				defer func() {
					close(st.RespChan)
				}()
				s.RegistryServer.Watch(i.(*WatchRequest), st)
			}()
			<-time.After(100 * time.Millisecond)
			return nil
		})
		return st, nil
	case "/registry.Registry/NewLocker":
		st := &RegistryStub_NewLockerStreamer{}
		st.Init(ctx)
		go s.RegistryServer.NewLocker(st)
		return st, nil
	}
	return nil, fmt.Errorf(method + "  not implemented")
}

type RegistryStub_SessionStreamer struct {
	stubs.BidirServerStreamerCore
}

func (s *RegistryStub_SessionStreamer) Recv() (*SessionRequest, error) {
	if req, o := <-s.ReqChan; o {
		return req.(*SessionRequest), nil
	} else {
		return nil, io.EOF
	}
}
func (s *RegistryStub_SessionStreamer) Send(response *EmptyResponse) error {
	s.RespChan <- response
	return nil
}

type RegistryStub_WatchStreamer struct {
	stubs.ClientServerStreamerCore
}

func (s *RegistryStub_WatchStreamer) Send(response *Result) error {
	s.RespChan <- response
	return nil
}

type RegistryStub_NewLockerStreamer struct {
	stubs.BidirServerStreamerCore
}

func (s *RegistryStub_NewLockerStreamer) Recv() (*NewLockerRequest, error) {
	if req, o := <-s.ReqChan; o {
		return req.(*NewLockerRequest), nil
	} else {
		return nil, io.EOF
	}
}
func (s *RegistryStub_NewLockerStreamer) Send(response *NewLockerResponse) error {
	s.RespChan <- response
	return nil
}
