package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/elarsaks/Go-blockchain/cmd/wallet_server/handlers"
	"github.com/elarsaks/Go-blockchain/pkg/utils"
	"github.com/gorilla/mux"
)

type WalletServer struct {
	port    uint16
	gateway string
}

// Make sure WalletServer implements handlers.WalletServer
var _ handlers.WalletServer = &WalletServer{}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

// Run the WalletServer
func (ws *WalletServer) Run() {
	// Create router
	router := mux.NewRouter()
	router.Use(utils.CorsMiddleware())

	// Create an instance of WalletServerHandler
	handler := handlers.NewWalletServerHandler(ws)

	// Define routes
	router.HandleFunc("/", handler.GetApiDescription)
	// These methods need to be implemented as methods of WalletServerHandler
	router.HandleFunc("/user/wallet", handler.GetUserWallet)
	router.HandleFunc("/wallet/balance", handler.GetWalletBalance)
	router.HandleFunc("/transaction", handler.CreateTransaction)
	router.HandleFunc("/miner/blocks", handler.GetBlocks)
	// router.HandleFunc("/miner/wallet", ws.GetMinerWallet) // TODO

	// Start server
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), router))
}

func init() {
	log.SetPrefix("Wallet Server: ")
}

func main() {
	// Retrieve gateway from environment variable
	gateway := os.Getenv("WALLET_SERVER_GATEWAY_TO_BLOCKCHAIN")
	if gateway == "" {
		gateway = "http://miner-2:5002" // Default value
	}

	fmt.Println("gateway: ", gateway)

	// Retrieve port from environment variable
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		port = 5000 // Default value
	}

	// Print gateway and port
	log.Printf("Gateway to blockchain: %s\n", gateway)
	log.Printf("Port: %d\n", port)

	app := NewWalletServer(uint16(port), gateway)
	app.Run()
}
