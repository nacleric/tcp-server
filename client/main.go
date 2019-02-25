package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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

	// Asks user to input their name
	fmt.Print("Enter Username: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	username = strings.TrimSuffix(username, "\n") //trims username from .ReadString
	if err != nil {
		log.Fatal(err)
	}

	// Handles the connection
	// Loop will exit when connection fails
	for conn != nil {

		// grabs user input
		// and assigns it to Message Struct
		fmt.Print("Enter Message: ")
		text, err := reader.ReadString('\n')
		//text = strings.TrimSuffix(text, "\n") //trims text from .ReadString
		if err != nil {
			log.Fatal(err)
		}
		var msg = &Message{username: username, body: text}

		// go routine handles messages being sent
		go func(m *Message, c net.Conn) {
			str := fmt.Sprintf("%s: %s", m.username, m.body)
			// converts string into a byteslice
			// sends bytes over tcp
			buf := []byte(str)
			c.Write(buf)

		}(msg, conn)

	}
}
