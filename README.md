
# Go-blockchain (🚧 UNDER DEVELOPMENT 🚧)
- [*LIVE EXAMPLE*](https://elarsaks.github.io/Go-blockchain/)
- [*PROJECT WIKI*](https://github.com/elarsaks/Go-blockchain/wiki)

# Table of Contents
1. [About](#about)
2. [Workflow](#workflow)
3. [Components](#components)
4. [Installation and Running](#installation-and-running)
   - [Run All in Docker](#run-all-in-docker)
   - [Run React Dashboard as Standalone](#run-react-dashboard-as-standalone)
   - [Run Wallet Server as Standalone](#run-wallet-server-as-standalone)
   - [Run Miner / Node as Standalone](#run-miner--node-as-standalone)

# About  
This project is a Docker-based blockchain application that is currently under development. It is written in Golang and features a user interface built in React. The application is composed of several key components, each serving a unique role in the overall functionality of the system.

**This project serves two main purposes:**

1. **Skill Development:** It provides an excellent opportunity to push the boundaries of my technical skills and deepen my understanding of Blockchain technology.

2. **Community Resource:** I'm dedicated to creating a robust codebase that can serve as a valuable learning resource for anyone interested in this technology.


![Topology Diagram](https://github.com/elarsaks/Go-blockchain/blob/main/docs/topology.png)


# Workflow
Users can access the client UI to interact with the blockchain. The API gateway processes these interactions and communicates with the blockchain miner nodes. These nodes then process transactions and maintain the distributed ledger of the blockchain.

**The Docker configuration of this application ensures easy deployment and scalability, making it a robust and flexible solution for blockchain applications.**


# Components
- **Client UI (React Dashboard):** The client user interface is a visual dashboard built using React. This dashboard serves as the primary point of interaction for users with the blockchain. It provides a user-friendly interface for executing various operations on the blockchain.

- **API Gateway (Wallet Server):** The API Gateway, also known as the Wallet Server, is a web server developed in Golang. It acts as a bridge between the client UI and the blockchain, managing requests from the client and forwarding them to the appropriate APIs.

- **Blockchain Miner Nodes (Golang APIs):** The backbone of the blockchain is formed by three Golang APIs that function as blockchain miner nodes. These nodes are responsible for performing mining operations and maintaining the integrity of the blockchain.

<img src="https://saks.digital/wp-content/uploads/2023/07/some.png" alt="Image Description" />

# Installation and running

## Run all in Docker
**Dependencies:**  
- Docker & Docker Compose  

**Running:**  
- To run full-stack application in docker, at the root of this project folder run:
```bash
docker-compose up --build
```

**Apps will start on ports:**
| App              | URL                                 |
|------------------|-------------------------------------|
| react_dashboard | [http://localhost:3000](http://localhost:3000) |
| wallet_server   | [http://localhost:5000](http://localhost:5000) |
| miner_1         | [http://localhost:5001](http://localhost:5001) |
| miner_2         | [http://localhost:5002](http://localhost:5002) |
| miner_3         | [http://localhost:5003](http://localhost:5003) |

---
<br></br>
## Run React Dashboard as standalone 
**Dependenies:**    
- Node v17  

**Installation:**  
To set up the UI (client), run the following command in this *Go-Blockchain/react_dashboard/*:
```bash
npm install
```

**Running:**  
To run the UI, use the following command in this folder:
```bash
npm start
```

**App will start on port:**
| App              | URL                                 |
|------------------|-------------------------------------|
| react_dashboard | [http://localhost:3000](http://localhost:3000) |


---
<br></br>
## Run Wallet Server as standalone
**Dependenies:**  
- Golang
- Air  

**Installation:**  
If you haven't already installed the project from the parent folder, follow these steps to set up the Wallet Server:
1. Navigate to the root folder of the project in your terminal.
2. Run the following command to download the necessary dependencies:
```bash
go mod tidy
```

**Running:**  
To run the app with the Air library (live reloading), execute the following command in *Go-Blockchain/wallet_server/*:
```bash
PORT=8081 HOST=127.0.0.1 GATEWAY_PORT=5001 air
```
To run it directly via Golang, execute the following command in this folder:
```bash
go run main.go blockchain_server.go -port 5001
```
**App will start on port:**
| App              | URL                                 |
|------------------|-------------------------------------|
| wallet_server | [http://localhost:3000](http://localhost:3000) |

---
<br></br>
## Run Miner / Node as standalone
**Dependenies:**  
- Golang
- Air  

**Installation:**  
1. Navigate to the root folder of the project in your terminal.
2. Run the following command to download the necessary dependencies:
```bash
go mod tidy
```

**Running:**  
To run the app with the Air library (live reloading), execute the following command in *Go-Blockchain/blockshain_server/*:
```bash
PORT=5001 air
```

**OR:**  

To run it directly via Golang, execute the following command in *Go-Blockchain/blockshain_server/*:
```bash
go run main.go wallet_server.go -port 8080 -gateway http://127.0.0.1:5001
```

**App will start on port:**
| App              | URL                                 |
|------------------|-------------------------------------|
| blockshain_server | [http://localhost:3000](http://localhost:3000) |

**NOTE:**  
**These commands run only one miner. If you want to run multiple miners, open multiple new terminals and run the same command with different port numbers.**

**Feel free to adjust the port numbers as needed to run multiple miners concurrently.**




