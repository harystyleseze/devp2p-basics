package main

import (
	"devp2p-basics/transport"
	"fmt"
	"io"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var PORT = 8080

	// Start the client connection in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done() // Notify when the client finishes

		conn, err := transport.ConnectToPeer("localhost:" + fmt.Sprint(PORT))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		fmt.Println("Connected to peer")

		// Send a simple message to the server
		_, err = conn.Write([]byte("Hello world from client"))
		if err != nil {
			log.Fatal("Failed to send message:", err)
		}
		fmt.Println("Message sent to server")
	}()

	// Start the server listening on port {port}
	listener, err := transport.CreateSecureListener(":" + fmt.Sprint(PORT))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Listening on port ", PORT)

	// Accept the incoming connection
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connection accepted from", conn.RemoteAddr(), "peer")

	// Read data from the client
	buffer := make([]byte, 1024) // Buffer to hold incoming data
	n, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		log.Fatal("Failed to read from connection:", err)
	}

	// Print the message received from the client
	if n > 0 {
		fmt.Println("Received message:", string(buffer[:n]))
	}

	// Wait for the client goroutine to finish before exiting
	wg.Wait()
}
