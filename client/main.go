package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

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
	defer conn.Close()

	//READSTRING MIGHT BE THE PROBLEM
	// Asks user to input their name
	fmt.Print("Enter Username: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Handles the connection
	// Loop will exit when connection fails
	for conn != nil {
		// grabs user input
		fmt.Print("Enter Message: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		var msg = &Message{username: username, body: text}

		//FIX THIS KEEPS ADDING NEWLINE
		//go routine handles messages being sent
		func(m *Message, c net.Conn) {
			//str := fmt.Sprintf("%s: %s", m.username, m.body)
			str := m.username + ": " + m.body
			// converts string into a byteslice
			// sends bytes over tcp
			//buf := []byte(string(m))
			buf := []byte(str)
			PrintBytes(buf)
			c.Write(buf)

		}(msg, conn)

	}
}
