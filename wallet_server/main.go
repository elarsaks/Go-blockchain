package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Wallet Server: ")
}

func main() {
	port := flag.Uint("port", 8080, "TCP Port Number for Wallet Server")
	gateway := flag.String("gateway", "http://127.0.0.1:5001", "Blockchain Server Gateway")
	flag.Parse()

	app := NewWalletServer(uint16(*port), *gateway)
	app.Run()
}
