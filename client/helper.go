package main

import (
	"fmt"
	"net"
)

type message struct {
	username string
	body     string
}

// Helper that displays what bytes are being sent
func printBytes(b []byte) {
	fmt.Println("[DEBUG]")
	fmt.Println(b)
	fmt.Println(string(b))
}

// turns a string into a byteslice and prepares it to be sent over tcp
func handleMessage(m *message, c net.Conn) {
	str := fmt.Sprintf("%s: %s", m.username, m.body)

	// converts string into a byteslice
	// sends bytes over tcp
	buf := []byte(str)
	c.Write(buf)
}
