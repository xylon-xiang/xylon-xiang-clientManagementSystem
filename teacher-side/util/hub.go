package util

import "github.com/gorilla/websocket"

type Connection struct {
	IpAddress string
	Conn *websocket.Conn
}

type Hub struct {
	Client map[*Connection]bool

	Broadcast chan []byte

	Register chan *Connection

	Unregister chan *Connection
}

func NewHub() *Hub {
	return &Hub{
		Client:       make(map[*Connection]bool),
		Broadcast:    make(chan []byte),
		Register:     make(chan *Connection),
		Unregister:   make(chan *Connection),
	}
}

func (h *Hub) Run() {
	for true {
		select {
		case client := <-h.Register:
			h.Client[client] = true

		case client := <-h.Unregister:
			if _, ok := h.Client[client]; ok {
				delete(h.Client, client)
				_ = client.Conn.Close()
			}

		case message := <-h.Broadcast:
			for client := range h.Client {
				_ = client.Conn.WriteMessage(websocket.TextMessage, message)
			}
		}
	}
}
