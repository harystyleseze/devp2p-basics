package main

import (
	"devp2p-basics/transport"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	bootnodePort := "8080" // Bootnode's port (initial point for peers)
	port := "8081"         // Peer's port

	// Start the server (this instance will listen on PORT)
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Start the listener to accept incoming connections
		listener, err := transport.CreateSecureListener(fmt.Sprintf("localhost:%s", port))
		if err != nil {
			log.Fatal("Error creating server:", err)
		}
		defer listener.Close()

		fmt.Printf("Listening on port %s\n", port)

		// Accept the incoming connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection:", err)
		}
		defer conn.Close()

		// Print message when connection is accepted
		fmt.Println("Connection accepted from", conn.RemoteAddr())

		// Read message from the peer
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatal("Failed to read message:", err)
		}
		fmt.Printf("Received message: %s\n", string(buffer[:n]))

		// Respond back to the peer
		_, err = conn.Write([]byte("Hello from the server"))
		if err != nil {
			log.Fatal("Failed to send response:", err)
		}
	}()

	// Simulate a delay to allow the server to start listening
	time.Sleep(1 * time.Second)

	// Start the client to connect to the bootnode
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Connect to the bootnode to join the network
		conn, err := transport.ConnectToPeer(fmt.Sprintf("localhost:%s", bootnodePort))
		if err != nil {
			log.Fatal("Error connecting to bootnode:", err)
		}
		defer conn.Close()

		fmt.Println("Connected to bootnode at", bootnodePort)

		// Read the bootnode's address (or list of peers)
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatal("Failed to read response from bootnode:", err)
		}
		bootnodeAddress := string(buffer[:n])
		fmt.Printf("Received bootnode address: %s\n", bootnodeAddress)

		// Now that we have the bootnode's address, connect to it
		conn, err = transport.ConnectToPeer(bootnodeAddress)
		if err != nil {
			log.Fatal("Error connecting to bootnode:", err)
		}
		defer conn.Close()

		// Send a message to the bootnode
		_, err = conn.Write([]byte("Hello world from client"))
		if err != nil {
			log.Fatal("Failed to send message:", err)
		}
		fmt.Println("Message sent to bootnode")

		// Read the response from the bootnode
		n, err = conn.Read(buffer)
		if err != nil {
			log.Fatal("Failed to read response:", err)
		}
		fmt.Printf("Received response: %s\n", string(buffer[:n]))
	}()

	// Wait for both the server and client to complete
	wg.Wait()
}
