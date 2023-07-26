// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/health/svc.proto

package healthconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	_ "members/grpc/api/v1/health"
	pkg "members/grpc/api/v1/health/pkg"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// HealthName is the fully-qualified name of the Health service.
	HealthName = "members.v1.health.svc.Health"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// HealthCheckProcedure is the fully-qualified name of the Health's Check RPC.
	HealthCheckProcedure = "/members.v1.health.svc.Health/Check"
	// HealthWatchProcedure is the fully-qualified name of the Health's Watch RPC.
	HealthWatchProcedure = "/members.v1.health.svc.Health/Watch"
)

// HealthClient is a client for the members.v1.health.svc.Health service.
type HealthClient interface {
	Check(context.Context, *connect_go.Request[pkg.HealthCheckRequest]) (*connect_go.Response[pkg.HealthCheckResponse], error)
	Watch(context.Context, *connect_go.Request[pkg.HealthCheckRequest]) (*connect_go.ServerStreamForClient[pkg.HealthCheckResponse], error)
}

// NewHealthClient constructs a client for the members.v1.health.svc.Health service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHealthClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) HealthClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &healthClient{
		check: connect_go.NewClient[pkg.HealthCheckRequest, pkg.HealthCheckResponse](
			httpClient,
			baseURL+HealthCheckProcedure,
			opts...,
		),
		watch: connect_go.NewClient[pkg.HealthCheckRequest, pkg.HealthCheckResponse](
			httpClient,
			baseURL+HealthWatchProcedure,
			opts...,
		),
	}
}

// healthClient implements HealthClient.
type healthClient struct {
	check *connect_go.Client[pkg.HealthCheckRequest, pkg.HealthCheckResponse]
	watch *connect_go.Client[pkg.HealthCheckRequest, pkg.HealthCheckResponse]
}

// Check calls members.v1.health.svc.Health.Check.
func (c *healthClient) Check(ctx context.Context, req *connect_go.Request[pkg.HealthCheckRequest]) (*connect_go.Response[pkg.HealthCheckResponse], error) {
	return c.check.CallUnary(ctx, req)
}

// Watch calls members.v1.health.svc.Health.Watch.
func (c *healthClient) Watch(ctx context.Context, req *connect_go.Request[pkg.HealthCheckRequest]) (*connect_go.ServerStreamForClient[pkg.HealthCheckResponse], error) {
	return c.watch.CallServerStream(ctx, req)
}

// HealthHandler is an implementation of the members.v1.health.svc.Health service.
type HealthHandler interface {
	Check(context.Context, *connect_go.Request[pkg.HealthCheckRequest]) (*connect_go.Response[pkg.HealthCheckResponse], error)
	Watch(context.Context, *connect_go.Request[pkg.HealthCheckRequest], *connect_go.ServerStream[pkg.HealthCheckResponse]) error
}

// NewHealthHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHealthHandler(svc HealthHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	healthCheckHandler := connect_go.NewUnaryHandler(
		HealthCheckProcedure,
		svc.Check,
		opts...,
	)
	healthWatchHandler := connect_go.NewServerStreamHandler(
		HealthWatchProcedure,
		svc.Watch,
		opts...,
	)
	return "/members.v1.health.svc.Health/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HealthCheckProcedure:
			healthCheckHandler.ServeHTTP(w, r)
		case HealthWatchProcedure:
			healthWatchHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedHealthHandler returns CodeUnimplemented from all methods.
type UnimplementedHealthHandler struct{}

func (UnimplementedHealthHandler) Check(context.Context, *connect_go.Request[pkg.HealthCheckRequest]) (*connect_go.Response[pkg.HealthCheckResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("members.v1.health.svc.Health.Check is not implemented"))
}

func (UnimplementedHealthHandler) Watch(context.Context, *connect_go.Request[pkg.HealthCheckRequest], *connect_go.ServerStream[pkg.HealthCheckResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("members.v1.health.svc.Health.Watch is not implemented"))
}
