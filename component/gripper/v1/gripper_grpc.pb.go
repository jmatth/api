// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: component/gripper/v1/gripper.proto

package v1

import (
	context "context"
	v1 "go.viam.com/api/common/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GripperServiceClient is the client API for GripperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GripperServiceClient interface {
	// Open opens a gripper of the underlying robot.
	Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error)
	// Grab requests a gripper of the underlying robot to grab.
	Grab(ctx context.Context, in *GrabRequest, opts ...grpc.CallOption) (*GrabResponse, error)
	// Stop stops a robot's gripper
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	// IsMoving reports if a component is in motion
	IsMoving(ctx context.Context, in *IsMovingRequest, opts ...grpc.CallOption) (*IsMovingResponse, error)
	// IsHoldingSomething returns whether the gripper is currently holding onto an object
	IsHoldingSomething(ctx context.Context, in *IsHoldingSomethingRequest, opts ...grpc.CallOption) (*IsHoldingSomethingResponse, error)
	// DoCommand sends/receives arbitrary commands
	DoCommand(ctx context.Context, in *v1.DoCommandRequest, opts ...grpc.CallOption) (*v1.DoCommandResponse, error)
	// GetGeometries returns the geometries of the component in their current configuration
	GetGeometries(ctx context.Context, in *v1.GetGeometriesRequest, opts ...grpc.CallOption) (*v1.GetGeometriesResponse, error)
	// GetKinematics returns the kinematics file for the component
	GetKinematics(ctx context.Context, in *v1.GetKinematicsRequest, opts ...grpc.CallOption) (*v1.GetKinematicsResponse, error)
}

type gripperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGripperServiceClient(cc grpc.ClientConnInterface) GripperServiceClient {
	return &gripperServiceClient{cc}
}

func (c *gripperServiceClient) Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error) {
	out := new(OpenResponse)
	err := c.cc.Invoke(ctx, "/viam.component.gripper.v1.GripperService/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gripperServiceClient) Grab(ctx context.Context, in *GrabRequest, opts ...grpc.CallOption) (*GrabResponse, error) {
	out := new(GrabResponse)
	err := c.cc.Invoke(ctx, "/viam.component.gripper.v1.GripperService/Grab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gripperServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/viam.component.gripper.v1.GripperService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gripperServiceClient) IsMoving(ctx context.Context, in *IsMovingRequest, opts ...grpc.CallOption) (*IsMovingResponse, error) {
	out := new(IsMovingResponse)
	err := c.cc.Invoke(ctx, "/viam.component.gripper.v1.GripperService/IsMoving", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gripperServiceClient) IsHoldingSomething(ctx context.Context, in *IsHoldingSomethingRequest, opts ...grpc.CallOption) (*IsHoldingSomethingResponse, error) {
	out := new(IsHoldingSomethingResponse)
	err := c.cc.Invoke(ctx, "/viam.component.gripper.v1.GripperService/IsHoldingSomething", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gripperServiceClient) DoCommand(ctx context.Context, in *v1.DoCommandRequest, opts ...grpc.CallOption) (*v1.DoCommandResponse, error) {
	out := new(v1.DoCommandResponse)
	err := c.cc.Invoke(ctx, "/viam.component.gripper.v1.GripperService/DoCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gripperServiceClient) GetGeometries(ctx context.Context, in *v1.GetGeometriesRequest, opts ...grpc.CallOption) (*v1.GetGeometriesResponse, error) {
	out := new(v1.GetGeometriesResponse)
	err := c.cc.Invoke(ctx, "/viam.component.gripper.v1.GripperService/GetGeometries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gripperServiceClient) GetKinematics(ctx context.Context, in *v1.GetKinematicsRequest, opts ...grpc.CallOption) (*v1.GetKinematicsResponse, error) {
	out := new(v1.GetKinematicsResponse)
	err := c.cc.Invoke(ctx, "/viam.component.gripper.v1.GripperService/GetKinematics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GripperServiceServer is the server API for GripperService service.
// All implementations must embed UnimplementedGripperServiceServer
// for forward compatibility
type GripperServiceServer interface {
	// Open opens a gripper of the underlying robot.
	Open(context.Context, *OpenRequest) (*OpenResponse, error)
	// Grab requests a gripper of the underlying robot to grab.
	Grab(context.Context, *GrabRequest) (*GrabResponse, error)
	// Stop stops a robot's gripper
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	// IsMoving reports if a component is in motion
	IsMoving(context.Context, *IsMovingRequest) (*IsMovingResponse, error)
	// IsHoldingSomething returns whether the gripper is currently holding onto an object
	IsHoldingSomething(context.Context, *IsHoldingSomethingRequest) (*IsHoldingSomethingResponse, error)
	// DoCommand sends/receives arbitrary commands
	DoCommand(context.Context, *v1.DoCommandRequest) (*v1.DoCommandResponse, error)
	// GetGeometries returns the geometries of the component in their current configuration
	GetGeometries(context.Context, *v1.GetGeometriesRequest) (*v1.GetGeometriesResponse, error)
	// GetKinematics returns the kinematics file for the component
	GetKinematics(context.Context, *v1.GetKinematicsRequest) (*v1.GetKinematicsResponse, error)
	mustEmbedUnimplementedGripperServiceServer()
}

// UnimplementedGripperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGripperServiceServer struct {
}

func (UnimplementedGripperServiceServer) Open(context.Context, *OpenRequest) (*OpenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedGripperServiceServer) Grab(context.Context, *GrabRequest) (*GrabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Grab not implemented")
}
func (UnimplementedGripperServiceServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedGripperServiceServer) IsMoving(context.Context, *IsMovingRequest) (*IsMovingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsMoving not implemented")
}
func (UnimplementedGripperServiceServer) IsHoldingSomething(context.Context, *IsHoldingSomethingRequest) (*IsHoldingSomethingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsHoldingSomething not implemented")
}
func (UnimplementedGripperServiceServer) DoCommand(context.Context, *v1.DoCommandRequest) (*v1.DoCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoCommand not implemented")
}
func (UnimplementedGripperServiceServer) GetGeometries(context.Context, *v1.GetGeometriesRequest) (*v1.GetGeometriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeometries not implemented")
}
func (UnimplementedGripperServiceServer) GetKinematics(context.Context, *v1.GetKinematicsRequest) (*v1.GetKinematicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKinematics not implemented")
}
func (UnimplementedGripperServiceServer) mustEmbedUnimplementedGripperServiceServer() {}

// UnsafeGripperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GripperServiceServer will
// result in compilation errors.
type UnsafeGripperServiceServer interface {
	mustEmbedUnimplementedGripperServiceServer()
}

func RegisterGripperServiceServer(s grpc.ServiceRegistrar, srv GripperServiceServer) {
	s.RegisterService(&GripperService_ServiceDesc, srv)
}

func _GripperService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripperServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.gripper.v1.GripperService/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripperServiceServer).Open(ctx, req.(*OpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GripperService_Grab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrabRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripperServiceServer).Grab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.gripper.v1.GripperService/Grab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripperServiceServer).Grab(ctx, req.(*GrabRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GripperService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripperServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.gripper.v1.GripperService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripperServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GripperService_IsMoving_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsMovingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripperServiceServer).IsMoving(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.gripper.v1.GripperService/IsMoving",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripperServiceServer).IsMoving(ctx, req.(*IsMovingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GripperService_IsHoldingSomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsHoldingSomethingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripperServiceServer).IsHoldingSomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.gripper.v1.GripperService/IsHoldingSomething",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripperServiceServer).IsHoldingSomething(ctx, req.(*IsHoldingSomethingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GripperService_DoCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DoCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripperServiceServer).DoCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.gripper.v1.GripperService/DoCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripperServiceServer).DoCommand(ctx, req.(*v1.DoCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GripperService_GetGeometries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetGeometriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripperServiceServer).GetGeometries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.gripper.v1.GripperService/GetGeometries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripperServiceServer).GetGeometries(ctx, req.(*v1.GetGeometriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GripperService_GetKinematics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetKinematicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripperServiceServer).GetKinematics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.gripper.v1.GripperService/GetKinematics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripperServiceServer).GetKinematics(ctx, req.(*v1.GetKinematicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GripperService_ServiceDesc is the grpc.ServiceDesc for GripperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GripperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "viam.component.gripper.v1.GripperService",
	HandlerType: (*GripperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _GripperService_Open_Handler,
		},
		{
			MethodName: "Grab",
			Handler:    _GripperService_Grab_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _GripperService_Stop_Handler,
		},
		{
			MethodName: "IsMoving",
			Handler:    _GripperService_IsMoving_Handler,
		},
		{
			MethodName: "IsHoldingSomething",
			Handler:    _GripperService_IsHoldingSomething_Handler,
		},
		{
			MethodName: "DoCommand",
			Handler:    _GripperService_DoCommand_Handler,
		},
		{
			MethodName: "GetGeometries",
			Handler:    _GripperService_GetGeometries_Handler,
		},
		{
			MethodName: "GetKinematics",
			Handler:    _GripperService_GetKinematics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "component/gripper/v1/gripper.proto",
}
