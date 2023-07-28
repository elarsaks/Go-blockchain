package wallet_server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/elarsaks/Go-blockchain/pkg/utils"
	"github.com/gorilla/mux"
)

type WalletServer struct {
	port    uint16
	gateway string
}

// Create a new instance of WalletServer
func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

// Get the port of the WalletServer
func (ws *WalletServer) Port() uint16 {
	return ws.port
}

// Get the gateway of the WalletServer
func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

// Run the WalletServer
func (ws *WalletServer) Run() {
	// Create router
	router := mux.NewRouter()
	router.Use(utils.CorsMiddleware())

	// Define routes
	router.HandleFunc("/wallet", ws.GetUserWallet)
	router.HandleFunc("/wallet/balance", ws.WalletBalance)
	router.HandleFunc("/transaction", ws.CreateTransaction)

	// Start server
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), router))
}
