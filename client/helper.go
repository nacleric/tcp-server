package main

import (
	"fmt"
)

type Message struct {
	username string
	body     string
	//date     time.Time
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
