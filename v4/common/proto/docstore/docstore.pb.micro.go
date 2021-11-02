// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: docstore.proto

package docstore

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DocStore service

type DocStoreService interface {
	PutDocument(ctx context.Context, in *PutDocumentRequest, opts ...client.CallOption) (*PutDocumentResponse, error)
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...client.CallOption) (*GetDocumentResponse, error)
	DeleteDocuments(ctx context.Context, in *DeleteDocumentsRequest, opts ...client.CallOption) (*DeleteDocumentsResponse, error)
	CountDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...client.CallOption) (*CountDocumentsResponse, error)
	ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...client.CallOption) (DocStore_ListDocumentsService, error)
}

type docStoreService struct {
	c    client.Client
	name string
}

func NewDocStoreService(name string, c client.Client) DocStoreService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "docstore"
	}
	return &docStoreService{
		c:    c,
		name: name,
	}
}

func (c *docStoreService) PutDocument(ctx context.Context, in *PutDocumentRequest, opts ...client.CallOption) (*PutDocumentResponse, error) {
	req := c.c.NewRequest(c.name, "DocStore.PutDocument", in)
	out := new(PutDocumentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *docStoreService) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...client.CallOption) (*GetDocumentResponse, error) {
	req := c.c.NewRequest(c.name, "DocStore.GetDocument", in)
	out := new(GetDocumentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *docStoreService) DeleteDocuments(ctx context.Context, in *DeleteDocumentsRequest, opts ...client.CallOption) (*DeleteDocumentsResponse, error) {
	req := c.c.NewRequest(c.name, "DocStore.DeleteDocuments", in)
	out := new(DeleteDocumentsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *docStoreService) CountDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...client.CallOption) (*CountDocumentsResponse, error) {
	req := c.c.NewRequest(c.name, "DocStore.CountDocuments", in)
	out := new(CountDocumentsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *docStoreService) ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...client.CallOption) (DocStore_ListDocumentsService, error) {
	req := c.c.NewRequest(c.name, "DocStore.ListDocuments", &ListDocumentsRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &docStoreServiceListDocuments{stream}, nil
}

type DocStore_ListDocumentsService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ListDocumentsResponse, error)
}

type docStoreServiceListDocuments struct {
	stream client.Stream
}

func (x *docStoreServiceListDocuments) Close() error {
	return x.stream.Close()
}

func (x *docStoreServiceListDocuments) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *docStoreServiceListDocuments) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *docStoreServiceListDocuments) Recv() (*ListDocumentsResponse, error) {
	m := new(ListDocumentsResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DocStore service

type DocStoreHandler interface {
	PutDocument(context.Context, *PutDocumentRequest, *PutDocumentResponse) error
	GetDocument(context.Context, *GetDocumentRequest, *GetDocumentResponse) error
	DeleteDocuments(context.Context, *DeleteDocumentsRequest, *DeleteDocumentsResponse) error
	CountDocuments(context.Context, *ListDocumentsRequest, *CountDocumentsResponse) error
	ListDocuments(context.Context, *ListDocumentsRequest, DocStore_ListDocumentsStream) error
}

func RegisterDocStoreHandler(s server.Server, hdlr DocStoreHandler, opts ...server.HandlerOption) error {
	type docStore interface {
		PutDocument(ctx context.Context, in *PutDocumentRequest, out *PutDocumentResponse) error
		GetDocument(ctx context.Context, in *GetDocumentRequest, out *GetDocumentResponse) error
		DeleteDocuments(ctx context.Context, in *DeleteDocumentsRequest, out *DeleteDocumentsResponse) error
		CountDocuments(ctx context.Context, in *ListDocumentsRequest, out *CountDocumentsResponse) error
		ListDocuments(ctx context.Context, stream server.Stream) error
	}
	type DocStore struct {
		docStore
	}
	h := &docStoreHandler{hdlr}
	return s.Handle(s.NewHandler(&DocStore{h}, opts...))
}

type docStoreHandler struct {
	DocStoreHandler
}

func (h *docStoreHandler) PutDocument(ctx context.Context, in *PutDocumentRequest, out *PutDocumentResponse) error {
	return h.DocStoreHandler.PutDocument(ctx, in, out)
}

func (h *docStoreHandler) GetDocument(ctx context.Context, in *GetDocumentRequest, out *GetDocumentResponse) error {
	return h.DocStoreHandler.GetDocument(ctx, in, out)
}

func (h *docStoreHandler) DeleteDocuments(ctx context.Context, in *DeleteDocumentsRequest, out *DeleteDocumentsResponse) error {
	return h.DocStoreHandler.DeleteDocuments(ctx, in, out)
}

func (h *docStoreHandler) CountDocuments(ctx context.Context, in *ListDocumentsRequest, out *CountDocumentsResponse) error {
	return h.DocStoreHandler.CountDocuments(ctx, in, out)
}

func (h *docStoreHandler) ListDocuments(ctx context.Context, stream server.Stream) error {
	m := new(ListDocumentsRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.DocStoreHandler.ListDocuments(ctx, m, &docStoreListDocumentsStream{stream})
}

type DocStore_ListDocumentsStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ListDocumentsResponse) error
}

type docStoreListDocumentsStream struct {
	stream server.Stream
}

func (x *docStoreListDocumentsStream) Close() error {
	return x.stream.Close()
}

func (x *docStoreListDocumentsStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *docStoreListDocumentsStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *docStoreListDocumentsStream) Send(m *ListDocumentsResponse) error {
	return x.stream.Send(m)
}
