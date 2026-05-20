// client/main.go
package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	// Attempt to connect to the server running on localhost at port 8080.
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// If the connection fails, print the error and exit.
		fmt.Println("Error connecting to server:", err)
		return
	}
	// Ensure the connection is properly closed when the function exits.
	defer conn.Close()

	// Define the message to be sent to the server.
	message := ""
	for i := 0; i < 10000; i++ {
		message += strconv.Itoa(i) + " "
	}
	// Send the message by writing its byte slice representation to the connection.
	_, err = conn.Write([]byte(message))
	if err != nil {
		// Print an error message if writing fails.
		fmt.Println("Error sending message:", err)
		return
	}

	// Create a buffer to store the response from the server.
	buffer := make([]byte, 1024)
	// Read the server's response into the buffer.
	n, err := conn.Read(buffer)
	if err != nil {
		// Print an error message if reading fails.
		fmt.Println("Error reading response:", err)
		return
	}

	// Convert the received bytes back to a string and print the server's response.
	fmt.Println("Server response:", string(buffer[:n]))
}
