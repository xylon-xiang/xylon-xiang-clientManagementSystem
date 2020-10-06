package object_operation

import (
	"clientManagementSystem/teacher-side/util"
	"github.com/gorilla/websocket"
)


// if err is nil but 'bool' is false, it means no such a Websocket connection
func SendScreenshotWsRequest(hub *util.Hub, studentName string) (bool, error) {

	for client := range hub.Client{
		if client.StudentName == studentName{

			err := client.Conn.WriteMessage(websocket.TextMessage, []byte("screenshot"))
			if err != nil{
				return false, err
			}

			return true, nil
		}
	}

	return false, nil
}
