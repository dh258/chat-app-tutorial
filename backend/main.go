package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dh258/chat-app-tutorial/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ServeWs] Request host: %s\n", r.Host)

	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Printf("[ServeWs] Error upgrading connection: %s\n", err)
		return
	}

	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	http.HandleFunc("/ws", serveWs)
}

func main() {
	log.Println("Chat App v0.0.1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
