// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-client-stub v1.1.0
// - protoc             (unknown)
// source: cells-idm.proto

package idm

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

type RoleServiceStub struct {
	RoleServiceServer
}

func (s *RoleServiceStub) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	fmt.Println("Serving", method, args, reply, opts)
	var e error
	switch method {
	case "/idm.RoleService/CreateRole":
		resp, er := s.RoleServiceServer.CreateRole(ctx, args.(*CreateRoleRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.RoleService/DeleteRole":
		resp, er := s.RoleServiceServer.DeleteRole(ctx, args.(*DeleteRoleRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.RoleService/CountRole":
		resp, er := s.RoleServiceServer.CountRole(ctx, args.(*SearchRoleRequest))
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
func (s *RoleServiceStub) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Println("Serving", method)
	switch method {
	case "/idm.RoleService/SearchRole":
		st := &RoleServiceStub_SearchRoleStreamer{}
		st.Init(ctx, func(i interface{}) error {
			go func() {
				defer func() {
					close(st.RespChan)
				}()
				s.RoleServiceServer.SearchRole(i.(*SearchRoleRequest), st)
			}()
			<-time.After(100 * time.Millisecond)
			return nil
		})
		return st, nil
	case "/idm.RoleService/StreamRole":
		st := &RoleServiceStub_StreamRoleStreamer{}
		st.Init(ctx)
		go s.RoleServiceServer.StreamRole(st)
		return st, nil
	}
	return nil, fmt.Errorf(method + "  not implemented")
}

type RoleServiceStub_SearchRoleStreamer struct {
	stubs.ClientServerStreamerCore
}

func (s *RoleServiceStub_SearchRoleStreamer) Send(response *SearchRoleResponse) error {
	s.RespChan <- response
	return nil
}

type RoleServiceStub_StreamRoleStreamer struct {
	stubs.BidirServerStreamerCore
}

func (s *RoleServiceStub_StreamRoleStreamer) Recv() (*SearchRoleRequest, error) {
	if req, o := <-s.ReqChan; o {
		return req.(*SearchRoleRequest), nil
	} else {
		return nil, io.EOF
	}
}
func (s *RoleServiceStub_StreamRoleStreamer) Send(response *SearchRoleResponse) error {
	s.RespChan <- response
	return nil
}

type UserServiceStub struct {
	UserServiceServer
}

func (s *UserServiceStub) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	fmt.Println("Serving", method, args, reply, opts)
	var e error
	switch method {
	case "/idm.UserService/CreateUser":
		resp, er := s.UserServiceServer.CreateUser(ctx, args.(*CreateUserRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.UserService/DeleteUser":
		resp, er := s.UserServiceServer.DeleteUser(ctx, args.(*DeleteUserRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.UserService/BindUser":
		resp, er := s.UserServiceServer.BindUser(ctx, args.(*BindUserRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.UserService/CountUser":
		resp, er := s.UserServiceServer.CountUser(ctx, args.(*SearchUserRequest))
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
func (s *UserServiceStub) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Println("Serving", method)
	switch method {
	case "/idm.UserService/SearchUser":
		st := &UserServiceStub_SearchUserStreamer{}
		st.Init(ctx, func(i interface{}) error {
			go func() {
				defer func() {
					close(st.RespChan)
				}()
				s.UserServiceServer.SearchUser(i.(*SearchUserRequest), st)
			}()
			<-time.After(100 * time.Millisecond)
			return nil
		})
		return st, nil
	case "/idm.UserService/StreamUser":
		st := &UserServiceStub_StreamUserStreamer{}
		st.Init(ctx)
		go s.UserServiceServer.StreamUser(st)
		return st, nil
	}
	return nil, fmt.Errorf(method + "  not implemented")
}

type UserServiceStub_SearchUserStreamer struct {
	stubs.ClientServerStreamerCore
}

func (s *UserServiceStub_SearchUserStreamer) Send(response *SearchUserResponse) error {
	s.RespChan <- response
	return nil
}

type UserServiceStub_StreamUserStreamer struct {
	stubs.BidirServerStreamerCore
}

func (s *UserServiceStub_StreamUserStreamer) Recv() (*SearchUserRequest, error) {
	if req, o := <-s.ReqChan; o {
		return req.(*SearchUserRequest), nil
	} else {
		return nil, io.EOF
	}
}
func (s *UserServiceStub_StreamUserStreamer) Send(response *SearchUserResponse) error {
	s.RespChan <- response
	return nil
}

type WorkspaceServiceStub struct {
	WorkspaceServiceServer
}

func (s *WorkspaceServiceStub) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	fmt.Println("Serving", method, args, reply, opts)
	var e error
	switch method {
	case "/idm.WorkspaceService/CreateWorkspace":
		resp, er := s.WorkspaceServiceServer.CreateWorkspace(ctx, args.(*CreateWorkspaceRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.WorkspaceService/DeleteWorkspace":
		resp, er := s.WorkspaceServiceServer.DeleteWorkspace(ctx, args.(*DeleteWorkspaceRequest))
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
func (s *WorkspaceServiceStub) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Println("Serving", method)
	switch method {
	case "/idm.WorkspaceService/SearchWorkspace":
		st := &WorkspaceServiceStub_SearchWorkspaceStreamer{}
		st.Init(ctx, func(i interface{}) error {
			go func() {
				defer func() {
					close(st.RespChan)
				}()
				s.WorkspaceServiceServer.SearchWorkspace(i.(*SearchWorkspaceRequest), st)
			}()
			<-time.After(100 * time.Millisecond)
			return nil
		})
		return st, nil
	case "/idm.WorkspaceService/StreamWorkspace":
		st := &WorkspaceServiceStub_StreamWorkspaceStreamer{}
		st.Init(ctx)
		go s.WorkspaceServiceServer.StreamWorkspace(st)
		return st, nil
	}
	return nil, fmt.Errorf(method + "  not implemented")
}

type WorkspaceServiceStub_SearchWorkspaceStreamer struct {
	stubs.ClientServerStreamerCore
}

func (s *WorkspaceServiceStub_SearchWorkspaceStreamer) Send(response *SearchWorkspaceResponse) error {
	s.RespChan <- response
	return nil
}

type WorkspaceServiceStub_StreamWorkspaceStreamer struct {
	stubs.BidirServerStreamerCore
}

func (s *WorkspaceServiceStub_StreamWorkspaceStreamer) Recv() (*SearchWorkspaceRequest, error) {
	if req, o := <-s.ReqChan; o {
		return req.(*SearchWorkspaceRequest), nil
	} else {
		return nil, io.EOF
	}
}
func (s *WorkspaceServiceStub_StreamWorkspaceStreamer) Send(response *SearchWorkspaceResponse) error {
	s.RespChan <- response
	return nil
}

type ACLServiceStub struct {
	ACLServiceServer
}

func (s *ACLServiceStub) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	fmt.Println("Serving", method, args, reply, opts)
	var e error
	switch method {
	case "/idm.ACLService/CreateACL":
		resp, er := s.ACLServiceServer.CreateACL(ctx, args.(*CreateACLRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.ACLService/ExpireACL":
		resp, er := s.ACLServiceServer.ExpireACL(ctx, args.(*ExpireACLRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.ACLService/DeleteACL":
		resp, er := s.ACLServiceServer.DeleteACL(ctx, args.(*DeleteACLRequest))
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
func (s *ACLServiceStub) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Println("Serving", method)
	switch method {
	case "/idm.ACLService/SearchACL":
		st := &ACLServiceStub_SearchACLStreamer{}
		st.Init(ctx, func(i interface{}) error {
			go func() {
				defer func() {
					close(st.RespChan)
				}()
				s.ACLServiceServer.SearchACL(i.(*SearchACLRequest), st)
			}()
			<-time.After(100 * time.Millisecond)
			return nil
		})
		return st, nil
	case "/idm.ACLService/StreamACL":
		st := &ACLServiceStub_StreamACLStreamer{}
		st.Init(ctx)
		go s.ACLServiceServer.StreamACL(st)
		return st, nil
	}
	return nil, fmt.Errorf(method + "  not implemented")
}

type ACLServiceStub_SearchACLStreamer struct {
	stubs.ClientServerStreamerCore
}

func (s *ACLServiceStub_SearchACLStreamer) Send(response *SearchACLResponse) error {
	s.RespChan <- response
	return nil
}

type ACLServiceStub_StreamACLStreamer struct {
	stubs.BidirServerStreamerCore
}

func (s *ACLServiceStub_StreamACLStreamer) Recv() (*SearchACLRequest, error) {
	if req, o := <-s.ReqChan; o {
		return req.(*SearchACLRequest), nil
	} else {
		return nil, io.EOF
	}
}
func (s *ACLServiceStub_StreamACLStreamer) Send(response *SearchACLResponse) error {
	s.RespChan <- response
	return nil
}

type UserMetaServiceStub struct {
	UserMetaServiceServer
}

func (s *UserMetaServiceStub) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	fmt.Println("Serving", method, args, reply, opts)
	var e error
	switch method {
	case "/idm.UserMetaService/UpdateUserMeta":
		resp, er := s.UserMetaServiceServer.UpdateUserMeta(ctx, args.(*UpdateUserMetaRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.UserMetaService/UpdateUserMetaNamespace":
		resp, er := s.UserMetaServiceServer.UpdateUserMetaNamespace(ctx, args.(*UpdateUserMetaNamespaceRequest))
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
func (s *UserMetaServiceStub) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Println("Serving", method)
	switch method {
	case "/idm.UserMetaService/SearchUserMeta":
		st := &UserMetaServiceStub_SearchUserMetaStreamer{}
		st.Init(ctx, func(i interface{}) error {
			go func() {
				defer func() {
					close(st.RespChan)
				}()
				s.UserMetaServiceServer.SearchUserMeta(i.(*SearchUserMetaRequest), st)
			}()
			<-time.After(100 * time.Millisecond)
			return nil
		})
		return st, nil
	case "/idm.UserMetaService/ListUserMetaNamespace":
		st := &UserMetaServiceStub_ListUserMetaNamespaceStreamer{}
		st.Init(ctx, func(i interface{}) error {
			go func() {
				defer func() {
					close(st.RespChan)
				}()
				s.UserMetaServiceServer.ListUserMetaNamespace(i.(*ListUserMetaNamespaceRequest), st)
			}()
			<-time.After(100 * time.Millisecond)
			return nil
		})
		return st, nil
	}
	return nil, fmt.Errorf(method + "  not implemented")
}

type UserMetaServiceStub_SearchUserMetaStreamer struct {
	stubs.ClientServerStreamerCore
}

func (s *UserMetaServiceStub_SearchUserMetaStreamer) Send(response *SearchUserMetaResponse) error {
	s.RespChan <- response
	return nil
}

type UserMetaServiceStub_ListUserMetaNamespaceStreamer struct {
	stubs.ClientServerStreamerCore
}

func (s *UserMetaServiceStub_ListUserMetaNamespaceStreamer) Send(response *ListUserMetaNamespaceResponse) error {
	s.RespChan <- response
	return nil
}

type PolicyEngineServiceStub struct {
	PolicyEngineServiceServer
}

func (s *PolicyEngineServiceStub) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	fmt.Println("Serving", method, args, reply, opts)
	var e error
	switch method {
	case "/idm.PolicyEngineService/IsAllowed":
		resp, er := s.PolicyEngineServiceServer.IsAllowed(ctx, args.(*PolicyEngineRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.PolicyEngineService/StorePolicyGroup":
		resp, er := s.PolicyEngineServiceServer.StorePolicyGroup(ctx, args.(*StorePolicyGroupRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.PolicyEngineService/ListPolicyGroups":
		resp, er := s.PolicyEngineServiceServer.ListPolicyGroups(ctx, args.(*ListPolicyGroupsRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/idm.PolicyEngineService/DeletePolicyGroup":
		resp, er := s.PolicyEngineServiceServer.DeletePolicyGroup(ctx, args.(*DeletePolicyGroupRequest))
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
func (s *PolicyEngineServiceStub) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Println("Serving", method)
	switch method {
	}
	return nil, fmt.Errorf(method + "  not implemented")
}
