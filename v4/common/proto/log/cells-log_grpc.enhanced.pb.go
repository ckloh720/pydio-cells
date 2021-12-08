// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.18.1
// source: cells-log.proto

package log

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
	enhancedLogRecorderServers = make(map[string]LogRecorderEnhancedServer)
)

type NamedLogRecorderServer interface {
	LogRecorderServer
	Name() string
}
type LogRecorderEnhancedServer map[string]NamedLogRecorderServer

// PutLog adds received log messages to the corresponding log repository.

func (m LogRecorderEnhancedServer) PutLog(s LogRecorder_PutLogServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method PutLog should have a context")
	}
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.PutLog(s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method PutLog not implemented")
}

// ListLogs performs a paginated search query in the log repository.

func (m LogRecorderEnhancedServer) ListLogs(r *ListLogRequest, s LogRecorder_ListLogsServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method ListLogs should have a context")
	}
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.ListLogs(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method ListLogs not implemented")
}

// DeleteLogs deletes logs based on a request (cannot be empty)

func (m LogRecorderEnhancedServer) DeleteLogs(ctx context.Context, r *ListLogRequest) (*DeleteLogsResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method DeleteLogs should have a context")
	}
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.DeleteLogs(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLogs not implemented")
}

// AggregatedLogs performs a query to retrieve log events of the given type, faceted by time range.

func (m LogRecorderEnhancedServer) AggregatedLogs(r *TimeRangeRequest, s LogRecorder_AggregatedLogsServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method AggregatedLogs should have a context")
	}
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.AggregatedLogs(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method AggregatedLogs not implemented")
}
func (m LogRecorderEnhancedServer) mustEmbedUnimplementedLogRecorderServer() {}
func RegisterLogRecorderEnhancedServer(s grpc.ServiceRegistrar, srv NamedLogRecorderServer) {
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedLogRecorderServers[addr]
	if !ok {
		m = LogRecorderEnhancedServer{}
		enhancedLogRecorderServers[addr] = m
		RegisterLogRecorderServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterLogRecorderEnhancedServer(s grpc.ServiceRegistrar, name string) {
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedLogRecorderServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}
