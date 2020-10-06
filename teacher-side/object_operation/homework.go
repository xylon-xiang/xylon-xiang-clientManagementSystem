package object_operation

import (
	"clientManagementSystem/module"
	"clientManagementSystem/teacher-side/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func GetHomeworkInfoTest() ([]module.HomeworkInfo, error) {

	file, err := os.Open("test/homeworkTest.json")
	if err != nil {
		return nil, err
	}

	bytesStream, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var homeworkInfo []module.HomeworkInfo
	err = json.Unmarshal(bytesStream, &homeworkInfo)
	if err != nil {
		return nil, err
	}

	return homeworkInfo, nil
}

// save the []HomeworkInfo into database and broadcast them
// the input homework info is the standard homework info, which means it contains the standard answer
func PublishHomeworkInfo(hub *util.Hub, homeworkInfo []module.HomeworkInfo) (success bool, err error) {


	err = util.Save(util.HOMEWORKINFO, homeworkInfo)
	if err != nil {
		return false, err
	}

	// remove the standard answer
	for key, _ := range homeworkInfo {
		homeworkInfo[key].HomeworkAnswer = ""
		homeworkInfo[key].HomeworkScore = 0
	}


	homeworkStr := fmt.Sprintf("%v", homeworkInfo)

	// broadcast the homework info
	hub.Broadcast <- []byte(homeworkStr)

	return true, nil

}

// this function is used after teacher correct some questions and give his subject points
func ChangeHomeworkStatus(studentStatus module.StudentStatus) error {

	startDate := strconv.FormatInt(studentStatus.ClassStartDate, 10)
	filter := map[string]string{
		"StudentId": studentStatus.StudentId,
		"ClassName": studentStatus.ClassName,
		"ClassStartDate": startDate,
	}

	result, err := util.FindOne(util.STUDENTSTATUS, filter)
	if err != nil{
		return err
	}

	status := result.(module.StudentStatus)

	for key, _ := range status.HomeworksInfo{

		for _, homework := range studentStatus.HomeworksInfo{

			if status.HomeworksInfo[key].HomeworkTitle == homework.HomeworkTitle ||
				status.HomeworksInfo[key].HomeworkType == homework.HomeworkType{
				status.HomeworksInfo[key].HomeworkScore = homework.HomeworkScore
			}
		}
	}

	return nil
}