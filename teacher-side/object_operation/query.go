package object_operation

import (
	"clientManagementSystem/teacher-side/module"
	"encoding/json"
)

/*
	all the operation in this file is aimed to one specific person
*/

type EachClassStatus struct {
	module.Class

	StudentId   string `json:"student_id"`
	StudentName string `json:"student_name"`

	SignStatus int `json:"sign_status"`
}

func QueryStudentStatus(studentId string, className string) ([]module.StudentStatus, error) {

	results, err := module.FindAll(module.STUDENTSTATUS, studentId, className)
	if err != nil {
		return nil, err
	}

	return results.([]module.StudentStatus), nil
}


// output the results as a json stream
func QueryEachClassStatus(studentStatuses []module.StudentStatus) (jsonByteStream []byte, err error){

	var EachClassStatuses []EachClassStatus

	for _, studentStatus := range studentStatuses{
		eachClassStatus := EachClassStatus{
			Class:       studentStatus.Class,
			StudentId:   studentStatus.StudentId,
			StudentName: studentStatus.StudentName,
			SignStatus:  studentStatus.SignStatus,
		}
		EachClassStatuses = append(EachClassStatuses, eachClassStatus)
	}

	jsonByteStream, err = json.Marshal(EachClassStatuses)
	if err != nil{
		return nil, err
	}

	return jsonByteStream, nil

}


// apart from absent situation, all situations are as attendance
func QueryAttendanceRate(studentStatus []module.StudentStatus) (rate float64) {
	count := float64(0)
	length := float64(0)

	for _, value := range studentStatus{
		length++
		if value.SignStatus != 0{
			count++
		}
	}

	return count / length
}

// TODO:
func QueryHomeworkStatus(studentStatus []module.StudentStatus) {

}

// TODO:
func QueryCumulativeScore(studentStatus []module.StudentStatus) {

}
