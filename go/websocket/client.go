// +build ignore
package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

type MsgType int

const (
	INITIAL MsgType = iota
	CONTRACT
	TRADE
	BOOK
)

type Msg struct {
	Type MsgType
	Buf  []byte
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/data"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	message := Msg{Type: INITIAL, Buf: []byte{'a', 'b'}}
	buf, err := json.Marshal(message)
	if err != nil {
		log.Println("write:", err)
		return
	}
	err = c.WriteMessage(websocket.TextMessage, buf)
	if err != nil {
		log.Println("write:", err)
		return
	}

	message = Msg{Type: TRADE, Buf: []byte{'c', 'b'}}
	buf, err = json.Marshal(message)
	if err != nil {
		log.Println("write:", err)
		return
	}
	err = c.WriteMessage(websocket.TextMessage, buf)
	if err != nil {
		log.Println("write:", err)
		return
	}

	message = Msg{Type: BOOK, Buf: []byte{'d', 'b'}}
	buf, err = json.Marshal(message)
	if err != nil {
		log.Println("write:", err)
		return
	}
	err = c.WriteMessage(websocket.TextMessage, buf)
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		select {
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			c.Close()
			return
		}
	}
}
