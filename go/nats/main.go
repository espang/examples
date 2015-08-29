package main

import (
	"os"

	"github.com/cihub/seelog"
	"github.com/nats-io/nats"
)

type Person struct {
	Name    string
	Address string
	Age     int
}

func main() {
	defer seelog.Flush()
	seelog.Trace("Start nats test program")
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		seelog.Errorf("connecting: %s", err)
		os.Exit(1)
	}
	c, err := nats.NewEncodedConn(nc, "json")
	if err != nil {
		seelog.Errorf("connecting: %s", err)
		os.Exit(1)
	}
	defer c.Close()

	err = c.Publish("foo", "Hello World")
	if err != nil {
		seelog.Errorf("publishing1: %s", err)
		os.Exit(1)
	}

	_, err = c.Subscribe("foo", func(s string) {
		seelog.Infof("Received a message: %s", s)
	})
	if err != nil {
		seelog.Errorf("subscribing1: %s", err)
		os.Exit(1)
	}

	_, err = c.Subscribe("hello", func(p *Person) {
		seelog.Infof("Received a person %#v", p)
	})
	if err != nil {
		seelog.Errorf("subscribing2: %s", err)
		os.Exit(1)
	}

	person := &Person{Name: "na", Age: 2, Address: "asdf 3773, Germany"}

	err = c.Publish("hello", person)
	if err != nil {
		seelog.Errorf("publishing2: %s", err)
		os.Exit(1)
	}
}
