package log

import "clientManagementSystem/teacher-side/module"

func CheckPassword(studentId string, studentPassword string) (isPasswordRight bool, err error) {

	filter := map[string]string{
		"StudentId": studentId,
	}

	results, err := module.FindOne(module.STUDENTINFO, filter)
	if err != nil{
		return false, err
	}

	studentInfo := results.(module.StudentInfo)

	return studentPassword == studentInfo.StudentPassword, nil


}

