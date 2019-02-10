package main

import (
	"fmt"
	"log"
	"net"
)

//type User struct { // might put this into its own file
//	username string
//}

func main() {
	port := ":2000"
	//fmt.Println("Enter in portnumber")
	//var p string
	//fmt.Scanln(&p)

	//Dial function connects to port 2000
	conn, err := net.Dial("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()

	// Handles the connection
	// Loop will exit when connection fails
	for conn != nil {
		// grabs user input
		var msg string
		fmt.Scanln(&msg)

		//go routine handles messages being sent
		go func(m string, c net.Conn) {
			// converts string into a byteslice
			// sends bytes over tcp
			buf := []byte(m)
			c.Write(buf)
		}(msg, conn)
	}
}
