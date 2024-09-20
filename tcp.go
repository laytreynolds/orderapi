package main

import (
	"fmt"
	"net"
	"os"
)

func runTCPServer() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("TCP server listening on port 8081")

}
