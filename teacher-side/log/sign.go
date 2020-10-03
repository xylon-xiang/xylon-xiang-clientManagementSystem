package log

import (
	"clientManagementSystem/teacher-side/constant"
	"clientManagementSystem/teacher-side/module"
	"clientManagementSystem/teacher-side/util"
	"github.com/gorilla/websocket"
	url2 "net/url"
	"strconv"
	"time"
)

func NewWebsocket(Hub *util.Hub, host string) error {
	dialer := websocket.Dialer{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		EnableCompression: true,
	}

	url := url2.URL{
		Scheme: "ws",
		Host: host,
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
	}
	(*Hub).Register <- connection


	msg := []byte(constant.ACCEPT)
	err = conn.WriteMessage(websocket.TextMessage, msg)

	return nil
}


func ChangeSignStatus(studentId string, className string,
	classStartDate int64, targetSignStatus int) (bool, error) {

	startTime := strconv.FormatInt(classStartDate, 10)

	filter := map[string]string{
		"studentId": studentId,
		"ClassName": className,
		"ClassStartDate": startTime,
	}

	result, err := module.FindOne(module.STUDENTSTATUS, filter)
	if err != nil{
		return false, err
	}

	studentStatus := result.(module.StudentStatus)
	studentStatus.SignStatus = targetSignStatus

	err = module.UpdateOne(module.STUDENTSTATUS, studentStatus)
	if err != nil{
		return false, err
	}

	return true, nil
}

func SetSignTime(signStatus string, classInfo module.ClassInfo) error {

	startDate := strconv.FormatInt(classInfo.ClassStartDate, 10)
	filter := map[string]string{
		"studentId": classInfo.StudentsInfo[0].StudentId,
		"className": classInfo.ClassName,
		"classStartDate": startDate,
	}

	result, err := module.FindOne(module.STUDENTSTATUS, filter)
	if err != nil{
		return err
	}

	studentStatus := result.(module.StudentStatus)

	switch signStatus {
	case constant.SIGNIN:
		studentStatus.SignInDate = time.Now().Unix()

	case constant.SIGNOUT:
		studentStatus.SignOutDate = time.Now().Unix()
	}


	err = module.UpdateOne(module.STUDENTSTATUS, studentStatus)
	if err != nil{
		return err
	}




	return nil
}
