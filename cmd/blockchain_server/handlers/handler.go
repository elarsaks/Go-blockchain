package handlers

import "github.com/elarsaks/Go-blockchain/pkg/wallet"

type BlockchainServer struct {
	Port   uint16
	Wallet *wallet.Wallet
	//* NOTE: In real world app we would not attach the wallet to the server
	// But for the sake of simplicity we will do it here,
	// because we dont store miners credentials in a database.
}

type BlockchainServerHandler struct {
	server BlockchainServer
}

func NewWalletServerHandler(s BlockchainServer) *BlockchainServerHandler {
	return &BlockchainServerHandler{server: s}
}
