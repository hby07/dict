// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/dict.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	libresp "github.com/pku-hit/libresp"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Dict service

func NewDictEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Dict service

type DictService interface {
	ListRoot(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*libresp.Response, error)
	ListChildren(ctx context.Context, in *wrappers.StringValue, opts ...client.CallOption) (*libresp.ListResponse, error)
	AddDict(ctx context.Context, in *AddDictRequest, opts ...client.CallOption) (*libresp.Response, error)
	DelDict(ctx context.Context, in *wrappers.StringValue, opts ...client.CallOption) (*libresp.Response, error)
	GetValue(ctx context.Context, in *wrappers.StringValue, opts ...client.CallOption) (*libresp.GenericResponse, error)
}

type dictService struct {
	c    client.Client
	name string
}

func NewDictService(name string, c client.Client) DictService {
	return &dictService{
		c:    c,
		name: name,
	}
}

func (c *dictService) ListRoot(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*libresp.Response, error) {
	req := c.c.NewRequest(c.name, "Dict.ListRoot", in)
	out := new(libresp.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictService) ListChildren(ctx context.Context, in *wrappers.StringValue, opts ...client.CallOption) (*libresp.ListResponse, error) {
	req := c.c.NewRequest(c.name, "Dict.ListChildren", in)
	out := new(libresp.ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictService) AddDict(ctx context.Context, in *AddDictRequest, opts ...client.CallOption) (*libresp.Response, error) {
	req := c.c.NewRequest(c.name, "Dict.AddDict", in)
	out := new(libresp.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictService) DelDict(ctx context.Context, in *wrappers.StringValue, opts ...client.CallOption) (*libresp.Response, error) {
	req := c.c.NewRequest(c.name, "Dict.DelDict", in)
	out := new(libresp.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictService) GetValue(ctx context.Context, in *wrappers.StringValue, opts ...client.CallOption) (*libresp.GenericResponse, error) {
	req := c.c.NewRequest(c.name, "Dict.GetValue", in)
	out := new(libresp.GenericResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dict service

type DictHandler interface {
	ListRoot(context.Context, *empty.Empty, *libresp.Response) error
	ListChildren(context.Context, *wrappers.StringValue, *libresp.ListResponse) error
	AddDict(context.Context, *AddDictRequest, *libresp.Response) error
	DelDict(context.Context, *wrappers.StringValue, *libresp.Response) error
	GetValue(context.Context, *wrappers.StringValue, *libresp.GenericResponse) error
}

func RegisterDictHandler(s server.Server, hdlr DictHandler, opts ...server.HandlerOption) error {
	type dict interface {
		ListRoot(ctx context.Context, in *empty.Empty, out *libresp.Response) error
		ListChildren(ctx context.Context, in *wrappers.StringValue, out *libresp.ListResponse) error
		AddDict(ctx context.Context, in *AddDictRequest, out *libresp.Response) error
		DelDict(ctx context.Context, in *wrappers.StringValue, out *libresp.Response) error
		GetValue(ctx context.Context, in *wrappers.StringValue, out *libresp.GenericResponse) error
	}
	type Dict struct {
		dict
	}
	h := &dictHandler{hdlr}
	return s.Handle(s.NewHandler(&Dict{h}, opts...))
}

type dictHandler struct {
	DictHandler
}

func (h *dictHandler) ListRoot(ctx context.Context, in *empty.Empty, out *libresp.Response) error {
	return h.DictHandler.ListRoot(ctx, in, out)
}

func (h *dictHandler) ListChildren(ctx context.Context, in *wrappers.StringValue, out *libresp.ListResponse) error {
	return h.DictHandler.ListChildren(ctx, in, out)
}

func (h *dictHandler) AddDict(ctx context.Context, in *AddDictRequest, out *libresp.Response) error {
	return h.DictHandler.AddDict(ctx, in, out)
}

func (h *dictHandler) DelDict(ctx context.Context, in *wrappers.StringValue, out *libresp.Response) error {
	return h.DictHandler.DelDict(ctx, in, out)
}

func (h *dictHandler) GetValue(ctx context.Context, in *wrappers.StringValue, out *libresp.GenericResponse) error {
	return h.DictHandler.GetValue(ctx, in, out)
}
