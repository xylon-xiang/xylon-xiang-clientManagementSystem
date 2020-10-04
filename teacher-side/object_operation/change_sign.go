package object_operation

import (
	"clientManagementSystem/teacher-side/module"
	"strconv"
)

func ChangeSignStatus(studentId string, className string,
	classStartDate int64, targetSignStatus int) (bool, error) {

	startTime := strconv.FormatInt(classStartDate, 10)

	filter := map[string]string{
		"StudentId": studentId,
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
