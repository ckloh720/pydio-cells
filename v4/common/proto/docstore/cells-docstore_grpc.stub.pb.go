// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-client-stub v1.1.0
// - protoc             v3.18.1
// source: cells-docstore.proto

package docstore

import (
	context "context"
	fmt "fmt"
	stubs "github.com/pydio/cells/v4/common/server/stubs"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

type DocStoreStub struct {
	DocStoreServer
}

func (s *DocStoreStub) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	fmt.Println("Serving", method, args, reply, opts)
	var e error
	switch method {
	case "/docstore.DocStore/PutDocument":
		resp, er := s.DocStoreServer.PutDocument(ctx, args.(*PutDocumentRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/docstore.DocStore/GetDocument":
		resp, er := s.DocStoreServer.GetDocument(ctx, args.(*GetDocumentRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/docstore.DocStore/DeleteDocuments":
		resp, er := s.DocStoreServer.DeleteDocuments(ctx, args.(*DeleteDocumentsRequest))
		if er == nil {
			e = stubs.AssignToInterface(resp, reply)
		} else {
			e = er
		}
	case "/docstore.DocStore/CountDocuments":
		resp, er := s.DocStoreServer.CountDocuments(ctx, args.(*ListDocumentsRequest))
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
func (s *DocStoreStub) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	fmt.Println("Serving", method)
	switch method {
	case "/docstore.DocStore/ListDocuments":
		st := &DocStoreStub_ListDocumentsStreamer{}
		st.Init(ctx, func(i interface{}) error {
			return s.DocStoreServer.ListDocuments(i.(*ListDocumentsRequest), st)
		})
		return st, nil
	}
	return nil, fmt.Errorf(method + "  not implemented")
}

type DocStoreStub_ListDocumentsStreamer struct {
	stubs.ClientServerStreamerCore
}

func (s *DocStoreStub_ListDocumentsStreamer) Send(response *ListDocumentsResponse) error {
	s.RespChan <- response
	return nil
}
