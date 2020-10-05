package log

import (
	"clientManagementSystem/module"
	"clientManagementSystem/teacher-side/util"
)

func CheckPassword(studentId string, studentPassword string) (isPasswordRight bool, err error) {

	filter := map[string]string{
		"StudentId": studentId,
	}

	results, err := util.FindOne(util.STUDENTINFO, filter)
	if err != nil{
		return false, err
	}

	studentInfo := results.(module.StudentInfo)

	return studentPassword == studentInfo.StudentPassword, nil


}

