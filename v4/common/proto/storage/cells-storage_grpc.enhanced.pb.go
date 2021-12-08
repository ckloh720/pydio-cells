// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.18.1
// source: cells-storage.proto

package storage

import (
	context "context"
	fmt "fmt"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var (
	enhancedStorageEndpointServers = make(map[string]StorageEndpointEnhancedServer)
)

type NamedStorageEndpointServer interface {
	StorageEndpointServer
	Name() string
}
type StorageEndpointEnhancedServer map[string]NamedStorageEndpointServer

func (m StorageEndpointEnhancedServer) Propose(ctx context.Context, r *ProposeRequest) (*ProposeResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method Propose should have a context")
	}
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.Propose(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method Propose not implemented")
}

func (m StorageEndpointEnhancedServer) Lookup(ctx context.Context, r *LookupRequest) (*LookupResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method Lookup should have a context")
	}
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.Lookup(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}
func (m StorageEndpointEnhancedServer) mustEmbedUnimplementedStorageEndpointServer() {}
func RegisterStorageEndpointEnhancedServer(s grpc.ServiceRegistrar, srv NamedStorageEndpointServer) {
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedStorageEndpointServers[addr]
	if !ok {
		m = StorageEndpointEnhancedServer{}
		enhancedStorageEndpointServers[addr] = m
		RegisterStorageEndpointServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterStorageEndpointEnhancedServer(s grpc.ServiceRegistrar, name string) {
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedStorageEndpointServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}
