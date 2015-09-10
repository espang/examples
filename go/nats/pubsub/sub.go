// +build ignore
package main

import (
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats"
)

const subject = "subj"

type Sub struct {
	conn *nats.EncodedConn

	mu    *sync.Mutex
	count uint64
}

func (s *Sub) log(format string, a ...interface{}) {
	log.Printf("SUBSCRIBE: "+format, a...)
}
func (s *Sub) subscribe() error {
	handler := func(msg string) {
		s.mu.Lock()
		s.count++
		s.mu.Unlock()
	}
	_, err := s.conn.Subscribe(subject, handler)
	return err
}
func (s *Sub) loop() {
	ticks := time.Tick(15 * time.Second)
	var msgsLastMinute uint64
	for _ = range ticks {
		s.mu.Lock()
		msgsLastMinute = s.count / 15
		s.count = 0
		s.mu.Unlock()
		s.log("%d msgs/s", msgsLastMinute)
	}
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("connecting: %s", err)
	}
	c, err := nats.NewEncodedConn(nc, "json")
	if err != nil {
		log.Fatalf("connecting: %s", err)
	}
	defer c.Close()

	s := &Sub{c, &sync.Mutex{}, 0}

	s.subscribe()
	s.loop()
}
