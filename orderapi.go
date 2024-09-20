package main

import (
	"flag"
	"fmt"
	"orderapi/servers"
	"os"
)

func main() {
	// Define a command-line flag for the server type
	s := flag.String("s", "", "Type of server to run (udp, tcp, http)")
	flag.Parse()

	// Check the server type and call the corresponding function
	switch *s {
	case "tcp", "TCP":
		servers.RunTCPServer()
	case "HTTP", "http":
		servers.RunHTTPServer()
	default:
		fmt.Println("Please specify a valid server type: udp, tcp, or http")
		os.Exit(1)
	}
}
