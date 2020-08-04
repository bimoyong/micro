// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service/debug/trace/proto/trace.proto

package go_micro_debug_trace

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
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

// Api Endpoints for Trace service

func NewTraceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Trace service

type TraceService interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*WriteResponse, error)
	Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (Trace_StreamService, error)
}

type traceService struct {
	c    client.Client
	name string
}

func NewTraceService(name string, c client.Client) TraceService {
	return &traceService{
		c:    c,
		name: name,
	}
}

func (c *traceService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Trace.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceService) Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*WriteResponse, error) {
	req := c.c.NewRequest(c.name, "Trace.Write", in)
	out := new(WriteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceService) Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (Trace_StreamService, error) {
	req := c.c.NewRequest(c.name, "Trace.Stream", &StreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &traceServiceStream{stream}, nil
}

type Trace_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamResponse, error)
}

type traceServiceStream struct {
	stream client.Stream
}

func (x *traceServiceStream) Close() error {
	return x.stream.Close()
}

func (x *traceServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *traceServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *traceServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *traceServiceStream) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Trace service

type TraceHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Write(context.Context, *WriteRequest, *WriteResponse) error
	Stream(context.Context, *StreamRequest, Trace_StreamStream) error
}

func RegisterTraceHandler(s server.Server, hdlr TraceHandler, opts ...server.HandlerOption) error {
	type trace interface {
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Write(ctx context.Context, in *WriteRequest, out *WriteResponse) error
		Stream(ctx context.Context, stream server.Stream) error
	}
	type Trace struct {
		trace
	}
	h := &traceHandler{hdlr}
	return s.Handle(s.NewHandler(&Trace{h}, opts...))
}

type traceHandler struct {
	TraceHandler
}

func (h *traceHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.TraceHandler.Read(ctx, in, out)
}

func (h *traceHandler) Write(ctx context.Context, in *WriteRequest, out *WriteResponse) error {
	return h.TraceHandler.Write(ctx, in, out)
}

func (h *traceHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.TraceHandler.Stream(ctx, m, &traceStreamStream{stream})
}

type Trace_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamResponse) error
}

type traceStreamStream struct {
	stream server.Stream
}

func (x *traceStreamStream) Close() error {
	return x.stream.Close()
}

func (x *traceStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *traceStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *traceStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *traceStreamStream) Send(m *StreamResponse) error {
	return x.stream.Send(m)
}
