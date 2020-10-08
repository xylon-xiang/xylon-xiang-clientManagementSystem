package main

import (
	"clientManagementSystem/config"
	"clientManagementSystem/gui"
	"clientManagementSystem/module"
	"clientManagementSystem/student-side/constant"
	"clientManagementSystem/student-side/log"
	"clientManagementSystem/student-side/screenshot"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	log2 "log"
	"net/http"
	"time"
)

var (
	websocketConnection *websocket.Conn
	inputValues         module.InputValue
	logStatus           chan bool
	homework            chan string
)

func main() {

	//ClassName = "soft"
	//ClassStartDate = int64(1)

	logStatus = make(chan bool)
	homework = make(chan string)

	go startWebsocketListening()

	go readWebsocketMessage()

	/*
		the main function is used to handle input and start specific function
	*/

	studentSideGUI := gui.NewGUI(&inputValues)

	GUIShow(studentSideGUI)

	// run in the last
	studentSideGUI.Application.Run()

	//
	//
	//for {
	//	_, _ = fmt.Scanln(&input)
	//
	//	switch input {

	//	case constant.UPLOADHOMEWORK:
	//		UploadHomeworkController()
	//
	//	case constant.LOGOUT:
	//		go closeWebsocket(websocketConnection)
	//	}
	//
	//}

}

func GUIShow(windows *gui.Windows) {

	const (
		LoginWindowsName = "login"

		MainWindowsName = "client management system Student-Side"

		AlarmWindowsName = "alarm"
	)

	windows.RegisterWindows(LoginWindowsName, 500, 300)
	windows.AddLoginForm(LoginWindowsName, &inputValues, LogController)
	loginWindowsHandle := (*windows).WindowsMap[LoginWindowsName]

	windows.RegisterWindows(MainWindowsName, 800, 600)
	mainWindowsHandle := (*windows).WindowsMap[MainWindowsName]

	(*loginWindowsHandle).Show()

	// read the logStatus
	go func() {
		passwordRight := <-logStatus

		if passwordRight {

			// init the main windows
			//windows.AddAlarmItem("alarm", "Log successful")
			(*loginWindowsHandle).Close()
			windows.StudentSideMainWindows(MainWindowsName, &inputValues, &homework)
			(*mainWindowsHandle).Show()

		} else {
			windows.RegisterWindows(AlarmWindowsName, 100, 80)
			alarmWindowsHandle := windows.WindowsMap[AlarmWindowsName]
			windows.AddAlarmItem(AlarmWindowsName, "log failure, please retry!")
			(*alarmWindowsHandle).Show()
		}

	}()

}

func LogController() {

	studentLogPost := module.StudentLogPost{
		StudentId:       inputValues.StudentId,
		StudentName:     inputValues.StudentName,
		StudentPassword: inputValues.StudentPassword,
		ClassName:       inputValues.ClassName,
		ClassStartDate:  inputValues.ClassStartDate,
	}

	isPasswordRight, err := log.SendLoginHttp(inputValues.StudentId, studentLogPost)
	if err != nil {

		logStatus <- false

		log2.Printf("checkpassword: %v\n", err)
		return
	}

	if !isPasswordRight {

		logStatus <- false

		log2.Printf("wrong password")
		return
	}

	logStatus <- true
}

func startWebsocketListening() {
	var addr = flag.String("addr", config.Config.APIConfig.WebsocketPort, "websocket service address")
	flag.Parse()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		wsConn, err := log.StartWebsocket(request, writer)
		if err != nil {
			log2.Printf("startwebsocket: %v\n", err)
			return
		}

		websocketConnection = wsConn
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log2.Printf("listenAndServer: %v", err)
		return
	}
}

// FIXME: fix the bug in close websocket function
func closeWebsocket(conn *websocket.Conn) {

	duration := time.Second *
		time.Duration(config.Config.APIConfig.WebsocketCloseDuration)

	err := conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, constant.LOGOUT),
		time.Now().Add(duration))

	if err != nil {
		log2.Printf("websocketClose error: %v", err)
		return
	}

	time.Sleep(duration)

	conn.Close()
}

// this function is used for get message from websocket co-talking with teacher
func readWebsocketMessage() {

	for {

		if websocketConnection != nil {
			_, msg, err := websocketConnection.ReadMessage()
			if err != nil {
				return
			}

			if string(msg) == "screenshot" {
				// send screenshot in a new go routine while log2 the error when it is not nil
				go func() {
					err := screenshot.SendScreenshot(&inputValues)
					if err != nil {
						log2.Printf("send screenshot error: %v\n", err)
						return
					}
				}()
			}

			homework <- string(msg)

			fmt.Println(string(msg))
		}

	}

}
