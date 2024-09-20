package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	port = ":8080"
)

func main() {
	// Define a command-line flag for the server type
	t := flag.String("type", "", "Type of server to run (udp, tcp, http)")
	flag.Parse()

	// Check the server type and call the corresponding function
	switch *t {
	case "tcp":
		runTCPServer()
	case "http":
		RunHTTPServer()
	default:
		fmt.Println("Please specify a valid server type: udp, tcp, or http")
		os.Exit(1)
	}
}
