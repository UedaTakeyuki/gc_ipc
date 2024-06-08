// refer https://learning.oreilly.com/library/view/network-programming-with/9781098128890/c07.xhtml#:-:text=The%20unixgram%20Datagram%20Socket
package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var address = "/tmp/ud_ucase"

func main() {
	os.Remove("/tmp/ud_ucase")
	ctx := context.Background()
	networktype := "unixgram"
	localaddress, err := datagramEchoServer(ctx, networktype, address)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("localaddress", localaddress)
	}
}

func datagramEchoServer(ctx context.Context, network string,
	addr string) (localAddr net.Addr, err error) {
	s, err := net.ListenPacket(network, addr)
	if err != nil {
		return nil, err
	}

	//	go func() {
	go func() {
		<-ctx.Done()
		_ = s.Close()
		if network == "unixgram" {
			_ = os.Remove(addr)
		}
	}()

	buf := make([]byte, 1024)
	for {
		var n int
		var clientAddr net.Addr
		n, clientAddr, err = s.ReadFrom(buf)
		if err != nil {
			return
		}
		elapse(buf[:n])

		_, err = s.WriteTo(getNowInNanoSec(), clientAddr)
		if err != nil {
			return
		}
	}
	//	}()

	localAddr = s.LocalAddr()
	return
}

func elapse(clientTimeLog []byte) {
	//	defer erapse.ShowErapsedTIme(time.Now())
	start := time.Now()
	defer fmt.Printf("eraps %d ns\n", time.Now().Sub(start).Nanoseconds())
	clientTimeUint64 := int64(binary.LittleEndian.Uint64(clientTimeLog))
	clientUnixTime := time.Unix(0, clientTimeUint64)
	elapseNano := time.Now().Sub(clientUnixTime).Nanoseconds()
	log.Println(elapseNano, "ns")
}

func getNowInNanoSec() (buf []byte) {
	now := time.Now().UnixNano()
	buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(now))
	return
}
