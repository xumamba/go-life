// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stream

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// StudentsClient is the client API for Students service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentsClient interface {
	GetStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Feature, error)
	ListStudents(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (Students_ListStudentsClient, error)
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (Students_RecordRouteClient, error)
}

type studentsClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentsClient(cc grpc.ClientConnInterface) StudentsClient {
	return &studentsClient{cc}
}

func (c *studentsClient) GetStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := c.cc.Invoke(ctx, "/stream.Students/GetStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentsClient) ListStudents(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (Students_ListStudentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Students_serviceDesc.Streams[0], "/stream.Students/ListStudents", opts...)
	if err != nil {
		return nil, err
	}
	x := &studentsListStudentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Students_ListStudentsClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type studentsListStudentsClient struct {
	grpc.ClientStream
}

func (x *studentsListStudentsClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *studentsClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (Students_RecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Students_serviceDesc.Streams[1], "/stream.Students/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &studentsRecordRouteClient{stream}
	return x, nil
}

type Students_RecordRouteClient interface {
	Send(*Student) error
	CloseAndRecv() (*StudentSummary, error)
	grpc.ClientStream
}

type studentsRecordRouteClient struct {
	grpc.ClientStream
}

func (x *studentsRecordRouteClient) Send(m *Student) error {
	return x.ClientStream.SendMsg(m)
}

func (x *studentsRecordRouteClient) CloseAndRecv() (*StudentSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StudentSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StudentsServer is the server API for Students service.
// All implementations must embed UnimplementedStudentsServer
// for forward compatibility
type StudentsServer interface {
	GetStudent(context.Context, *Student) (*Feature, error)
	ListStudents(*Rectangle, Students_ListStudentsServer) error
	RecordRoute(Students_RecordRouteServer) error
	mustEmbedUnimplementedStudentsServer()
}

// UnimplementedStudentsServer must be embedded to have forward compatible implementations.
type UnimplementedStudentsServer struct {
}

func (UnimplementedStudentsServer) GetStudent(context.Context, *Student) (*Feature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (UnimplementedStudentsServer) ListStudents(*Rectangle, Students_ListStudentsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListStudents not implemented")
}
func (UnimplementedStudentsServer) RecordRoute(Students_RecordRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordRoute not implemented")
}
func (UnimplementedStudentsServer) mustEmbedUnimplementedStudentsServer() {}

// UnsafeStudentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentsServer will
// result in compilation errors.
type UnsafeStudentsServer interface {
	mustEmbedUnimplementedStudentsServer()
}

func RegisterStudentsServer(s grpc.ServiceRegistrar, srv StudentsServer) {
	s.RegisterService(&_Students_serviceDesc, srv)
}

func _Students_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentsServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.Students/GetStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentsServer).GetStudent(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

func _Students_ListStudents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StudentsServer).ListStudents(m, &studentsListStudentsServer{stream})
}

type Students_ListStudentsServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type studentsListStudentsServer struct {
	grpc.ServerStream
}

func (x *studentsListStudentsServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _Students_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StudentsServer).RecordRoute(&studentsRecordRouteServer{stream})
}

type Students_RecordRouteServer interface {
	SendAndClose(*StudentSummary) error
	Recv() (*Student, error)
	grpc.ServerStream
}

type studentsRecordRouteServer struct {
	grpc.ServerStream
}

func (x *studentsRecordRouteServer) SendAndClose(m *StudentSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *studentsRecordRouteServer) Recv() (*Student, error) {
	m := new(Student)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Students_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stream.Students",
	HandlerType: (*StudentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudent",
			Handler:    _Students_GetStudent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListStudents",
			Handler:       _Students_ListStudents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _Students_RecordRoute_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
