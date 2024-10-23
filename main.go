package main

import (
	"flag"
	"fmt"
	"github.com/zhulida1234/go-rpc-service/config"
	services "github.com/zhulida1234/go-rpc-service/server"
	"strconv"
)

func main() {
	var f = flag.String("c", "config.yml", "config path")
	flag.Parse()
	conf, err := config.New(*f)
	if err != nil {
		fmt.Println("failed to load config", "err", err)
		return
	}
	port, err := strconv.Atoi(conf.Server.Port)
	if err != nil {
		return
	}
	grpcServerCfg := &services.RpcServerConfig{
		GrpcHostname: conf.Server.Host,
		GrpcPort:     port,
	}

	rpcServer, err := services.NewRpcServer(grpcServerCfg)
	if err != nil {
		fmt.Println("error create rpc server")
		return
	}
	err = rpcServer.Start()
	if err != nil {
		fmt.Println("error create rpc server", "err", err)
		return
	}

	<-(chan struct{})(nil)
}
