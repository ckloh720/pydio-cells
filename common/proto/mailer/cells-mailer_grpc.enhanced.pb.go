// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             (unknown)
// source: cells-mailer.proto

package mailer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	sync "sync"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var (
	enhancedMailerServiceServers     = make(map[string]MailerServiceEnhancedServer)
	enhancedMailerServiceServersLock = sync.RWMutex{}
)

type idMailerServiceServer interface {
	ID() string
}
type MailerServiceEnhancedServer interface {
	MailerServiceServer
	AddFilter(func(context.Context, interface{}) bool)
	addHandler(MailerServiceServer)
	filter(context.Context) []MailerServiceServer
}
type MailerServiceEnhancedServerImpl struct {
	filters  []func(context.Context, interface{}) bool
	handlers []MailerServiceServer
}

func (m *MailerServiceEnhancedServerImpl) AddFilter(f func(context.Context, interface{}) bool) {
	m.filters = append(m.filters, f)
}
func (m *MailerServiceEnhancedServerImpl) addHandler(srv MailerServiceServer) {
	m.handlers = append(m.handlers, srv)
}
func (m *MailerServiceEnhancedServerImpl) filter(ctx context.Context) []MailerServiceServer {
	var ret []MailerServiceServer
	for _, i := range m.handlers {
		valid := true
		for _, filter := range m.filters {
			if !filter(ctx, i) {
				valid = false
				break
			}
			if valid {
				ret = append(ret, i)
			}
		}
	}
	return ret
}

func (m *MailerServiceEnhancedServerImpl) SendMail(ctx context.Context, r *SendMailRequest) (*SendMailResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.SendMail(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}

func (m *MailerServiceEnhancedServerImpl) ConsumeQueue(ctx context.Context, r *ConsumeQueueRequest) (*ConsumeQueueResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.ConsumeQueue(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method ConsumeQueue not implemented")
}
func (m *MailerServiceEnhancedServerImpl) mustEmbedUnimplementedMailerServiceServer() {}
func RegisterMailerServiceEnhancedServer(s grpc.ServiceRegistrar, srv MailerServiceServer) {
	idServer, ok := s.(idMailerServiceServer)
	if ok {
		enhancedMailerServiceServersLock.Lock()
		defer enhancedMailerServiceServersLock.Unlock()
		instance, ok := enhancedMailerServiceServers[idServer.ID()]
		if !ok {
			instance = &MailerServiceEnhancedServerImpl{}
			enhancedMailerServiceServers[idServer.ID()] = instance
		}
		instance.addHandler(srv)
		RegisterMailerServiceServer(s, instance)
	} else {
		RegisterMailerServiceServer(s, srv)
	}
}
