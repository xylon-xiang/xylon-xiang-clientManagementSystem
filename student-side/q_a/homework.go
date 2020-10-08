package q_a

import (
	"bytes"
	"clientManagementSystem/config"
	"clientManagementSystem/module"
	"clientManagementSystem/student-side/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// this function is just for getting the test data
func GetStudentStatus() (*module.StudentStatus, error) {

	file, err := os.Open("test/uploadStatus.json")
	if err != nil{
		return nil, err
	}

	body, err := ioutil.ReadAll(file)
	if err != nil{
		return nil, err
	}

	status := new(module.StudentStatus)
	if err := json.Unmarshal(body, status); err != nil{
		return nil, err
	}

	return status, nil
}

func UploadHomework(status module.StudentStatus) (string, error) {

	hostUrl := config.Config.APIConfig.TeacherHost +
		config.Config.APIConfig.HomeworkAPI.Path

	url := util.GetRealUrl(hostUrl, status.StudentId)

	requestBody, err := json.Marshal(status)
	if err != nil{
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil{
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return "", err
	}

	return string(body), nil
}

// TODO: send the file to homework file receive API