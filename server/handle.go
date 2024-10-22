package server

import (
	"context"
	"fmt"
	"github.com/zhulida1234/go-rpc-service.git/protobuf/wallet"
	addresss "github.com/zhulida1234/go-rpc-service.git/server/address"
	"strconv"
)

func (s *RpcServer) GetSupportCoins(ctx context.Context, in *wallet.SupportCoinsRequest) (*wallet.SupportCoinsResponse, error) {
	return &wallet.SupportCoinsResponse{
		Code:    strconv.Itoa(200),
		Msg:     "success request",
		Support: true,
	}, nil
}

func (s *RpcServer) GetWalletAddress(ctx context.Context, in *wallet.WalletAddressRequest) (*wallet.WalletAddressResponse, error) {
	addressInfo, err := addresss.CreateAddressFromPrivateKey()
	if err != nil {
		fmt.Println("err create address")
		return &wallet.WalletAddressResponse{
			Code:      strconv.Itoa(400),
			Msg:       "create address fail",
			Address:   "",
			PublicKey: "",
		}, nil
	}
	return &wallet.WalletAddressResponse{
		Code:      strconv.Itoa(200),
		Msg:       "success request",
		Address:   addressInfo.Address,
		PublicKey: addressInfo.PublicKey,
	}, nil
}
