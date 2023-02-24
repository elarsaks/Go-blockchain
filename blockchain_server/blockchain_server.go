package main

import (
	"Go-blockchain/block"
	"log"
	"net/http"
	"strconv"
)

var cache map[string]*block.Blockchain = make(map[string]*Block.Blockchain)

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func (bcs *BlockchainServer) Blockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]

	if !ok {
		minersWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minersWallet.Address(), bcs.Port())
		cache["blockchain"] = bc
		log.Printf("private_key: %v", minersWallet.PrivateKeyStr())
		log.Printf("public_key: %v", minersWallet.PublicKeyStr())
		log.Printf("blockchain_address: %v", minersWallet.BlockchainAddress())
	}

	return bc
}

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, r *http.Request) {
	switch req.Method {
	case "GET":
		bc := bcs.Blockchain()
	
	default: 
		log.Printf("Error: Invalid request method: %v", req.Method)

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
