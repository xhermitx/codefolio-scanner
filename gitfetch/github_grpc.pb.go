// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: github.proto

package gitfetch

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Github_FetchData_FullMethodName = "/github/FetchData"
)

// GithubClient is the client API for Github service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GithubClient interface {
	FetchData(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Response, error)
}

type githubClient struct {
	cc grpc.ClientConnInterface
}

func NewGithubClient(cc grpc.ClientConnInterface) GithubClient {
	return &githubClient{cc}
}

func (c *githubClient) FetchData(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Github_FetchData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GithubServer is the server API for Github service.
// All implementations must embed UnimplementedGithubServer
// for forward compatibility
type GithubServer interface {
	FetchData(context.Context, *Profile) (*Response, error)
	mustEmbedUnimplementedGithubServer()
}

// UnimplementedGithubServer must be embedded to have forward compatible implementations.
type UnimplementedGithubServer struct {
}

func (UnimplementedGithubServer) FetchData(context.Context, *Profile) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchData not implemented")
}
func (UnimplementedGithubServer) mustEmbedUnimplementedGithubServer() {}

// UnsafeGithubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GithubServer will
// result in compilation errors.
type UnsafeGithubServer interface {
	mustEmbedUnimplementedGithubServer()
}

func RegisterGithubServer(s grpc.ServiceRegistrar, srv GithubServer) {
	s.RegisterService(&Github_ServiceDesc, srv)
}

func _Github_FetchData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubServer).FetchData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Github_FetchData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubServer).FetchData(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

// Github_ServiceDesc is the grpc.ServiceDesc for Github service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Github_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github",
	HandlerType: (*GithubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchData",
			Handler:    _Github_FetchData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.proto",
}