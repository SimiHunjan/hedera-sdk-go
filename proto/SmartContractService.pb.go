// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/SmartContractService.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() {
	proto.RegisterFile("proto/SmartContractService.proto", fileDescriptor_858920b839a6ded9)
}

var fileDescriptor_858920b839a6ded9 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x4f, 0xf3, 0x10, 0xc6, 0x86, 0x75, 0x30, 0xe8, 0x45, 0x8f, 0x5e, 0xd6, 0x8a, 0x7a,
	0xf0, 0x24, 0xa3, 0x1b, 0xca, 0x40, 0x0f, 0x6e, 0xf3, 0xe2, 0x2d, 0x4b, 0x9e, 0x6d, 0xb1, 0xcd,
	0x2b, 0xc9, 0x9b, 0x98, 0x7f, 0xcb, 0xbf, 0x50, 0x4c, 0xda, 0x32, 0x9c, 0x30, 0xdd, 0x4e, 0xa1,
	0xef, 0xfb, 0x7e, 0xdf, 0xf7, 0xd2, 0x96, 0x9d, 0x55, 0x1a, 0x09, 0xe3, 0x45, 0xc9, 0x35, 0x4d,
	0x50, 0x91, 0xe6, 0x82, 0x16, 0xa0, 0xdf, 0x73, 0x01, 0x91, 0x93, 0x82, 0x8e, 0x3b, 0xc2, 0x53,
	0x6f, 0x5c, 0x6a, 0xae, 0x0c, 0x17, 0x94, 0xa3, 0x9a, 0x83, 0xa9, 0x50, 0x99, 0xda, 0x17, 0x1e,
	0x7b, 0xc3, 0xd3, 0x1a, 0xb4, 0xad, 0x47, 0x03, 0x3f, 0xfa, 0x61, 0x1c, 0x6e, 0x25, 0x79, 0xe1,
	0xf2, 0xb3, 0xc3, 0x06, 0xbf, 0x2d, 0x12, 0x8c, 0x59, 0x4f, 0x68, 0xe0, 0x04, 0x8d, 0x10, 0x04,
	0x1e, 0x89, 0x36, 0x42, 0xc2, 0x70, 0x7b, 0xd6, 0x34, 0x7f, 0x27, 0xac, 0x2b, 0x79, 0x48, 0xc2,
	0x94, 0x05, 0xa2, 0x66, 0x27, 0xbc, 0x28, 0x1e, 0x81, 0x32, 0x94, 0xff, 0x4e, 0xb9, 0x60, 0xfd,
	0x14, 0xda, 0xfb, 0xcd, 0xd4, 0x2b, 0x06, 0xdd, 0xda, 0xee, 0x5e, 0x5c, 0xd8, 0xaf, 0x9f, 0x5a,
	0xe2, 0x86, 0x0d, 0x37, 0x7b, 0x1f, 0x50, 0xf0, 0xa6, 0x7c, 0x07, 0x79, 0xcd, 0x4e, 0x9a, 0xa2,
	0x7b, 0xa0, 0xc4, 0x12, 0x08, 0x94, 0xb0, 0x8b, 0xf2, 0x1b, 0x26, 0x76, 0x81, 0x45, 0x2e, 0x73,
	0xb2, 0xb3, 0xe9, 0x1f, 0x36, 0x4c, 0x81, 0x96, 0x1f, 0x73, 0x10, 0xa8, 0x65, 0x62, 0xdb, 0xdb,
	0xed, 0x24, 0xc7, 0xac, 0x27, 0xa1, 0x80, 0x03, 0xbe, 0xca, 0x2d, 0xeb, 0x1a, 0x6b, 0x08, 0xca,
	0xa9, 0xcb, 0xd9, 0xe7, 0xbf, 0xf0, 0xfc, 0xb3, 0x92, 0x7b, 0x25, 0x24, 0x77, 0x2c, 0x14, 0x58,
	0x46, 0x19, 0x48, 0xd0, 0x3c, 0xca, 0xb8, 0xc9, 0x52, 0xcd, 0xab, 0xcc, 0x13, 0x2f, 0xe7, 0x69,
	0x4e, 0xd9, 0x7a, 0x15, 0x09, 0x2c, 0xe3, 0x56, 0x8b, 0xbd, 0x79, 0x64, 0xe4, 0xdb, 0x28, 0xc5,
	0xd8, 0x39, 0x57, 0x47, 0xee, 0xb8, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x96, 0xda, 0xd5, 0x12,
	0x91, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SmartContractServiceClient is the client API for SmartContractService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmartContractServiceClient interface {
	CreateContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	UpdateContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	ContractCallMethod(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	GetContractInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	ContractCallLocalMethod(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	ContractGetBytecode(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	GetBySolidityID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	GetTxRecordByContractID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	DeleteContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	SystemDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	SystemUndelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type smartContractServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmartContractServiceClient(cc grpc.ClientConnInterface) SmartContractServiceClient {
	return &smartContractServiceClient{cc}
}

func (c *smartContractServiceClient) CreateContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/createContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) UpdateContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/updateContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) ContractCallMethod(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/contractCallMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) GetContractInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/getContractInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) ContractCallLocalMethod(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/contractCallLocalMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) ContractGetBytecode(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/ContractGetBytecode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) GetBySolidityID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/getBySolidityID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) GetTxRecordByContractID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/getTxRecordByContractID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) DeleteContract(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/deleteContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) SystemDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/systemDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartContractServiceClient) SystemUndelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.SmartContractService/systemUndelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmartContractServiceServer is the server API for SmartContractService service.
type SmartContractServiceServer interface {
	CreateContract(context.Context, *Transaction) (*TransactionResponse, error)
	UpdateContract(context.Context, *Transaction) (*TransactionResponse, error)
	ContractCallMethod(context.Context, *Transaction) (*TransactionResponse, error)
	GetContractInfo(context.Context, *Query) (*Response, error)
	ContractCallLocalMethod(context.Context, *Query) (*Response, error)
	ContractGetBytecode(context.Context, *Query) (*Response, error)
	GetBySolidityID(context.Context, *Query) (*Response, error)
	GetTxRecordByContractID(context.Context, *Query) (*Response, error)
	DeleteContract(context.Context, *Transaction) (*TransactionResponse, error)
	SystemDelete(context.Context, *Transaction) (*TransactionResponse, error)
	SystemUndelete(context.Context, *Transaction) (*TransactionResponse, error)
}

// UnimplementedSmartContractServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSmartContractServiceServer struct {
}

func (*UnimplementedSmartContractServiceServer) CreateContract(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContract not implemented")
}
func (*UnimplementedSmartContractServiceServer) UpdateContract(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContract not implemented")
}
func (*UnimplementedSmartContractServiceServer) ContractCallMethod(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractCallMethod not implemented")
}
func (*UnimplementedSmartContractServiceServer) GetContractInfo(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContractInfo not implemented")
}
func (*UnimplementedSmartContractServiceServer) ContractCallLocalMethod(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractCallLocalMethod not implemented")
}
func (*UnimplementedSmartContractServiceServer) ContractGetBytecode(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractGetBytecode not implemented")
}
func (*UnimplementedSmartContractServiceServer) GetBySolidityID(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySolidityID not implemented")
}
func (*UnimplementedSmartContractServiceServer) GetTxRecordByContractID(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxRecordByContractID not implemented")
}
func (*UnimplementedSmartContractServiceServer) DeleteContract(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContract not implemented")
}
func (*UnimplementedSmartContractServiceServer) SystemDelete(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDelete not implemented")
}
func (*UnimplementedSmartContractServiceServer) SystemUndelete(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUndelete not implemented")
}

func RegisterSmartContractServiceServer(s *grpc.Server, srv SmartContractServiceServer) {
	s.RegisterService(&_SmartContractService_serviceDesc, srv)
}

func _SmartContractService_CreateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).CreateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/CreateContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).CreateContract(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_UpdateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).UpdateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/UpdateContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).UpdateContract(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_ContractCallMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).ContractCallMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/ContractCallMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).ContractCallMethod(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_GetContractInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).GetContractInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/GetContractInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).GetContractInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_ContractCallLocalMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).ContractCallLocalMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/ContractCallLocalMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).ContractCallLocalMethod(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_ContractGetBytecode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).ContractGetBytecode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/ContractGetBytecode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).ContractGetBytecode(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_GetBySolidityID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).GetBySolidityID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/GetBySolidityID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).GetBySolidityID(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_GetTxRecordByContractID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).GetTxRecordByContractID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/GetTxRecordByContractID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).GetTxRecordByContractID(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_DeleteContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).DeleteContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/DeleteContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).DeleteContract(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_SystemDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).SystemDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/SystemDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).SystemDelete(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartContractService_SystemUndelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartContractServiceServer).SystemUndelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SmartContractService/SystemUndelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartContractServiceServer).SystemUndelete(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmartContractService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SmartContractService",
	HandlerType: (*SmartContractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createContract",
			Handler:    _SmartContractService_CreateContract_Handler,
		},
		{
			MethodName: "updateContract",
			Handler:    _SmartContractService_UpdateContract_Handler,
		},
		{
			MethodName: "contractCallMethod",
			Handler:    _SmartContractService_ContractCallMethod_Handler,
		},
		{
			MethodName: "getContractInfo",
			Handler:    _SmartContractService_GetContractInfo_Handler,
		},
		{
			MethodName: "contractCallLocalMethod",
			Handler:    _SmartContractService_ContractCallLocalMethod_Handler,
		},
		{
			MethodName: "ContractGetBytecode",
			Handler:    _SmartContractService_ContractGetBytecode_Handler,
		},
		{
			MethodName: "getBySolidityID",
			Handler:    _SmartContractService_GetBySolidityID_Handler,
		},
		{
			MethodName: "getTxRecordByContractID",
			Handler:    _SmartContractService_GetTxRecordByContractID_Handler,
		},
		{
			MethodName: "deleteContract",
			Handler:    _SmartContractService_DeleteContract_Handler,
		},
		{
			MethodName: "systemDelete",
			Handler:    _SmartContractService_SystemDelete_Handler,
		},
		{
			MethodName: "systemUndelete",
			Handler:    _SmartContractService_SystemUndelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/SmartContractService.proto",
}
