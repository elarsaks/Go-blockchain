package handlers

import (
	"fmt"
	"reflect"

	"github.com/elarsaks/Go-blockchain/pkg/wallet"
)

func LogMethods(i interface{}) {
	t := reflect.TypeOf(i)

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println(method.Name)
	}
}

type BlockchainServer interface {
	Port() uint16
	GetWallet() *wallet.Wallet
}

type BlockchainServerHandler struct {
	server BlockchainServer
}

func NewBlockchainServerHandler(s BlockchainServer) *BlockchainServerHandler {
	LogMethods(s)
	return &BlockchainServerHandler{server: s}
}
