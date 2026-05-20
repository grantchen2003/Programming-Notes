// server/main.go
package main

import (
	"fmt"
	"net"
)

func main() {
	// Start listening on TCP port 8080.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		// Print error if unable to start the server.
		fmt.Println("Error starting server:", err)
		return
	}
	// Ensure the listener is closed when main exits.
	defer listener.Close()
	fmt.Println("Server listening on port 8080")

	// Continuously accept incoming connections.
	for {
		conn, err := listener.Accept()
		if err != nil {
			// Log the error and continue accepting new connections.
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// Handle the connection concurrently.
		go handleConnection(conn)
	}
}

// handleConnection manages an individual client connection.
func handleConnection(conn net.Conn) {
	// Ensure the connection is closed when function exits.
	defer conn.Close()

	// Create a buffer to hold incoming data.
	buffer := make([]byte, 1024)

	// Read loop for handling client messages.
	for {
		// Read data into the buffer.
		n, err := conn.Read(buffer)
		if err != nil {
			// Log error if reading fails or connection is closed.
			fmt.Println("Connection closed or error reading:", err)
			return
		}
		// Print the received message from the client.
		fmt.Println("Received:", string(buffer[:n]))
		// Echo the message back to the client.
		_, err = conn.Write(buffer[:n])
		if err != nil {
			// Log error if writing fails.
			fmt.Println("Error writing response:", err)
			return
		}
	}
}
