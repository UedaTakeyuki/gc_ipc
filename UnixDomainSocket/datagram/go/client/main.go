package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/UedaTakeyuki/erapse"
)

var now int64

func main() {
	socketType := "unixgram" // or "unixgram" or "unixpacket"
	laddr := net.UnixAddr{"/tmp/unixdomaincli", socketType}
	conn, err := net.DialUnix(socketType, &laddr, /*can be nil*/
		&net.UnixAddr{"/tmp/ud_ucase", socketType})
	if err != nil {
		panic(err)
	}
	defer os.Remove("/tmp/unixdomaincli")

	buff := make([]byte, 10)
	var n int
	n, err = writeAndRead(conn, &buff)
	//	n, err = conn.Read(buff)
	if err != nil {
		fmt.Printf("server() read socket failed with err: %v\n", err)
	} else {
		log.Println("n:", n)
		log.Println(string(buff[:n]))
	}
	conn.Close()
}

func writeAndRead(conn net.Conn, buff *[]byte) (n int, err error) {
	defer erapse.ShowErapsedTIme(time.Now())

	_, err = conn.Write(getNowInNanoSec())
	if err != nil {
		panic(err)
	}
	n, err = conn.Read(*buff)
	return
}

func getNowInNanoSec() (buf []byte) {
	now = time.Now().UnixNano()
	buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(now))
	return
}
