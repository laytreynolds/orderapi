package servers

import (
	"fmt"
	"io"
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

	packet := make([]byte, 4096)
	tmp := make([]byte, 4096)

	defer c.Close()
	
	for {
		_, err := c.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			println("END OF FILE")
			break
		}
		packet = append(packet, tmp...)
	}
	num, _ := c.Write(packet)
	fmt.Printf("Wrote back %d bytes, the payload is %s\n", num, string(packet))

}
