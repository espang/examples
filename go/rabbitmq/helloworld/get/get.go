package main

import (
	"log"

	"github.com/streadway/amqp"
)

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

	wc := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Got a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for msgs.")
	<-wc
}
