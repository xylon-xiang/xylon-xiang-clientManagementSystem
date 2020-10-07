package screenshot

import (
	"bytes"
	"clientManagementSystem/config"
	"clientManagementSystem/module"
	"clientManagementSystem/student-side/util"
	"encoding/json"
	"net/http"
)

func ReportFrozenToTeacher(studentId string) error {

	body := module.StudentStatus{
		Class: module.Class{
			ClassName: "Software Engineering 1210",
			ClassStartDate: int64(1601790639),
		},
		StudentInfo: module.StudentInfo{
			StudentId: "U0000",
		},
	}

	client := http.Client{}

	bodyJson, err := json.Marshal(body)
	if err != nil{
		return err
	}

	host := config.Config.APIConfig.TeacherHost +
		config.Config.APIConfig.ScreenshotAPI.UpdateFrozenDurationPath
	uri := util.GetRealUrl(host, studentId)

	req, err := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(bodyJson))
	if err != nil{
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, err = client.Do(req)
	if err != nil{
		return err
	}

	return nil
}
