package main

import (
	"fmt"
	"net"
)

const port = 8081

func main() {
	addr := net.UDPAddr{
		Port: port,
		// IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("Server is running on port", port)

	buffer := make([]byte, 1024)
	

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Printf("server got: %s from %s\n", string(buffer[:n]), remoteAddr)
	}
}