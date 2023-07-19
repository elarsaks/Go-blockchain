# WALLET - Wallet Server

## About
The Wallet Server is a web server that handles the connection between clients and the blockchain network.

## Dependenies
- Golang
- Air

## Installation

If you haven't already installed the project from the parent folder, follow these steps to set up the Wallet Server:

1. Navigate to the parent folder of the project in your terminal.

2. Run the following command to download the necessary dependencies:

```bash
go mod tidy
```

## Running
To run the app with the Air library (live reloading), execute the following command in this folder:
```bash
PORT=5000 HOST=127.0.0.1 GATEWAY_PORT=5001 air
```

To run it directly via Golang, execute the following command in this folder:
```bash
go run main.go
```



