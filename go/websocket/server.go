// +build ignore
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

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

func update(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		var m Msg
		err = json.Unmarshal(message, &m)
		if err != nil {
			log.Printf("could not unmarshal message: %v", err)
		} else {
			log.Printf("Messagetype: %d, Messagebuf: %s", m.Type, m.Buf)
		}
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	fmt.Printf("Socket closed\n")
}

func wsHandler(ws *websocket.Conn) {
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/data", update)
	log.Fatal(http.ListenAndServe(":8080", router))

}
