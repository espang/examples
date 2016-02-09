// +build ignore
package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

const subject = "subj"

type Pub struct {
	ch    *amqp.Channel
	qName string

	mu    *sync.Mutex
	count uint64
}

func (p *Pub) log(format string, a ...interface{}) {
	log.Printf("PUBLISH:  "+format, a...)
}
func (p *Pub) publish(s string) error {
	return p.ch.Publish(
		"",
		p.qName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(s),
		},
	)
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
	conn, err := amqp.Dial("amqp://TestUser:TestUser@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	p := &Pub{ch, q.Name, &sync.Mutex{}, 0}
	p.loop()
}
