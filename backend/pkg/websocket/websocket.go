package websocket

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[Upgrade] Error upgrading connection: %s\n", err)
		return ws, err
	}
	return ws, nil
}

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("[Reader] Error reading message: %s\n", err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("[Reader] Error writing message: %s\n", err)
			return
		}
	}
}

func Writer(conn *websocket.Conn) {
	for {
		log.Printf("[Writer] Sending message")
		messageType, r, err := conn.NextReader()
		if err != nil {
			log.Printf("[Writer] Error reading message: %s\n", err)
			return
		}

		w, err := conn.NextWriter(messageType)
		if err != nil {
			log.Printf("[Writer] Error writing message type: %s\n", err)
			return
		}

		if _, err := io.Copy(w, r); err != nil {
			log.Printf("[Writer] Error copying: %s\n", err)
			return
		}

		if err := w.Close(); err != nil {
			log.Printf("[Writer] Error closing connection: %s\n", err)
			return
		}
	}
}

