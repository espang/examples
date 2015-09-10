package main

import (
	"log"
	"strings"
	"time"

	"github.com/nats-io/nats"
)

const subject = "***test***"

type Req struct {
	nc *nats.Conn
	ec *nats.EncodedConn
}

func (r *Req) log(format string, a ...interface{}) {
	log.Printf("Req:  "+format, a...)
}
func (r *Req) request(s string) (string, error) {
	msg, err := r.nc.Request(subject, []byte(s), 1*time.Second)
	if err != nil {
		return "", err
	}
	return string(msg.Data), nil
}
func (r *Req) loop() {
	tick := time.Tick(time.Second)
	for _ = range tick {
		msg, err := r.request("text in")
		if err != nil {
			r.log("Error requesting 'text in': %v", err)
			continue
		}
		r.log("Got response: '%s'", msg)
	}
}

type Res struct {
	conn *nats.EncodedConn
}

func (r *Res) log(format string, a ...interface{}) {
	log.Printf("RES: "+format, a...)
}
func (r *Res) subscribe() error {
	handler := func(m *nats.Msg) {
		r.log("got: %s", m.Data)
		r.conn.Publish(m.Reply, strings.Repeat(string(m.Data), 3))
	}

	_, err := r.conn.Subscribe(subject, handler)
	return err
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

	r := &Req{nc, c}
	re := &Res{c}

	re.subscribe()
	r.loop()
}
