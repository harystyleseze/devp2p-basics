package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// Bootnode will listen on port 8080 using TLS encryption
	port := "8080"

	// Load the server's TLS certificate and key
	cert, err := tls.LoadX509KeyPair("/home/harystyles/go/src/networking/devp2p-basics/server-cert.pem", "/home/harystyles/go/src/networking/devp2p-basics/server-key.pem")
	if err != nil {
		log.Fatal("Error loading certificate:", err)
	}

	// Configure the TLS settings
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		// InsecureSkipVerify: true, // We could use this in real-world scenarios, but for simplicity, we'll keep verification.
	}

	// Create a TLS listener for the bootnode
	listener, err := tls.Listen("tcp", "localhost:"+port, config)
	if err != nil {
		log.Fatal("Error creating bootnode listener:", err)
	}
	defer listener.Close()

	fmt.Printf("Bootnode listening on port %s\n", port)

	// Accept connections from peers
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("New peer connected from %s\n", conn.RemoteAddr())

	// Send the bootnode's address to the new peer
	peerAddress := "localhost:8080" // Provide the bootnode's address (or other peers)
	_, err := conn.Write([]byte(peerAddress))
	if err != nil {
		log.Println("Error sending peer address:", err)
		return
	}

	// Simulate a handshake or waiting for the new peer to get initialized
	time.Sleep(1 * time.Second)
	fmt.Println("Handshake complete with new peer")
}
