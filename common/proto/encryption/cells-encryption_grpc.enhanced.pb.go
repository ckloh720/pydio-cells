// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             (unknown)
// source: cells-encryption.proto

package encryption

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
	enhancedUserKeyStoreServers     = make(map[string]UserKeyStoreEnhancedServer)
	enhancedUserKeyStoreServersLock = sync.RWMutex{}
)

type idUserKeyStoreServer interface {
	ID() string
}
type UserKeyStoreEnhancedServer interface {
	UserKeyStoreServer
	AddFilter(func(context.Context, interface{}) bool)
	addHandler(UserKeyStoreServer)
	filter(context.Context) []UserKeyStoreServer
}
type UserKeyStoreEnhancedServerImpl struct {
	filters  []func(context.Context, interface{}) bool
	handlers []UserKeyStoreServer
}

func (m *UserKeyStoreEnhancedServerImpl) AddFilter(f func(context.Context, interface{}) bool) {
	m.filters = append(m.filters, f)
}
func (m *UserKeyStoreEnhancedServerImpl) addHandler(srv UserKeyStoreServer) {
	m.handlers = append(m.handlers, srv)
}
func (m *UserKeyStoreEnhancedServerImpl) filter(ctx context.Context) []UserKeyStoreServer {
	var ret []UserKeyStoreServer
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

func (m *UserKeyStoreEnhancedServerImpl) AddKey(ctx context.Context, r *AddKeyRequest) (*AddKeyResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.AddKey(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method AddKey not implemented")
}

func (m *UserKeyStoreEnhancedServerImpl) GetKey(ctx context.Context, r *GetKeyRequest) (*GetKeyResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.GetKey(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
}

func (m *UserKeyStoreEnhancedServerImpl) AdminListKeys(ctx context.Context, r *AdminListKeysRequest) (*AdminListKeysResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.AdminListKeys(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method AdminListKeys not implemented")
}

func (m *UserKeyStoreEnhancedServerImpl) AdminCreateKey(ctx context.Context, r *AdminCreateKeyRequest) (*AdminCreateKeyResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.AdminCreateKey(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateKey not implemented")
}

func (m *UserKeyStoreEnhancedServerImpl) AdminDeleteKey(ctx context.Context, r *AdminDeleteKeyRequest) (*AdminDeleteKeyResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.AdminDeleteKey(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteKey not implemented")
}

func (m *UserKeyStoreEnhancedServerImpl) AdminExportKey(ctx context.Context, r *AdminExportKeyRequest) (*AdminExportKeyResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.AdminExportKey(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method AdminExportKey not implemented")
}

func (m *UserKeyStoreEnhancedServerImpl) AdminImportKey(ctx context.Context, r *AdminImportKeyRequest) (*AdminImportKeyResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.AdminImportKey(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method AdminImportKey not implemented")
}
func (m *UserKeyStoreEnhancedServerImpl) mustEmbedUnimplementedUserKeyStoreServer() {}
func RegisterUserKeyStoreEnhancedServer(s grpc.ServiceRegistrar, srv UserKeyStoreServer) {
	idServer, ok := s.(idUserKeyStoreServer)
	if ok {
		enhancedUserKeyStoreServersLock.Lock()
		defer enhancedUserKeyStoreServersLock.Unlock()
		instance, ok := enhancedUserKeyStoreServers[idServer.ID()]
		if !ok {
			instance = &UserKeyStoreEnhancedServerImpl{}
			enhancedUserKeyStoreServers[idServer.ID()] = instance
		}
		instance.addHandler(srv)
		RegisterUserKeyStoreServer(s, instance)
	} else {
		RegisterUserKeyStoreServer(s, srv)
	}
}

var (
	enhancedNodeKeyManagerServers     = make(map[string]NodeKeyManagerEnhancedServer)
	enhancedNodeKeyManagerServersLock = sync.RWMutex{}
)

type idNodeKeyManagerServer interface {
	ID() string
}
type NodeKeyManagerEnhancedServer interface {
	NodeKeyManagerServer
	AddFilter(func(context.Context, interface{}) bool)
	addHandler(NodeKeyManagerServer)
	filter(context.Context) []NodeKeyManagerServer
}
type NodeKeyManagerEnhancedServerImpl struct {
	filters  []func(context.Context, interface{}) bool
	handlers []NodeKeyManagerServer
}

func (m *NodeKeyManagerEnhancedServerImpl) AddFilter(f func(context.Context, interface{}) bool) {
	m.filters = append(m.filters, f)
}
func (m *NodeKeyManagerEnhancedServerImpl) addHandler(srv NodeKeyManagerServer) {
	m.handlers = append(m.handlers, srv)
}
func (m *NodeKeyManagerEnhancedServerImpl) filter(ctx context.Context) []NodeKeyManagerServer {
	var ret []NodeKeyManagerServer
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

func (m *NodeKeyManagerEnhancedServerImpl) GetNodeInfo(ctx context.Context, r *GetNodeInfoRequest) (*GetNodeInfoResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.GetNodeInfo(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeInfo not implemented")
}

func (m *NodeKeyManagerEnhancedServerImpl) GetNodePlainSize(ctx context.Context, r *GetNodePlainSizeRequest) (*GetNodePlainSizeResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.GetNodePlainSize(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetNodePlainSize not implemented")
}

func (m *NodeKeyManagerEnhancedServerImpl) SetNodeInfo(s NodeKeyManager_SetNodeInfoServer) error {
	for _, handler := range m.filter(s.Context()) {
		return handler.SetNodeInfo(s)
	}
	return status.Errorf(codes.Unimplemented, "method SetNodeInfo not implemented")
}

func (m *NodeKeyManagerEnhancedServerImpl) CopyNodeInfo(ctx context.Context, r *CopyNodeInfoRequest) (*CopyNodeInfoResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.CopyNodeInfo(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method CopyNodeInfo not implemented")
}

func (m *NodeKeyManagerEnhancedServerImpl) DeleteNode(ctx context.Context, r *DeleteNodeRequest) (*DeleteNodeResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.DeleteNode(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}

func (m *NodeKeyManagerEnhancedServerImpl) DeleteNodeKey(ctx context.Context, r *DeleteNodeKeyRequest) (*DeleteNodeKeyResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.DeleteNodeKey(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodeKey not implemented")
}

func (m *NodeKeyManagerEnhancedServerImpl) DeleteNodeSharedKey(ctx context.Context, r *DeleteNodeSharedKeyRequest) (*DeleteNodeSharedKeyResponse, error) {
	for _, handler := range m.filter(ctx) {
		return handler.DeleteNodeSharedKey(ctx, r)
	}
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodeSharedKey not implemented")
}
func (m *NodeKeyManagerEnhancedServerImpl) mustEmbedUnimplementedNodeKeyManagerServer() {}
func RegisterNodeKeyManagerEnhancedServer(s grpc.ServiceRegistrar, srv NodeKeyManagerServer) {
	idServer, ok := s.(idNodeKeyManagerServer)
	if ok {
		enhancedNodeKeyManagerServersLock.Lock()
		defer enhancedNodeKeyManagerServersLock.Unlock()
		instance, ok := enhancedNodeKeyManagerServers[idServer.ID()]
		if !ok {
			instance = &NodeKeyManagerEnhancedServerImpl{}
			enhancedNodeKeyManagerServers[idServer.ID()] = instance
		}
		instance.addHandler(srv)
		RegisterNodeKeyManagerServer(s, instance)
	} else {
		RegisterNodeKeyManagerServer(s, srv)
	}
}
