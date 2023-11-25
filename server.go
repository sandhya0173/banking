package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// Read data from the client
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	clientData := string(buffer[:n])
	fmt.Printf("Received data from client: %s\n", clientData)

	// Send a response back to the client
	response := "Hello from server!"
	conn.Write([]byte(response))

	// Close the connection
	conn.Close()
}

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on :9999")

	for {
		// Accept a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}
