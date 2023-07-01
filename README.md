# Go-blockchain

## WALLETS

### To run

IN: \Go-blockchain\wallet_server>

- Run wallet 1: go run main.go wallet_server.go -port 8080 -gateway http://127.0.0.1:5001
- Run wallet 2: go run main.go wallet_server.go -port 8081 -gateway http://127.0.0.1:5001

### TODO:

- Start using Mux Router
- Add live reloading, such as 'air'

## MINERS

### To run

IN: \Go-blockchain\blockchain_server>

- Run miner 1: go run main.go blockchain_server.go -port 5001
- Run miner 2: go run main.go blockchain_server.go -port 5002
- Run miner 3: go run main.go blockchain_server.go -port 5003

### TODO:

- Add live reloading, such as 'air'

## DASHBOARD

### To run

IN: \Go-blockchain\react_dashboard> - npm start
It starts in port 3000

### TODO:

- Connect Miner Wallet to API
- Connect Reguar Wallet to API
- Implement sending crypto from miner to regular wallet

## GLOBAL TODO:

- Containerize applications
