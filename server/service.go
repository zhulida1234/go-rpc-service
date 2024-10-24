package services

import (
	"context"
	"fmt"
	"github.com/zhulida1234/go-rpc-service/database"
	"github.com/zhulida1234/go-rpc-service/protobuf/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync/atomic"
)

const MaxRecvMessageSize = 1024 * 1024 * 300

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     int
}

type RpcServer struct {
	*RpcServerConfig
	db *database.DB

	wallet.UnimplementedWalletServiceServer

	stopped atomic.Bool
}

func (s *RpcServer) Stop(ctx context.Context) error {
	s.stopped.Store(true)
	return nil
}

func (s *RpcServer) Stopped() bool {
	//TODO implement me
	panic("implement me")
}

func NewRpcServer(db *database.DB, config *RpcServerConfig) (*RpcServer, error) {
	return &RpcServer{
		RpcServerConfig: config,
		db:              db,
	}, nil
}

func (s *RpcServer) Start(ctx context.Context) error {
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
