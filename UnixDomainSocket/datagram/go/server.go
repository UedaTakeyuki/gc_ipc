package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	os.Remove("/tmp/ud_ucase")
	conn, err := net.ListenUnixgram("unixgram", &net.UnixAddr{"/tmp/ud_ucase", "unixgram"})
	if err != nil {
		panic(err)
	}
	defer os.Remove("/tmp/ud_ucase")

	for {
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(buf[:n]))
	}
}
