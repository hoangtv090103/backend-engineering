package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// Send a message to the client
	fmt.Fprintln(conn, "Hello.")

	// Read data from the client
	reader := bufio.NewReader(conn)
	for {
		message, _ := reader.ReadString('\n')
		fmt.Println("Received:", message)
	}
}

const port = 8080

func main() {
	// Start listening on port 8080
	// listener, err := net.Listen("tcp", ":8080")
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{
		Port: port,
	})
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port", port)

	for {
		// Accept connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn) // Handle each connection concurrently
	}
}
