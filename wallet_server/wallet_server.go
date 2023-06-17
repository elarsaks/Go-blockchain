package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"

	"github.com/elarsaks/Go-blockchain/block"
	"github.com/elarsaks/Go-blockchain/utils"
	"github.com/elarsaks/Go-blockchain/wallet"
)

const tempDir = "templates/"

// WalletServer represents a server that serves a wallet application.
type WalletServer struct {
	port    uint16
	gateway string
}

// NewWalletServer creates a new instance of WalletServer with the specified port and gateway.
func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

// Port returns the port on which the server is running.
func (ws *WalletServer) Port() uint16 {
	return ws.port
}

// Gateway returns the gateway address used by the server.
func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

// Index handles the HTTP GET request for the index page.
func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(tempDir, "index.html"))
		t.Execute(w, ws)
	default:
		log.Printf("Error: Invalid HTTP Method")
	}
}

// Returns starting data for the wallet page
func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error: Invalid HTTP Method")
	}
}

// CreateTransaction handles the HTTP POST request for creating a transaction.
func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("Error: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
		}
		if !t.Validate() {
			log.Printf("Error: Missing required fields")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		publicKey := utils.PublicKeyFromString(t.SenderPublicKey)
		privateKey := utils.PrivateKeyFromString(t.SenderPrivateKey, publicKey)
		value, err := strconv.ParseFloat(t.Value, 32)

		if err != nil {
			log.Println("Error: parse error")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		value32 := float32(value)

		w.Header().Add("Content-Type", "application/json")

		transaction := wallet.NewTransaction(
			privateKey,
			publicKey,
			*t.SenderBlockchainAddress,
			*t.RecipientBlockchainAddress,
			value32)

		signature := transaction.GenerateSignature()
		signatureStr := signature.String()

		bt := &block.TransactionRequest{
			t.SenderBlockchainAddress,
			t.RecipientBlockchainAddress,
			t.SenderPublicKey,
			&value32,
			&signatureStr,
		}

		m, _ := json.Marshal(bt)
		buf := bytes.NewBuffer(m)
		resp, err := http.Post(ws.Gateway()+"/transaction", "application/json", buf)

		if resp.StatusCode == 201 {
			io.WriteString(w, string(utils.JsonStatus("success")))
			return
		}

		io.WriteString(w, string(utils.JsonStatus("fail")))

	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error: Invalid HTTP Method")
	}
}

// Run starts the server and listens for incoming HTTP requests.
func (ws *WalletServer) Run() {
	fmt.Printf("Wallet Server Listening on Port %d\n", ws.Port())
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
