// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: install.proto

/*
Package install is a generated protocol buffer package.

It is generated from these files:
	install.proto

It has these top-level messages:
	InstallConfig
	ProxyConfig
	TLSSelfSigned
	TLSLetsEncrypt
	TLSCertificate
	CheckResult
	PerformCheckRequest
	PerformCheckResponse
	GetDefaultsRequest
	GetDefaultsResponse
	GetAgreementRequest
	GetAgreementResponse
	InstallRequest
	InstallResponse
	InstallEventsRequest
	InstallEventsResponse
*/
package install

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Install service

type InstallClient interface {
	GetDefaults(ctx context.Context, in *GetDefaultsRequest, opts ...client.CallOption) (*GetDefaultsResponse, error)
	Install(ctx context.Context, in *InstallRequest, opts ...client.CallOption) (*InstallResponse, error)
	PerformCheck(ctx context.Context, in *PerformCheckRequest, opts ...client.CallOption) (*PerformCheckResponse, error)
}

type installClient struct {
	c           client.Client
	serviceName string
}

func NewInstallClient(serviceName string, c client.Client) InstallClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "install"
	}
	return &installClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *installClient) GetDefaults(ctx context.Context, in *GetDefaultsRequest, opts ...client.CallOption) (*GetDefaultsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Install.GetDefaults", in)
	out := new(GetDefaultsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installClient) Install(ctx context.Context, in *InstallRequest, opts ...client.CallOption) (*InstallResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Install.Install", in)
	out := new(InstallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installClient) PerformCheck(ctx context.Context, in *PerformCheckRequest, opts ...client.CallOption) (*PerformCheckResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Install.PerformCheck", in)
	out := new(PerformCheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Install service

type InstallHandler interface {
	GetDefaults(context.Context, *GetDefaultsRequest, *GetDefaultsResponse) error
	Install(context.Context, *InstallRequest, *InstallResponse) error
	PerformCheck(context.Context, *PerformCheckRequest, *PerformCheckResponse) error
}

func RegisterInstallHandler(s server.Server, hdlr InstallHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Install{hdlr}, opts...))
}

type Install struct {
	InstallHandler
}

func (h *Install) GetDefaults(ctx context.Context, in *GetDefaultsRequest, out *GetDefaultsResponse) error {
	return h.InstallHandler.GetDefaults(ctx, in, out)
}

func (h *Install) Install(ctx context.Context, in *InstallRequest, out *InstallResponse) error {
	return h.InstallHandler.Install(ctx, in, out)
}

func (h *Install) PerformCheck(ctx context.Context, in *PerformCheckRequest, out *PerformCheckResponse) error {
	return h.InstallHandler.PerformCheck(ctx, in, out)
}
