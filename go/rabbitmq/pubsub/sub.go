// +build ignore
package main

import (
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

const subject = "subj"

type Sub struct {
	deliveries <-chan amqp.Delivery

	mu    *sync.Mutex
	count uint64
}

func (s *Sub) log(format string, a ...interface{}) {
	log.Printf("SUBSCRIBE: "+format, a...)
}
func (s *Sub) loop() {
	ticks := time.Tick(15 * time.Second)
	var msgsLastMinute uint64
	go func() {
		for _ = range ticks {
			s.mu.Lock()
			msgsLastMinute = s.count / 15
			s.count = 0
			s.mu.Unlock()
			s.log("%d msgs/s", msgsLastMinute)
		}
	}()
	for _ = range s.deliveries {
		//got a msg
		s.mu.Lock()
		s.count++
		s.mu.Unlock()
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	s := &Sub{msgs, &sync.Mutex{}, 0}

	s.loop()
}
