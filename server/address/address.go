package addresss

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthAddress struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
}

func CreateAddressFromPrivateKey() (*EthAddress, error) {
	prvKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	address := &EthAddress{
		PrivateKey: hex.EncodeToString(crypto.FromECDSA(prvKey)),
		PublicKey:  hex.EncodeToString(crypto.FromECDSAPub(&prvKey.PublicKey)),
		Address:    crypto.PubkeyToAddress(prvKey.PublicKey).String(),
	}
	return address, nil
}

func PublicKeyToAddress(publicKey string) (string, error) {
	publicKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		return "", err
	}
	addressCommon := common.BytesToAddress(crypto.Keccak256(publicKeyBytes[1:])[12:])
	return addressCommon.String(), nil
}
