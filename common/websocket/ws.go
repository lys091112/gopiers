package websocket

import (
	"log"

	ws "github.com/gorilla/websocket"
)

func Get() {
	url := "ws://10.26.26.167:8001/aiwchat/ws/voice/update?brokerId=203620515"

	conn, _, err := ws.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}

	defer conn.Close()

	conn.WriteMessage(ws.BinaryMessage, []byte("hello world"))
	conn.WriteMessage(ws.BinaryMessage, []byte("hello 58"))
	err = conn.WriteMessage(ws.BinaryMessage, []byte("End"))
	if err != nil {
		log.Fatal("write: ", err)
	}

	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("read: ", err)
	}
	log.Println(string(message))
}
