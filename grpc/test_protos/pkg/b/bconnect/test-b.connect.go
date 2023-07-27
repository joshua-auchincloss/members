// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: test_protos/pkg/b/test-b.proto

package bconnect

import (
	context "context"
	errors "errors"
	b "members/grpc/test_protos/pkg/b"
	http "net/http"
	strings "strings"

	connect_go "github.com/bufbuild/connect-go"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// HelloServiceName is the fully-qualified name of the HelloService service.
	HelloServiceName = "pbtestb.HelloService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// HelloServiceSayHelloProcedure is the fully-qualified name of the HelloService's SayHello RPC.
	HelloServiceSayHelloProcedure = "/pbtestb.HelloService/SayHello"
)

// HelloServiceClient is a client for the pbtestb.HelloService service.
type HelloServiceClient interface {
	SayHello(context.Context, *connect_go.Request[b.Outer]) (*connect_go.Response[b.OuterInner], error)
}

// NewHelloServiceClient constructs a client for the pbtestb.HelloService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHelloServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) HelloServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &helloServiceClient{
		sayHello: connect_go.NewClient[b.Outer, b.OuterInner](
			httpClient,
			baseURL+HelloServiceSayHelloProcedure,
			opts...,
		),
	}
}

// helloServiceClient implements HelloServiceClient.
type helloServiceClient struct {
	sayHello *connect_go.Client[b.Outer, b.OuterInner]
}

// SayHello calls pbtestb.HelloService.SayHello.
func (c *helloServiceClient) SayHello(ctx context.Context, req *connect_go.Request[b.Outer]) (*connect_go.Response[b.OuterInner], error) {
	return c.sayHello.CallUnary(ctx, req)
}

// HelloServiceHandler is an implementation of the pbtestb.HelloService service.
type HelloServiceHandler interface {
	SayHello(context.Context, *connect_go.Request[b.Outer]) (*connect_go.Response[b.OuterInner], error)
}

// NewHelloServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHelloServiceHandler(svc HelloServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	helloServiceSayHelloHandler := connect_go.NewUnaryHandler(
		HelloServiceSayHelloProcedure,
		svc.SayHello,
		opts...,
	)
	return "/pbtestb.HelloService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HelloServiceSayHelloProcedure:
			helloServiceSayHelloHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedHelloServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedHelloServiceHandler struct{}

func (UnimplementedHelloServiceHandler) SayHello(context.Context, *connect_go.Request[b.Outer]) (*connect_go.Response[b.OuterInner], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("pbtestb.HelloService.SayHello is not implemented"))
}
