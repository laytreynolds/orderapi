package servers

import (
	"fmt"
	"log"
	"net"
	"orderapi/helpers"
)

func RunTCPServer() {

	listener, err := net.Listen("tcp", port)

	helpers.Logger("Listening on %v", port)

	if err != nil {
		log.Fatal(err)

	}
	defer listener.Close()

	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}

}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

}
