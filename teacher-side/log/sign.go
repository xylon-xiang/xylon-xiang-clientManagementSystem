package log

import (
	"clientManagementSystem/config"
	"clientManagementSystem/module"
	"clientManagementSystem/teacher-side/constant"
	"clientManagementSystem/teacher-side/util"
	"github.com/gorilla/websocket"
	url2 "net/url"
	"strconv"
	"time"
)

func NewWebsocket(Hub *util.Hub, host string, studentName string) error {
	dialer := websocket.Dialer{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		EnableCompression: true,
	}

	url := url2.URL{
		Scheme: "ws",
		Host: "localhost" + config.Config.APIConfig.WebsocketPort,
		Path: "/",
	}

	conn, _, err := dialer.Dial(url.String(), nil)
	if err != nil{
		return err
	}

	// add the new connection into the all map
	connection := &util.Connection{
		Conn: conn,
		IpAddress: host,
		StudentName: studentName,
	}
	(*Hub).Register <- connection


	msg := []byte(constant.ACCEPT)
	err = conn.WriteMessage(websocket.TextMessage, msg)

	return nil
}


func SetSignTime(signStatus string, studentStatus module.StudentStatus) error {

	startDate := strconv.FormatInt(studentStatus.ClassStartDate, 10)
	filter := map[string]string{
		"StudentId": studentStatus.StudentId,
		"ClassName": studentStatus.ClassName,
		"ClassStartDate": startDate,
	}

	result, err := util.FindOne(util.STUDENTSTATUS, filter)
	if err != nil{
		return err
	}

	newStudentStatus := result.(module.StudentStatus)

	switch signStatus {
	case constant.SIGNIN:
		newStudentStatus.SignInDate = time.Now().Unix()

	case constant.SIGNOUT:
		newStudentStatus.SignOutDate = time.Now().Unix()
	}


	err = util.UpdateOne(util.STUDENTSTATUS, newStudentStatus)
	if err != nil{
		return err
	}

	return nil
}
