package main

import (
	"clientManagementSystem/config"
	"clientManagementSystem/module"
	"clientManagementSystem/student-side/constant"
	"clientManagementSystem/student-side/log"
	"clientManagementSystem/student-side/quiz"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	log2 "log"
	"net/http"
	"time"
)

var websocketConnection *websocket.Conn

func main() {

	go startWebsocketListening()

	go readWebsocketMessage()

	/*
		the main function is used to handle input and start specific function
	*/
	var (
		input           string
		studentId       string
		studentName     string
		studentPassword string
		QuestionContent string
	)

	for {
		_, _ = fmt.Scanln(&input)

		switch input {
		case constant.LOGIN:
			fmt.Println("Please enter the studentId")
			_, _ = fmt.Scanln(&studentId)
			fmt.Println("Please enter the studentPassword")
			_, _ = fmt.Scanln(&studentPassword)
			fmt.Println("Please enter the studentName")
			_, _ = fmt.Scanln(&studentName)

			LogController(studentId, studentPassword)

		case constant.QUIZ:
			fmt.Println("Please enter the Question content:")
			_, _ = fmt.Scanln(&QuestionContent)
			QuizController(studentId, studentName, QuestionContent)

		case constant.LOGOUT:
			go closeWebsocket(websocketConnection)
		}

	}

}

func LogController(studentId string, studentPassword string) {

	studentLogPost := module.StudentLogPost{
		StudentPassword: studentPassword,
		ClassName: "Software Engineering 1210",
		ClassStartDate: 1601790639,
	}

	isPasswordRight, err := log.SendLoginHttp(studentId, studentLogPost)
	if err != nil {
		log2.Printf("checkpassword: %v\n", err)
		return
	}

	if !isPasswordRight {

		/*
			this part of codes is used to alarming student to change his password
		*/
		log2.Printf("wrong password")

		return
	}
}

func QuizController(studentId string, studentName string, question string) {
	answer, err := quiz.HandUp(studentId, studentName, question)
	if err != nil {
		log2.Printf("hand up err: %v\n", err)
		return
	}

	log2.Printf("answer is: \n%v", answer)
	return
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

func closeWebsocket(conn *websocket.Conn) {

	duration := time.Second *
		time.Duration(config.Config.APIConfig.WebsocketCloseDuration)

	err := conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, constant.LOGOUT),
		time.Now().Add(duration))

	if err != nil {
		log2.Printf("websocketClose: %v", err)
		return
	}

	time.Sleep(duration)

	conn.Close()
}

// this function is used for test
func readWebsocketMessage() {

	for {

		if websocketConnection != nil {
			_, msg, err := websocketConnection.ReadMessage()
			if err != nil {
				return
			}

			fmt.Println(string(msg))
		}

	}

}
