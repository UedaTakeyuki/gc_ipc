package main

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"time"

	"github.com/UedaTakeyuki/erapse"
	"github.com/fsnotify/fsnotify"
)

func main() {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		var times = 0
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
					if event.Name == "/tmp/testinotify" {
						if times == 1 {
							elapse("/tmp/testinotify")
							times = 0
						} else {
							log.Println("Ignored first write.")
							times += 1
						}
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add("/tmp")
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

func elapse(file string) {
	defer erapse.ShowErapsedTIme(time.Now())
	clientTimeLog, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("clientTimeLog", clientTimeLog)
	}
	//	if len(clientTimeLog) > 0 {
	clientTimeUint64 := int64(binary.LittleEndian.Uint64(clientTimeLog))
	clientUnixTime := time.Unix(0, clientTimeUint64)
	elapseNano := time.Now().Sub(clientUnixTime).Microseconds()
	log.Println(elapseNano, "ns")
	// }
}
