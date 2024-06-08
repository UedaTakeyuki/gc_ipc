// refer https://gist.github.com/szaydel/423d89bc9fe11d85d332
package main

import (
	"encoding/binary"
	"log"
	"net"
	"os"
	"time"

	"github.com/UedaTakeyuki/erapse"
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
		elapse(buf[:n])
	}
}

func elapse(clientTimeLog []byte) {
	defer erapse.ShowErapsedTIme(time.Now())
	clientTimeUint64 := int64(binary.LittleEndian.Uint64(clientTimeLog))
	clientUnixTime := time.Unix(0, clientTimeUint64)
	elapseNano := time.Now().Sub(clientUnixTime).Nanoseconds()
	log.Println(elapseNano, "ns")
}
