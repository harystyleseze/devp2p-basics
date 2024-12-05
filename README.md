# DevP2P (Decentralized Peer-to-Peer Networking) in Go

This project demonstrates a basic peer-to-peer (P2P) network in Go where each peer connects to a bootnode to discover other peers and then communicates with them. The communication between peers is encrypted using **TLS** (Transport Layer Security), and each peer works as both a **client** and a **server**.

## Table of Contents
- [Overview](#overview)
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [Setup and Installation](#setup-and-installation)
- [Running the Project](#running-the-project)
- [Key Concepts](#key-concepts)
  - [Go Goroutines](#go-goroutines)
  - [Go Channels](#go-channels)
  - [TLS Encryption](#tls-encryption)
  - [Client-Server Communication](#client-server-communication)
- [Testing the Network](#testing-the-network)
- [License](#license)

---

## Overview

This project demonstrates a basic **decentralized P2P network** where:
- **Peers** (nodes) can connect to a **bootnode** to join the network and discover other peers.
- Each peer can act as both a **server** (listening for connections) and a **client** (connecting to other peers, including the bootnode).
- The communication between peers is secured with **TLS** encryption, using self-signed certificates.
- The bootnode acts as an entry point into the network, allowing peers to join and interact with each other.

### Components:
- **Main Peer**: This program runs both the client and server parts. It listens for incoming peer connections and connects to the bootnode to discover peers.
- **Bootnode**: A server that listens on a fixed port (`8080`) and provides its address to peers when they connect, allowing them to discover each other.
- **Transport Layer**: Contains functions to set up secure listeners and establish secure connections using TLS.

---

## Project Structure

```
/devp2p-basics
|-- main.go                     # Main peer program that acts as both client and server
|-- bootnode/bootnode.go        # Bootnode program that peers connect to for discovery
|-- transport/transport.go      # Functions for setting up secure communication (TLS)
|-- server-cert.pem             # Self-signed server certificate for TLS
|-- server-key.pem              # Self-signed private key for TLS
```

---

## Technologies Used

- **Go Programming Language**: For building the networking and concurrent parts of the application.
- **TLS (Transport Layer Security)**: For secure peer-to-peer communication.
- **Goroutines**: For concurrent execution of server and client operations in the same peer.
- **Go Channels**: Used for synchronization and communication between different goroutines.

---

## Setup and Installation

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/yourusername/devp2p-basics.git
   cd devp2p-basics
   ```

2. Ensure that you have Go installed. If not, you can install it from the official site: https://golang.org/dl/.

3. Create the self-signed certificates (`server-cert.pem` and `server-key.pem`) in the project directory:
   - You can use OpenSSL to generate a self-signed certificate and private key:
     ```bash
     openssl req -x509 -newkey rsa:4096 -keyout server-key.pem -out server-cert.pem -days 365
     ```
   - When prompted, provide the necessary information, and make sure to leave the password for the private key empty.

---

## Running the Project

1. Start the **bootnode**:
   - In one terminal window, run the bootnode program:
     ```bash
     go run bootnode.go
     ```
   - This will start the bootnode, which listens on `localhost:8080` for incoming connections from peers.

2. Start the **main peer**:
   - In another terminal window, run the main peer program:
     ```bash
     go run main.go
     ```
   - This peer will connect to the bootnode at `localhost:8080` and establish secure communication.

3. You should see output indicating that the main peer has successfully connected to the bootnode, and the peer-server interactions will take place.

---

## Key Concepts

### Go Goroutines

- **Goroutines** are lightweight concurrent threads of execution in Go. They allow us to execute multiple functions concurrently without the need for complex thread management.
- In the code, we use Goroutines to handle both the server and client parts of the peer concurrently:
  - The main peer starts a Goroutine to listen for incoming peer connections and act as a server.
  - Another Goroutine is used to handle the client part, connecting to the bootnode and sending messages.

Example from the code:
```go
go func() {
    // Client logic here
}()
```

This allows the program to execute both tasks (client and server) at the same time, without blocking the other tasks.

### Go Channels

- **Channels** in Go are used to communicate between Goroutines. They allow us to send and receive values safely between different parts of the program.
- In this project, a **sync.WaitGroup** is used to wait for all Goroutines to complete before the program terminates, ensuring that all tasks (server and client) finish their execution.

Example from the code:
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Handle server/client task
}()
wg.Wait()
```
This ensures that the main function waits for both the server and client Goroutines to complete before terminating.

### TLS Encryption

- **TLS** is used to encrypt the communication between peers, providing security and privacy.
- The peer-to-peer communication is encrypted using the self-signed certificates (`server-cert.pem` and `server-key.pem`). The `tls` package in Go is used to configure secure listeners and clients.
- **InsecureSkipVerify** is set to `true` for simplicity in this demo, but in real-world scenarios, certificate validation should be done to ensure secure and trusted communication.

Example from `transport.go`:
```go
config := &tls.Config{
    InsecureSkipVerify: true, // Not recommended for production
}
```

### Client-Server Communication

- Each peer acts as both a server (listening for incoming connections) and a client (initiating connections to other peers or the bootnode).
- The **server** listens for incoming connections and handles them by accepting connections, reading messages, and sending responses.
- The **client** connects to the bootnode to join the network, retrieves the bootnode’s address, and sends messages to the network.

Example of server communication in `main.go`:
```go
conn, err := listener.Accept()
if err != nil {
    log.Fatal("Error accepting connection:", err)
}
```

Example of client communication:
```go
conn, err := transport.ConnectToPeer(fmt.Sprintf("localhost:%s", bootnodePort))
```

---

## Testing the Network

To test the peer-to-peer network:

1. Run the **bootnode** in one terminal.
2. Run multiple instances of the **main peer** in other terminals.
3. Each peer will connect to the bootnode, exchange messages, and print the received data.

---

## License

This project is licensed under the MIT License.