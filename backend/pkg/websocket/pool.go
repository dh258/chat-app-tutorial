package websocket

import (
	"log"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			log.Printf("[Start] Registering client %s to pool...\n", client.ID)
			pool.Clients[client] = true

			log.Println("[Start] Connection pool size: ", len(pool.Clients))

			for client := range pool.Clients {
				log.Println("[Start] Notifying user joined to client: ", client)
				client.Conn.WriteJSON(Message{
					Type: 1,
					Body: "New User Joined...",
				})
			}
		case client := <-pool.Unregister:
			log.Printf("[Start] Unregistering client %s from pool...\n", client.ID)
			delete(pool.Clients, client)

			log.Println("[Start] Connection pool size: ", len(pool.Clients))

			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{
					Type: 1,
					Body: "User Disconnected...",
				})
			}
		case message := <-pool.Broadcast:
			log.Println("[Start] Sending message to clients in Pool...")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Println("[Start] Error broadcasting message:", err)
					return
				}
			}
		}

	}
}
