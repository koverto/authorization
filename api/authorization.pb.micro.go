// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: authorization.proto

package authorization

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/koverto/uuid"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Authorization service

type AuthorizationService interface {
	Create(ctx context.Context, in *Claims, opts ...client.CallOption) (*Token, error)
	Validate(ctx context.Context, in *Token, opts ...client.CallOption) (*Claims, error)
	Invalidate(ctx context.Context, in *Claims, opts ...client.CallOption) (*Claims, error)
}

type authorizationService struct {
	c    client.Client
	name string
}

func NewAuthorizationService(name string, c client.Client) AuthorizationService {
	return &authorizationService{
		c:    c,
		name: name,
	}
}

func (c *authorizationService) Create(ctx context.Context, in *Claims, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.name, "Authorization.Create", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationService) Validate(ctx context.Context, in *Token, opts ...client.CallOption) (*Claims, error) {
	req := c.c.NewRequest(c.name, "Authorization.Validate", in)
	out := new(Claims)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationService) Invalidate(ctx context.Context, in *Claims, opts ...client.CallOption) (*Claims, error) {
	req := c.c.NewRequest(c.name, "Authorization.Invalidate", in)
	out := new(Claims)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authorization service

type AuthorizationHandler interface {
	Create(context.Context, *Claims, *Token) error
	Validate(context.Context, *Token, *Claims) error
	Invalidate(context.Context, *Claims, *Claims) error
}

func RegisterAuthorizationHandler(s server.Server, hdlr AuthorizationHandler, opts ...server.HandlerOption) error {
	type authorization interface {
		Create(ctx context.Context, in *Claims, out *Token) error
		Validate(ctx context.Context, in *Token, out *Claims) error
		Invalidate(ctx context.Context, in *Claims, out *Claims) error
	}
	type Authorization struct {
		authorization
	}
	h := &authorizationHandler{hdlr}
	return s.Handle(s.NewHandler(&Authorization{h}, opts...))
}

type authorizationHandler struct {
	AuthorizationHandler
}

func (h *authorizationHandler) Create(ctx context.Context, in *Claims, out *Token) error {
	return h.AuthorizationHandler.Create(ctx, in, out)
}

func (h *authorizationHandler) Validate(ctx context.Context, in *Token, out *Claims) error {
	return h.AuthorizationHandler.Validate(ctx, in, out)
}

func (h *authorizationHandler) Invalidate(ctx context.Context, in *Claims, out *Claims) error {
	return h.AuthorizationHandler.Invalidate(ctx, in, out)
}
