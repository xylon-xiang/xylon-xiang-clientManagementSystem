package main

import (
	"clientManagementSystem/config"
	"clientManagementSystem/module"
	"clientManagementSystem/teacher-side/constant"
	"clientManagementSystem/teacher-side/log"
	"clientManagementSystem/teacher-side/object_operation"
	"clientManagementSystem/teacher-side/q_a"
	"clientManagementSystem/teacher-side/screenshot"
	"clientManagementSystem/teacher-side/util"
	"fmt"
	"github.com/labstack/echo"
	log2 "log"
	"net/http"
	"strconv"
)

var (
	Hub        *util.Hub
	inputMutex bool = true
)

// TODO: start a go routine to set the signStatus when class is over
func init() {
	Hub = util.NewHub()
	go Hub.Run()
	//go Hub.ClearWebConnection()
}

func main() {

	go handleInput()

	e := echo.New()

	// section 1
	e.POST(config.Config.APIConfig.StudentLogAPI.Path, LoginController)

	// section 2
	e.POST(config.Config.APIConfig.StudentHandUpAPI.Path, QuizController)

	// section 3, Q&A homework function
	e.POST(config.Config.APIConfig.HomeworkAPI.Path, HomeworkController)
	e.POST(config.Config.APIConfig.HomeworkAPI.FileSavePath, HomeworkFileSaveController)

	e.POST(config.Config.APIConfig.ScreenshotAPI.Path, ScreenshotController)
	e.PUT(config.Config.APIConfig.ScreenshotAPI.UpdateFrozenDurationPath, ScreenFrozenController)

	e.Logger.Fatal(e.Start(":1234"))

}

func handleInput() {
	var (
		input            string
		studentId        string
		className        string
		classStartDate   int64
		targetSignStatus int
	)
	for true {

		if !inputMutex {
			continue
		}

		fmt.Println("Please enter the object_operation")
		_, _ = fmt.Scanln(&input)

		switch input {
		case constant.IMPORT:
			object_operation.ImportClassTestCase()

		case constant.CHANGESIGNSTATUS:

			// handle the input

			//fmt.Println("please enter the studentId :")
			//_, _ = fmt.Scanln(&studentId)
			studentId = "U0000"

			//fmt.Println("please enter the className :")
			//_, _ = fmt.Scanln(&className)
			className = "Software Engineering 1210"

			//fmt.Println("please enter the class start time: ")
			//_, _ = fmt.Scanln(&classStartDate)
			classStartDate = int64(1601790639)

			fmt.Println("please enter the target sign status: ")
			_, _ = fmt.Scanln(&targetSignStatus)

			//classStart := int64(0)
			changeSuccess, err := object_operation.ChangeSignStatus(
				studentId, className, classStartDate, targetSignStatus)
			if err != nil {
				log2.Printf("change sign status error: %v", err)
				continue
			}
			if changeSuccess {
				log2.Println("change successful")
				continue
			} else {
				log2.Println("change failure: ")
				continue
			}

			// TODO: handle the input
			// TODO: handle the error

		case constant.GETATTENDANCERATE:

			//handle the input
			fmt.Println("please enter the studentId :")
			_, _ = fmt.Scanln(&studentId)

			fmt.Println("please enter the className :")
			_, _ = fmt.Scanln(&className)

			studentStatus, err := object_operation.QueryStudentStatus(studentId, className)
			if err != nil {
				continue
			}
			rate := object_operation.QueryAttendanceRate(studentStatus)
			rateStr := strconv.FormatFloat(rate, 'E', -1, 64)
			fmt.Println("this student's attendance rate: " + rateStr)

		case constant.ASSIGNMENT:
			AssignmentController()

		case constant.GETCUMULATIVESCORE:
			//object_operation.QueryCumulativeScore(db.StudentStatus{})

		case constant.GETEACHCLASSSTATUS:
			studentStatus, err := object_operation.QueryStudentStatus(studentId, className)
			if err != nil {
				continue
			}
			byteStream, err := object_operation.QueryEachClassStatus(studentStatus)
			if err != nil {
				continue
			}
			fmt.Println(string(byteStream))

		case constant.PUBLISHHOMEWORK:
			homeworkInfo, err := object_operation.GetHomeworkInfoTest()
			if err != nil {
				log2.Printf("get homework info err: %v\n", err)
				continue
			}

			success, err := object_operation.PublishHomeworkInfo(Hub, homeworkInfo)
			if err != nil || !success {
				log2.Printf("publish homework info err: %v\n", err)
				continue
			}
			log2.Printf("publish homework success!\n")

		case constant.CHANGESTUDENTHOMEWORKSCORE:
			ChangeStudentHomeworkScoreController()

		case constant.SCREENSHOT:
			SendScreenRequestController()

		case constant.GETHOMEWORKSTATUS:
			//object_operation.QueryHomeworkStatus(db.StudentStatus{})
		}
	}
}

func LoginController(context echo.Context) error {

	studentStatus := new(module.StudentStatus)
	if err := context.Bind(studentStatus); err != nil {
		return err
	}

	studentStatus.StudentId = context.Param("studentId")

	isMatch, err := log.CheckPassword(studentStatus.StudentId, studentStatus.StudentPassword)
	if err != nil {
		return err
	}

	if isMatch {
		host := context.Request().RemoteAddr

		err = log.NewWebsocket(Hub, host, studentStatus.StudentName)
		if err != nil {
			return err
		}

		err = log.SetSignTime(constant.SIGNIN, *studentStatus)
		if err != nil {
			return err
		}

		return context.String(http.StatusOK, constant.ACCEPT)

	} else {
		return context.String(http.StatusUnauthorized, constant.FAILURE)
	}

}

func QuizController(context echo.Context) error {

	quizBody := new(module.QuizPostBody)
	if err := context.Bind(quizBody); err != nil {
		return err
	}

	inputMutex = false
	log2.Printf("hand up: %v \n", quizBody)
	fmt.Println("please enter your answer")
	var answer string
	_, _ = fmt.Scanln(&answer)
	inputMutex = true

	return context.String(http.StatusOK, "I am coming")
}

func AssignmentController() {

	// TODO: this part of code should be transport as teachers' inputting
	question := module.HomeworkInfo{
		HomeworkTitle: "Question 1: how old are you?",
		HomeworkType:  constant.TEXT,
	}

	questionStr := fmt.Sprintf("%v", question)
	byteStream := []byte(questionStr)

	Hub.Broadcast <- byteStream

}

func HomeworkController(context echo.Context) error {

	studentStatus := new(module.StudentStatus)
	if err := context.Bind(studentStatus); err != nil{
		return err
	}

	if err := q_a.AutoCorrectHomeWork(&((*studentStatus).HomeworksInfo)); err != nil{
		return err
	}

	if err := q_a.UpdateStudentHomeworkStatus(*studentStatus); err != nil{
		return err
	}

	return context.String(http.StatusOK, "")
}

func HomeworkFileSaveController(context echo.Context) error {

	startDate, err := strconv.ParseInt(context.FormValue("classStartDate"), 10, 64)
	if err != nil{
		return err
	}

	studentStatus := module.StudentStatus{
		StudentInfo: module.StudentInfo{
			StudentId: context.FormValue("studentId"),
		},
		Class: module.Class{
			ClassName: context.FormValue("className"),
			ClassStartDate: startDate,
		},
	}

	questionTitle := context.FormValue("questionTitle")
	file, err := context.FormFile("file")
	if err != nil{
		return err
	}

	// save student upload file in local
	err = q_a.SaveFileQuestionAsFile(studentStatus, questionTitle, file)
	if err != nil{
		return err
	}

	return context.String(http.StatusOK, "upload success")
}

func ChangeStudentHomeworkScoreController() {

	status := module.StudentStatus{

	}

	err := object_operation.ChangeHomeworkStatus(status)
	if err != nil{
		log2.Printf("change homework status error: %v\n", err)
		return
	}

	log2.Println("change student homework status success!")
	return
}

func SendScreenRequestController() {

	success, err := object_operation.SendScreenshotWsRequest(Hub, "Bob")
	if err != nil{
		log2.Printf("send screenshot websocket request error: %v\n", err)
		return
	}

	if !success{
		log2.Printf("this student is absent\n")
		return
	}

	log2.Printf("see this student's desktop OK \n")
	return

}

func ScreenshotController(context echo.Context) error {

	startDate, err := strconv.ParseInt(context.FormValue("classStartDate"), 10, 64)
	if err != nil{
		return err
	}

	studentStatus := module.StudentStatus{
		StudentInfo: module.StudentInfo{
			StudentId: context.FormValue("studentId"),
		},
		Class: module.Class{
			ClassName: context.FormValue("className"),
			ClassStartDate: startDate,
		},
	}

	pic, err := context.FormFile("screenshot")
	if err != nil{
		return err
	}

	// save the screenshot at path "/screenshot"
	err = screenshot.SaveScreenshot(studentStatus.StudentName, pic)
	if err != nil{
		return err
	}

	return context.NoContent(http.StatusOK)
}

func ScreenFrozenController(context echo.Context) error {

	studentStatus := new(module.StudentStatus)
	if err := context.Bind(studentStatus); err != nil{
		return err
	}

	err := screenshot.UpdateScreenFrozenDuration(*studentStatus)
	if err != nil{
		log2.Printf("update screen frozen duration error: %v\n", err)
		return err
	}

	return context.String(http.StatusOK, "Attention! Your desktop has little change!")
}
