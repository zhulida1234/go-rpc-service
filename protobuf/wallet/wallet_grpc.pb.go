package wallet

import (
	"context"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion7

const (
	WalletService_GetSupportCoins_fullMethodName  = "/the_web_three.wallet.WalletService/GetSupportCoins"
	WalletService_GetWalletAddress_fullMethodName = "/the_web_three.wallet.WalletService/getWalletAddress"
)

type WalletServiceClient interface {
	GetSupportCoins(ctx context.Context, in *SupportCoinsRequest, opts ...grpc.CallOption) (*SupportCoinsResponse, error)
	GetWalletAddress(ctx context.Context, in *WalletAddressRequest, opts ...grpc.CallOption) (*WalletAddressResponse, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) GetSupportCoins(ctx context.Context, in *SupportCoinsRequest, opts ...grpc.CallOption) (*SupportCoinsResponse, error) {
	out := new(SupportCoinsResponse)
	err := c.cc.Invoke(ctx, "/the_web_three.wallet.WalletService/GetSupportCoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, err
}

func (c *walletServiceClient) GetWalletAddress(ctx context.Context, in *WalletAddressRequest, opts ...grpc.CallOption) (*WalletAddressResponse, error) {
	out := new(WalletAddressResponse)
	err := c.cc.Invoke(ctx, "/the_web_three.wallet.WalletService/GetWalletAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type UnimplementedWalletServiceServer struct{}

func (UnimplementedWalletServiceServer) GetSupportCoins(ctx context.Context, in *SupportCoinsRequest, opts ...grpc.CallOption) (*SupportCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportCoins not implemented")
}

func (UnimplementedWalletServiceServer) GetWalletAddress(ctx context.Context, in *WalletAddressRequest, opts ...grpc.CallOption) (*WalletAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletAddress not implemented")
}

type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	s.RegisterService(&WalletService_ServceDesc, srv)
}

func _WalletService_GetSupportCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetSupportCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_GetSupportCoins_fullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetSupportCoins(ctx, req.(*SupportCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetWalletAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetWalletAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_GetWalletAddress_fullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetWalletAddress(ctx, req.(*WalletAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "the_web_three.wallet.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSupportCoins",
			Handler:    _WalletService_GetSupportCoins_Handler,
		},
		{
			MethodName: "GetWalletAddress",
			Handler:    _WalletService_GetWalletAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}
