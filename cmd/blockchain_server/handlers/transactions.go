package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/elarsaks/Go-blockchain/pkg/block"
	"github.com/elarsaks/Go-blockchain/pkg/utils"
)

func (h *BlockchainServerHandler) HandleGetTransaction(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	bc := h.server.GetBlockchain()

	transactions := bc.TransactionPool()

	m, _ := json.Marshal(struct {
		Transactions []*block.Transaction `json:"transactions"`
		Length       int                  `json:"length"`
	}{
		Transactions: transactions,
		Length:       len(transactions),
	})

	io.WriteString(w, string(m[:]))
}

func (h *BlockchainServerHandler) HandlePostTransaction(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t block.TransactionRequest
	err := decoder.Decode(&t)

	if err != nil {
		log.Printf("ERROR: %v", err)
		io.WriteString(w, string(utils.JsonStatus("fail")))
		return
	}

	if !t.Validate() {
		log.Println("ERROR: missing field(s)")
		io.WriteString(w, string(utils.JsonStatus("fail")))
		return
	}

	publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
	signature := utils.SignatureFromString(*t.Signature)

	bc := h.server.GetBlockchain()

	isCreated, err := bc.CreateTransaction(*t.SenderBlockchainAddress,
		*t.RecipientBlockchainAddress, *t.Message, *t.Value, publicKey, signature)

	w.Header().Add("Content-Type", "application/json")

	var m []byte
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "fail",
			Message: err.Error(),
		}

		m, _ = json.Marshal(errMsg)
	} else if !isCreated {
		w.WriteHeader(http.StatusBadRequest)
		m = utils.JsonStatus("fail")
	} else {
		w.WriteHeader(http.StatusCreated)
		m = utils.JsonStatus("success")
	}

	io.WriteString(w, string(m))
}

func (h *BlockchainServerHandler) HandlePutTransaction(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t block.TransactionRequest
	err := decoder.Decode(&t)

	if err != nil {
		log.Printf("ERROR: %v", err)
		io.WriteString(w, string(utils.JsonStatus("fail")))
		return
	}

	if !t.Validate() {
		log.Println("ERROR: missing field(s)")
		io.WriteString(w, string(utils.JsonStatus("fail")))
		return
	}

	publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
	signature := utils.SignatureFromString(*t.Signature)

	bc := h.server.GetBlockchain()

	isUpdated, err := bc.AddTransaction(*t.SenderBlockchainAddress,
		*t.RecipientBlockchainAddress,
		*t.Message,
		*t.Value,
		publicKey,
		signature)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "fail", "error": err.Error()})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}

	w.Header().Add("Content-Type", "application/json")

	var m []byte
	if !isUpdated {
		w.WriteHeader(http.StatusBadRequest)
		m = utils.JsonStatus("fail")
	} else {
		m = utils.JsonStatus("success")
	}

	io.WriteString(w, string(m))
}

func (h *BlockchainServerHandler) HandleDeleteTransaction(w http.ResponseWriter, req *http.Request) {
	bc := h.server.GetBlockchain()

	bc.ClearTransactionPool()

	io.WriteString(w, string(utils.JsonStatus("success")))
}

func (h *BlockchainServerHandler) Transactions(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.HandleGetTransaction(w, req)
	case http.MethodPost:
		h.HandlePostTransaction(w, req)
	case http.MethodPut:
		h.HandlePutTransaction(w, req)
	case http.MethodDelete:
		h.HandleDeleteTransaction(w, req)
	default:
		log.Println("ERROR: Invalid HTTP Method")
		w.WriteHeader(http.StatusBadRequest)
	}
}
