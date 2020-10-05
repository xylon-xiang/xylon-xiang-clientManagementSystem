package quiz

import (
	"bytes"
	"clientManagementSystem/config"
	"clientManagementSystem/module"
	"clientManagementSystem/student-side/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HandUp(studentId string, studentName string, question string) (string, error) {

	// Post URL
	hostUrl := config.Config.APIConfig.TeacherHost +
		config.Config.APIConfig.StudentHandUpAPI.Path

	urlStr := util.GetRealUrl(hostUrl, studentId)

	//POST Body
	seatInfo := util.GetSeatInfo()
	quizBody := module.QuizPostBody{
		StudentName:     studentName,
		StudentId:       studentId,
		SeatPosition:    seatInfo,
		QuestionContent: question,
	}

	requestBody, err := json.Marshal(quizBody)
	if err != nil{
		return "", err
	}

	resp, err  := http.Post(urlStr, "application/json", bytes.NewBuffer(requestBody))
	if err != nil{
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return "", err
	}

	return string(body), nil

}
