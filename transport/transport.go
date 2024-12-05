package transport

import (
	"crypto/tls"
	// "log"
	"net"
	// "os"
)

// Create a TCP listener (a server) with TLS encryption using the self-signed certificate
func CreateSecureListener(address string) (net.Listener, error) {
	// loads self-signed certificate and private key use for encryption
	cert, err := tls.LoadX509KeyPair("server-cert.pem", "server-key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		// Skip peer cert verification, this ensures that the application doesn't fail due to certificate validation errors
		InsecureSkipVerify: true,
	}

	// Create a TCP listener
	// The tcp protocol means we are using the Transmission Control Protocol, which is a connection-based protocol.
	listener, err := tls.Listen("tcp", address, config)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

// Connect to another server or peer with TLS using the self-signed certificate on the network
// used by the client (or peer) to connect to a server using TLS encryption
func ConnectToPeer(address string) (net.Conn, error) {
	// In real-world scenarios, you would verify the server's certificate here
	config := &tls.Config{
		InsecureSkipVerify: true, // Only for testing
	}

	// Connect to the server
	//creates a secure TLS connection to the specified address (the server you want to connect to).
	conn, err := tls.Dial("tcp", address, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
