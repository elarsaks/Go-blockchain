# Go-blockchain

# To run wallets
IN: \Go-blockchain\wallet_server>
    - Run wallet 1: go run main.go wallet_server.go -port 8080 -gateway http://127.0.0.1:5001
    - Run wallet 2: go run main.go wallet_server.go -port 8081 -gateway http://127.0.0.1:5001

# To run miners
IN: \Go-blockchain\blockchain_server>
    - Run miner 1: go run main.go blockchain_server.go -port 5001
    - Run miner 2: go run main.go blockchain_server.go -port 5002
    - Run miner 3: go run main.go blockchain_server.go -port 5003

# To run dashboard
IN: \Go-blockchain\react_dashboard>
    - npm start