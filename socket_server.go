package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("its working")

	// net.Listen() creates the server
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	listener, err := net.Listen("tcp", C.Port) //C.Port is accessed in the config file
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		// Wait for a connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("someone has connected")
		}

		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		// reads and echoes messages
		go func(c net.Conn) {
			for {
				buf := make([]byte, 0, 4096) // big buffer
				tmp := make([]byte, 2048)

				incoming, err := c.Read(tmp)
				if err != nil {
					log.Println(err)
					break
				}
				// Reconstructs bytes into a string
				buf = append(buf, tmp[:incoming]...)
				msg := string(buf)
				fmt.Println("message:", msg)

			}
		}(conn)
		//conn.Close()
	}
}
