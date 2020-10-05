package object_operation

import (
	"clientManagementSystem/module"
	"clientManagementSystem/teacher-side/util"
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

	result, err := util.FindOne(util.STUDENTSTATUS, filter)
	if err != nil{
		return false, err
	}

	studentStatus := result.(module.StudentStatus)
	studentStatus.SignStatus = targetSignStatus

	err = util.UpdateOne(util.STUDENTSTATUS, studentStatus)
	if err != nil{
		return false, err
	}

	return true, nil
}
