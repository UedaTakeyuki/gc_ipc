package main

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"time"

	"github.com/UedaTakeyuki/erapse"
)

func main() {
	defer erapse.ShowErapsedTIme(time.Now())
	now := time.Now().UnixNano()
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(now))
	err := ioutil.WriteFile("/tmp/testinotify", buf, 0777)
	if err != nil {
		log.Println(err)
	}
}
