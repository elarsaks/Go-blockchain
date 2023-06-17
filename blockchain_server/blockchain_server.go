package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/elarsaks/Go-blockchain/block"
	"github.com/elarsaks/Go-blockchain/utils"
	"github.com/elarsaks/Go-blockchain/wallet"
)

// cache stores the blockchain instance
var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

// BlockchainServer represents a server that handles blockchain requests.
type BlockchainServer struct {
	port uint16
}

// NewBlockchainServer creates a new instance of BlockchainServer with the specified port.
func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

// Port returns the port on which the server is running.
func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

// GetBlockchain returns the blockchain associated with the server.
// If the blockchain doesn't exist in the cache, it creates a new one.
func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]

	if !ok {
		minersWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minersWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc
		log.Printf("private_key: %v", minersWallet.PrivateKeyStr())
		log.Printf("public_key: %v", minersWallet.PublicKeyStr())
		log.Printf("blockchain_address: %v", minersWallet.BlockchainAddress())
	}

	return bc
}

// GetChain handles the HTTP GET request for retrieving the blockchain.
func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("Error: Invalid HTTP Method")
	}
}

func (bcs *BlockchainServer) Transactions(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		transactions := bc.TransactionPool()
		m, _ := json.Marshal(struct {
			transactions []*block.Transaction `json:"transactions"`
			Length       int                  `json:"length"`
		}{
			Transactions: transactions,
			Length:       len(transactions),
		})
		io.WriteString(w, string(m))

	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t *wallet.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("Error: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if !t.Validate() {
			log.Printf("Error: missing field(s)")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		signature := utils.SignatureFromStrings(*t.Signature)
		bc := bcs.GetBlockchain()
		isCreated := bc.CreateTransaction(
			publicKey,
			*t.SenderBlockchainAddress,
			*t.RecipientBlockchainAddress,
			*t.Value,
			signature)

		w.Header().Add("Content-Type", "application/json")
		var m []byte

		if !isCreated {
			w.WriteHeader(http.StatusBadRequest)
			m, _ = utils.JsonStatus("fail")
		} else {
			w.WriteHeader(http.StatusCreated)
			m, _ = utils.JsonStatus("success")
		}

		io.WriteString(w, string(m))
	default:
		log.Printf("Error: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}

}

// Run starts the server and listens for incoming HTTP requests.
func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	http.HandleFunc("/transactions", bcs.Transactions)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
