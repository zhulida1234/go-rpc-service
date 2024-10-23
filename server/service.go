package services

import (
	"fmt"
	"github.com/zhulida1234/go-rpc-service/protobuf/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const MaxRecvMessageSize = 1024 * 1024 * 300

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     int
}

type RpcServer struct {
	*RpcServerConfig
	wallet.UnimplementedWalletServiceServer
}

func NewRpcServer(config *RpcServerConfig) (*RpcServer, error) {
	return &RpcServer{
		RpcServerConfig: config,
	}, nil
}

func (s *RpcServer) Start() error {
	go func(s *RpcServer) {
		addr := fmt.Sprintf("%s:%d", s.GrpcHostname, s.GrpcPort)
		fmt.Println("start rpc server", "addr", addr)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			fmt.Println("Could not start rpc server", "err", err)
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)

		gs := grpc.NewServer(opt, grpc.ChainUnaryInterceptor(nil))
		reflection.Register(gs)

		wallet.RegisterWalletServiceServer(gs, s)

		fmt.Println("start rpc server", "port", s.GrpcPort, "address", listener.Addr())
		if err := gs.Serve(listener); err != nil {
			fmt.Println("start rpc server", "err", err)
		}
	}(s)
	return nil
}
