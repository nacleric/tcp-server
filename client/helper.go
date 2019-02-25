package main

import (
	"fmt"
	"net"
)

type Message struct {
	username string
	body     string
}

/*
//encodes Message and allows it to be passed over tcp
func gobject(m *Message) {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(m); err != nil {
		panic(err)
	}
	fmt.Println("encoded struct: ", b)

}
*/
// Helper that displays what bytes are being sent
func PrintBytes(b []byte) {
	//fmt.Println("\n[DEBUG]", "byteOutput:", b, "string:", string(b))
	fmt.Println("[DEBUG]")
	fmt.Println(b)
	fmt.Println(string(b))
}

func HandleMessage(m *Message, c net.Conn) {
	str := fmt.Sprintf("%s: %s", m.username, m.body)

	// converts string into a byteslice
	// sends bytes over tcp
	buf := []byte(str)
	c.Write(buf)
}
