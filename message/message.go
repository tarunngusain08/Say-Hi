package main

import (
	"Say-Hi/message/contracts"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type MessageHandler struct {
	clients  map[*contracts.Client]bool
	mu       sync.Mutex
	upgrader websocket.Upgrader
}

func NewMessageHandler(upgrader websocket.Upgrader) *MessageHandler {
	return &MessageHandler{
		clients:  make(map[*contracts.Client]bool),
		mu:       sync.Mutex{},
		upgrader: upgrader,
	}
}

func (m *MessageHandler) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	user1 := r.URL.Query().Get("user1")
	user2 := r.URL.Query().Get("user2")
	client1 := &contracts.Client{Username: user1, Conn: conn}
	client2 := &contracts.Client{Username: user2, Conn: conn}

	m.mu.Lock()
	m.clients[client1] = true
	m.clients[client2] = true
	m.mu.Unlock()

	defer func() {
		m.mu.Lock()
		delete(m.clients, client1)
		delete(m.clients, client2)
		m.mu.Unlock()
	}()

	for {
		var msg contracts.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}

		// Handle your message logic here (e.g., store in database, send to recipient)
		// For simplicity, we'll broadcast the message to all connected clients.

		m.mu.Lock()
		for c := range m.clients {
			if c.Username == msg.Recipient {
				if err := client1.Conn.WriteJSON(msg); err != nil {
					log.Println(err)
					return
				}
			}
		}
		m.mu.Unlock()
	}
}

func main() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	h := NewMessageHandler(upgrader)
	http.HandleFunc("/ws", h.handleWebSocket)
	port := 8080
	fmt.Printf("Server is running on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}