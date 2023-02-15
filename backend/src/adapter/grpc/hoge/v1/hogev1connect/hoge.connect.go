// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: hoge/v1/hoge.proto

package hogev1connect

import (
	v1 "connect-tutorial/src/adapter/grpc/hoge/v1"
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
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
	// HogeServiceName is the fully-qualified name of the HogeService service.
	HogeServiceName = "hoge.v1.HogeService"
)

// HogeServiceClient is a client for the hoge.v1.HogeService service.
type HogeServiceClient interface {
	GetHoge(context.Context, *connect_go.Request[v1.HogeGetRequest]) (*connect_go.Response[v1.HogeGetResponse], error)
	CreateHoge(context.Context, *connect_go.Request[v1.HogeCreateRequest]) (*connect_go.Response[v1.HogeCreateResponse], error)
}

// NewHogeServiceClient constructs a client for the hoge.v1.HogeService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHogeServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) HogeServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &hogeServiceClient{
		getHoge: connect_go.NewClient[v1.HogeGetRequest, v1.HogeGetResponse](
			httpClient,
			baseURL+"/hoge.v1.HogeService/GetHoge",
			opts...,
		),
		createHoge: connect_go.NewClient[v1.HogeCreateRequest, v1.HogeCreateResponse](
			httpClient,
			baseURL+"/hoge.v1.HogeService/CreateHoge",
			opts...,
		),
	}
}

// hogeServiceClient implements HogeServiceClient.
type hogeServiceClient struct {
	getHoge    *connect_go.Client[v1.HogeGetRequest, v1.HogeGetResponse]
	createHoge *connect_go.Client[v1.HogeCreateRequest, v1.HogeCreateResponse]
}

// GetHoge calls hoge.v1.HogeService.GetHoge.
func (c *hogeServiceClient) GetHoge(ctx context.Context, req *connect_go.Request[v1.HogeGetRequest]) (*connect_go.Response[v1.HogeGetResponse], error) {
	return c.getHoge.CallUnary(ctx, req)
}

// CreateHoge calls hoge.v1.HogeService.CreateHoge.
func (c *hogeServiceClient) CreateHoge(ctx context.Context, req *connect_go.Request[v1.HogeCreateRequest]) (*connect_go.Response[v1.HogeCreateResponse], error) {
	return c.createHoge.CallUnary(ctx, req)
}

// HogeServiceHandler is an implementation of the hoge.v1.HogeService service.
type HogeServiceHandler interface {
	GetHoge(context.Context, *connect_go.Request[v1.HogeGetRequest]) (*connect_go.Response[v1.HogeGetResponse], error)
	CreateHoge(context.Context, *connect_go.Request[v1.HogeCreateRequest]) (*connect_go.Response[v1.HogeCreateResponse], error)
}

// NewHogeServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHogeServiceHandler(svc HogeServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/hoge.v1.HogeService/GetHoge", connect_go.NewUnaryHandler(
		"/hoge.v1.HogeService/GetHoge",
		svc.GetHoge,
		opts...,
	))
	mux.Handle("/hoge.v1.HogeService/CreateHoge", connect_go.NewUnaryHandler(
		"/hoge.v1.HogeService/CreateHoge",
		svc.CreateHoge,
		opts...,
	))
	return "/hoge.v1.HogeService/", mux
}

// UnimplementedHogeServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedHogeServiceHandler struct{}

func (UnimplementedHogeServiceHandler) GetHoge(context.Context, *connect_go.Request[v1.HogeGetRequest]) (*connect_go.Response[v1.HogeGetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("hoge.v1.HogeService.GetHoge is not implemented"))
}

func (UnimplementedHogeServiceHandler) CreateHoge(context.Context, *connect_go.Request[v1.HogeCreateRequest]) (*connect_go.Response[v1.HogeCreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("hoge.v1.HogeService.CreateHoge is not implemented"))
}