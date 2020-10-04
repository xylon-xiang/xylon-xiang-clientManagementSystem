package util

import "github.com/gorilla/websocket"

func ReadWebsocketMessage(conn *websocket.Conn) (msg []byte, err error) {

	_, msg, err = conn.ReadMessage()
	if err != nil{
		return nil, err
	}

	return msg, nil

}

