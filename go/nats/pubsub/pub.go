// +build ignore
package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/nats-io/nats"
)

const subject = "subj"

type Pub struct {
	conn *nats.EncodedConn

	mu    *sync.Mutex
	count uint64
}

func (p *Pub) log(format string, a ...interface{}) {
	log.Printf("PUBLISH:  "+format, a...)
}
func (p *Pub) publish(s string) error {
	return p.conn.Publish(subject, s)
}
func (p *Pub) loop() {

	ticks := time.Tick(60 * time.Second)
	var msgsLastMinute uint64
	go func() {
		for _ = range ticks {
			p.mu.Lock()
			msgsLastMinute = p.count / 60
			p.count = 0
			p.mu.Unlock()
			p.log("%d msgs/s", msgsLastMinute)
		}
	}()

	var v float64
	var counter = 0
	for {
		v = rand.Float64()
		err := p.publish(fmt.Sprintf("value=%f", v))
		if err != nil {
			p.log("Error publishsing '%s': %v", v, err)
		}
		counter++
		if counter >= 200 {
			time.Sleep(1 * time.Nanosecond)
			counter = 0
			p.mu.Lock()
			p.count += 200
			p.mu.Unlock()
		}
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

	p := &Pub{c, &sync.Mutex{}, 0}

	p.loop()
}
