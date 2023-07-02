
# Go-blockchain

## About
This Docker-based blockchain application consists of the following components:

- **Client UI (React Dashboard):**  
A user interface built using React, providing a visual dashboard for interacting with the blockchain.  

- **API Gateway (Wallet Server):**  
A Golang web server serving as a gateway between the client UI and the blockchain.  
It handles requests and forwards them to the appropriate APIs.  

- **Blockchain Miner Nodes (Golang APIs):**  
Three Golang APIs that serve as blockchain miner nodes.  
They perform mining operations and maintain the integrity of the blockchain.

With this setup, users can access the client UI to interact with the blockchain through the API gateway, which communicates with the blockchain miner nodes to process transactions and maintain the blockchain's distributed ledger.

The Docker configuration ensures easy deployment and scalability of the blockchain application.


## Dependencies
- Docker & Docker Compose
- Golang
- Air [GitHub repository](https://github.com/cosmtrek/air)
- Node v17

## Installation & Running

1. To install Golang dependencies, at the root of this project folder run:

```bash
go mod tidy
```

2. To install JavaScript dependencies, at the /react_dashboard folder run:

```bash
npm install
```

3. To run it in the docker, at the root of this project folder run:
```bash
docker-compose up --build
```

### Apps will start on ports:
| App              | URL                                 |
|------------------|-------------------------------------|
| react_dashboard | [http://localhost:3000](http://localhost:3000) |
| wallet_server   | [http://localhost:5000](http://localhost:5000) |
| miner_1         | [http://localhost:5001](http://localhost:5001) |
| miner_2         | [http://localhost:5002](http://localhost:5002) |
| miner_3         | [http://localhost:5003](http://localhost:5003) |

**Note:** To run each apps separately, check for ReadMe files in each app folder:
- `react_dashboard/ReadMe.md`
- `wallet_server/ReadMe.md`
- `blockchain_server/Readme.md`


### TODO:
- Connect Miner Wallet to API
- Connect Reguar Wallet to API
- Implement sending crypto from miner to regular wallet


