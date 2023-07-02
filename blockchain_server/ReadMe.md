# MINER - Blockchain Server

## About
Miner is a web server that facilitates the process of verifying and adding new blocks to the blockchain network. In the world of blockchain, it is commonly referred to as a node.

## Dependenies
- Golang
- Air

## Installation

Follow the steps below to set up the Blockchain Server:

1. Navigate to the parent folder of the project in your terminal.

2. Run the following command to download the necessary dependencies:

```bash
go mod tidy
```

## Running
To run the app with the Air library (live reloading), execute the following command in this folder:
```bash
PORT=5001 air
```

To run it directly via Golang, execute the following command in this folder:
```bash
go run main.go wallet_server.go -port 8080 -gateway http://127.0.0.1:5001
```

**Note:**  
These commands run only one miner. If you want to run multiple miners, open multiple new terminals and run the same command with different port numbers.

Feel free to adjust the port numbers as needed to run multiple miners concurrently.


