// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package enginepb

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

// DataRWServiceClient is the client API for DataRWService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataRWServiceClient interface {
	ReadLines(ctx context.Context, in *ReadLinesRequest, opts ...grpc.CallOption) (DataRWService_ReadLinesClient, error)
	WriteLines(ctx context.Context, opts ...grpc.CallOption) (DataRWService_WriteLinesClient, error)
	GenerateData(ctx context.Context, in *GenerateDataRequest, opts ...grpc.CallOption) (*GenerateDataResponse, error)
	ListFiles(ctx context.Context, in *ListFilesReq, opts ...grpc.CallOption) (*ListFilesResponse, error)
	IsReady(ctx context.Context, in *IsReadyRequest, opts ...grpc.CallOption) (*IsReadyResponse, error)
	CheckDir(ctx context.Context, in *CheckDirRequest, opts ...grpc.CallOption) (*CheckDirResponse, error)
}

type dataRWServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataRWServiceClient(cc grpc.ClientConnInterface) DataRWServiceClient {
	return &dataRWServiceClient{cc}
}

func (c *dataRWServiceClient) ReadLines(ctx context.Context, in *ReadLinesRequest, opts ...grpc.CallOption) (DataRWService_ReadLinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataRWService_ServiceDesc.Streams[0], "/enginepb.DataRWService/ReadLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataRWServiceReadLinesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataRWService_ReadLinesClient interface {
	Recv() (*ReadLinesResponse, error)
	grpc.ClientStream
}

type dataRWServiceReadLinesClient struct {
	grpc.ClientStream
}

func (x *dataRWServiceReadLinesClient) Recv() (*ReadLinesResponse, error) {
	m := new(ReadLinesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataRWServiceClient) WriteLines(ctx context.Context, opts ...grpc.CallOption) (DataRWService_WriteLinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataRWService_ServiceDesc.Streams[1], "/enginepb.DataRWService/WriteLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataRWServiceWriteLinesClient{stream}
	return x, nil
}

type DataRWService_WriteLinesClient interface {
	Send(*WriteLinesRequest) error
	CloseAndRecv() (*WriteLinesResponse, error)
	grpc.ClientStream
}

type dataRWServiceWriteLinesClient struct {
	grpc.ClientStream
}

func (x *dataRWServiceWriteLinesClient) Send(m *WriteLinesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataRWServiceWriteLinesClient) CloseAndRecv() (*WriteLinesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteLinesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataRWServiceClient) GenerateData(ctx context.Context, in *GenerateDataRequest, opts ...grpc.CallOption) (*GenerateDataResponse, error) {
	out := new(GenerateDataResponse)
	err := c.cc.Invoke(ctx, "/enginepb.DataRWService/GenerateData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataRWServiceClient) ListFiles(ctx context.Context, in *ListFilesReq, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	out := new(ListFilesResponse)
	err := c.cc.Invoke(ctx, "/enginepb.DataRWService/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataRWServiceClient) IsReady(ctx context.Context, in *IsReadyRequest, opts ...grpc.CallOption) (*IsReadyResponse, error) {
	out := new(IsReadyResponse)
	err := c.cc.Invoke(ctx, "/enginepb.DataRWService/IsReady", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataRWServiceClient) CheckDir(ctx context.Context, in *CheckDirRequest, opts ...grpc.CallOption) (*CheckDirResponse, error) {
	out := new(CheckDirResponse)
	err := c.cc.Invoke(ctx, "/enginepb.DataRWService/CheckDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataRWServiceServer is the server API for DataRWService service.
// All implementations should embed UnimplementedDataRWServiceServer
// for forward compatibility
type DataRWServiceServer interface {
	ReadLines(*ReadLinesRequest, DataRWService_ReadLinesServer) error
	WriteLines(DataRWService_WriteLinesServer) error
	GenerateData(context.Context, *GenerateDataRequest) (*GenerateDataResponse, error)
	ListFiles(context.Context, *ListFilesReq) (*ListFilesResponse, error)
	IsReady(context.Context, *IsReadyRequest) (*IsReadyResponse, error)
	CheckDir(context.Context, *CheckDirRequest) (*CheckDirResponse, error)
}

// UnimplementedDataRWServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDataRWServiceServer struct {
}

func (UnimplementedDataRWServiceServer) ReadLines(*ReadLinesRequest, DataRWService_ReadLinesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadLines not implemented")
}
func (UnimplementedDataRWServiceServer) WriteLines(DataRWService_WriteLinesServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteLines not implemented")
}
func (UnimplementedDataRWServiceServer) GenerateData(context.Context, *GenerateDataRequest) (*GenerateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateData not implemented")
}
func (UnimplementedDataRWServiceServer) ListFiles(context.Context, *ListFilesReq) (*ListFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedDataRWServiceServer) IsReady(context.Context, *IsReadyRequest) (*IsReadyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReady not implemented")
}
func (UnimplementedDataRWServiceServer) CheckDir(context.Context, *CheckDirRequest) (*CheckDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDir not implemented")
}

// UnsafeDataRWServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataRWServiceServer will
// result in compilation errors.
type UnsafeDataRWServiceServer interface {
	mustEmbedUnimplementedDataRWServiceServer()
}

func RegisterDataRWServiceServer(s grpc.ServiceRegistrar, srv DataRWServiceServer) {
	s.RegisterService(&DataRWService_ServiceDesc, srv)
}

func _DataRWService_ReadLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadLinesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataRWServiceServer).ReadLines(m, &dataRWServiceReadLinesServer{stream})
}

type DataRWService_ReadLinesServer interface {
	Send(*ReadLinesResponse) error
	grpc.ServerStream
}

type dataRWServiceReadLinesServer struct {
	grpc.ServerStream
}

func (x *dataRWServiceReadLinesServer) Send(m *ReadLinesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DataRWService_WriteLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataRWServiceServer).WriteLines(&dataRWServiceWriteLinesServer{stream})
}

type DataRWService_WriteLinesServer interface {
	SendAndClose(*WriteLinesResponse) error
	Recv() (*WriteLinesRequest, error)
	grpc.ServerStream
}

type dataRWServiceWriteLinesServer struct {
	grpc.ServerStream
}

func (x *dataRWServiceWriteLinesServer) SendAndClose(m *WriteLinesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataRWServiceWriteLinesServer) Recv() (*WriteLinesRequest, error) {
	m := new(WriteLinesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataRWService_GenerateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataRWServiceServer).GenerateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.DataRWService/GenerateData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataRWServiceServer).GenerateData(ctx, req.(*GenerateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataRWService_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataRWServiceServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.DataRWService/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataRWServiceServer).ListFiles(ctx, req.(*ListFilesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataRWService_IsReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataRWServiceServer).IsReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.DataRWService/IsReady",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataRWServiceServer).IsReady(ctx, req.(*IsReadyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataRWService_CheckDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataRWServiceServer).CheckDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.DataRWService/CheckDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataRWServiceServer).CheckDir(ctx, req.(*CheckDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataRWService_ServiceDesc is the grpc.ServiceDesc for DataRWService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataRWService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "enginepb.DataRWService",
	HandlerType: (*DataRWServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateData",
			Handler:    _DataRWService_GenerateData_Handler,
		},
		{
			MethodName: "ListFiles",
			Handler:    _DataRWService_ListFiles_Handler,
		},
		{
			MethodName: "IsReady",
			Handler:    _DataRWService_IsReady_Handler,
		},
		{
			MethodName: "CheckDir",
			Handler:    _DataRWService_CheckDir_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadLines",
			Handler:       _DataRWService_ReadLines_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WriteLines",
			Handler:       _DataRWService_WriteLines_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "engine/proto/datarw.proto",
}