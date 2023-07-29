package handlers

import (
	blockchain_server "github.com/elarsaks/Go-blockchain/pkg/blockchain_server"
)

type BlockchainServerHandler struct {
	server *blockchain_server.BlockchainServer
}

func NewBlockchainServerHandler(s *blockchain_server.BlockchainServer) *BlockchainServerHandler {
	return &BlockchainServerHandler{server: s}
}
