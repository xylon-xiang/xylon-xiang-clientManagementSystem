package util

import (
	"clientManagementSystem/config"
	"github.com/gorilla/websocket"
	log2 "log"
	"time"
)

type Connection struct {
	IpAddress string
	StudentName string
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
			log2.Printf("WebSocket Register: ip: %v\n", client.IpAddress)

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


// FIXME: the close function contains some bugs
// TODO: when close the websocket connection, set the logout time
func (h *Hub) ClearWebConnection() {
	for true {
		for client := range h.Client{
			messageType, msg, err := client.Conn.ReadMessage()
			if err == websocket.ErrCloseSent{

				if messageType == websocket.CloseMessage{
					log2.Printf("closeWebsocket: %v\n", msg)
					h.Unregister <- client
				}

				time.Sleep(time.Duration(config.Config.APIConfig.WebsocketCloseDuration))

				continue
			}

			if err != nil{
				log2.Printf("ReadMessageError: %v\n", err)
			}



		}
	}
}
