// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.0
// source: ss_student.proto

package student

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

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	CreateStudent(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*Student, error)
	GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*Student, error)
	ListStudents(ctx context.Context, in *ListStudentRequest, opts ...grpc.CallOption) (*ListStudentResponse, error)
	PatchStudent(ctx context.Context, in *UpdateStudentRequest, opts ...grpc.CallOption) (*Student, error)
	DeleteStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) CreateStudent(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/student.StudentService/CreateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/student.StudentService/GetStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) ListStudents(ctx context.Context, in *ListStudentRequest, opts ...grpc.CallOption) (*ListStudentResponse, error) {
	out := new(ListStudentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/ListStudents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) PatchStudent(ctx context.Context, in *UpdateStudentRequest, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/student.StudentService/PatchStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) DeleteStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/DeleteStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	CreateStudent(context.Context, *CreateStudentRequest) (*Student, error)
	GetStudent(context.Context, *GetStudentRequest) (*Student, error)
	ListStudents(context.Context, *ListStudentRequest) (*ListStudentResponse, error)
	PatchStudent(context.Context, *UpdateStudentRequest) (*Student, error)
	DeleteStudent(context.Context, *GetStudentRequest) (*SimpleResponse, error)
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) CreateStudent(context.Context, *CreateStudentRequest) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (UnimplementedStudentServiceServer) GetStudent(context.Context, *GetStudentRequest) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (UnimplementedStudentServiceServer) ListStudents(context.Context, *ListStudentRequest) (*ListStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStudents not implemented")
}
func (UnimplementedStudentServiceServer) PatchStudent(context.Context, *UpdateStudentRequest) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchStudent not implemented")
}
func (UnimplementedStudentServiceServer) DeleteStudent(context.Context, *GetStudentRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudent not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_CreateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).CreateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/CreateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).CreateStudent(ctx, req.(*CreateStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/GetStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudent(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_ListStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).ListStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/ListStudents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).ListStudents(ctx, req.(*ListStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_PatchStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).PatchStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/PatchStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).PatchStudent(ctx, req.(*UpdateStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_DeleteStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).DeleteStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/DeleteStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).DeleteStudent(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStudent",
			Handler:    _StudentService_CreateStudent_Handler,
		},
		{
			MethodName: "GetStudent",
			Handler:    _StudentService_GetStudent_Handler,
		},
		{
			MethodName: "ListStudents",
			Handler:    _StudentService_ListStudents_Handler,
		},
		{
			MethodName: "PatchStudent",
			Handler:    _StudentService_PatchStudent_Handler,
		},
		{
			MethodName: "DeleteStudent",
			Handler:    _StudentService_DeleteStudent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ss_student.proto",
}
