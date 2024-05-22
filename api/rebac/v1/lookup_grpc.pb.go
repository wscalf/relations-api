// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: rebac/v1/lookup.proto

package v1

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
	Lookup_Subjects_FullMethodName = "/api.rebac.v1.Lookup/Subjects"
)

// LookupClient is the client API for Lookup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LookupClient interface {
	Subjects(ctx context.Context, in *LookupSubjectsRequest, opts ...grpc.CallOption) (Lookup_SubjectsClient, error)
}

type lookupClient struct {
	cc grpc.ClientConnInterface
}

func NewLookupClient(cc grpc.ClientConnInterface) LookupClient {
	return &lookupClient{cc}
}

func (c *lookupClient) Subjects(ctx context.Context, in *LookupSubjectsRequest, opts ...grpc.CallOption) (Lookup_SubjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Lookup_ServiceDesc.Streams[0], Lookup_Subjects_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &lookupSubjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Lookup_SubjectsClient interface {
	Recv() (*LookupSubjectsResponse, error)
	grpc.ClientStream
}

type lookupSubjectsClient struct {
	grpc.ClientStream
}

func (x *lookupSubjectsClient) Recv() (*LookupSubjectsResponse, error) {
	m := new(LookupSubjectsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LookupServer is the server API for Lookup service.
// All implementations must embed UnimplementedLookupServer
// for forward compatibility
type LookupServer interface {
	Subjects(*LookupSubjectsRequest, Lookup_SubjectsServer) error
	mustEmbedUnimplementedLookupServer()
}

// UnimplementedLookupServer must be embedded to have forward compatible implementations.
type UnimplementedLookupServer struct {
}

func (UnimplementedLookupServer) Subjects(*LookupSubjectsRequest, Lookup_SubjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method Subjects not implemented")
}
func (UnimplementedLookupServer) mustEmbedUnimplementedLookupServer() {}

// UnsafeLookupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LookupServer will
// result in compilation errors.
type UnsafeLookupServer interface {
	mustEmbedUnimplementedLookupServer()
}

func RegisterLookupServer(s grpc.ServiceRegistrar, srv LookupServer) {
	s.RegisterService(&Lookup_ServiceDesc, srv)
}

func _Lookup_Subjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LookupSubjectsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LookupServer).Subjects(m, &lookupSubjectsServer{stream})
}

type Lookup_SubjectsServer interface {
	Send(*LookupSubjectsResponse) error
	grpc.ServerStream
}

type lookupSubjectsServer struct {
	grpc.ServerStream
}

func (x *lookupSubjectsServer) Send(m *LookupSubjectsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Lookup_ServiceDesc is the grpc.ServiceDesc for Lookup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lookup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.rebac.v1.Lookup",
	HandlerType: (*LookupServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subjects",
			Handler:       _Lookup_Subjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rebac/v1/lookup.proto",
}