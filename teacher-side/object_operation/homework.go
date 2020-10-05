package object_operation

import (
	"clientManagementSystem/module"
	"clientManagementSystem/teacher-side/constant"
	"clientManagementSystem/teacher-side/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func getHomeworkInfoTest() ([]module.HomeworkInfo, error) {

	file, err := os.Open("test/homeworkTest.json")
	if err!= nil{
		return nil, err
	}

	bytesStream, err := ioutil.ReadAll(file)
	if err != nil{
		return nil, err
	}

	var homeworkInfo []module.HomeworkInfo
	err = json.Unmarshal(bytesStream, &homeworkInfo)
	if err != nil{
		return nil, err
	}

	return homeworkInfo, nil
}

// save the []HomeworkInfo into database and broadcast them
// the input homework info is the standard homework info, which means it contains the standard answer
func PublishHomeworkInfo(hub *util.Hub, homeworkInfo []module.HomeworkInfo) (success bool, err error) {

	// remove the standard answer
	for _, value := range homeworkInfo{
		value.HomeworkAnswer = ""
		value.HomeworkScore = 0
	}

	homeworkStr := fmt.Sprintf("%v", homeworkInfo)

	err = util.Save(util.HOMEWORKINFO, homeworkInfo)
	if err != nil{
		return false, err
	}

	// broadcast the homework info
	hub.Broadcast <- []byte(homeworkStr)

	return true, nil

}

// correct homework by exact match the int or string
// pass the homework whose type is text/file
func AutoCorrectHomeWork(homeworkInfo *[]module.HomeworkInfo) error {

	results, err := util.FindAll(util.HOMEWORKINFO, "")
	if err != nil{
		return err
	}

	standardHomeworkInfo := results.([]module.HomeworkInfo)

	// auto correct the homework whose type is select or fill blank
	for _, answer := range standardHomeworkInfo{

		for _, homework := range *homeworkInfo{

			// if over ddl, then , no score
			if time.Now().Unix() > answer.HomeworkDDL {
				homework.HomeworkScore = constant.HOMEWORKSCOREZERO
				continue
			}

			// if it is the homework type is text/file, pass it
			if homework.HomeworkType == constant.TEXT ||
				homework.HomeworkType == constant.FILE{
				continue
			}

			// if all matched, then they are the same questions
			if answer.HomeworkType == homework.HomeworkType &&
				answer.HomeworkTitle == homework.HomeworkTitle{

				// auto correct it
				if answer.HomeworkAnswer == homework.HomeworkAnswer{
					homework.HomeworkScore = constant.HOMEWORKSCOREMAX
				}else {
					homework.HomeworkScore = constant.HOMEWORKSCOREZERO
				}

			}
		}
	}

	return nil
}


// TODO: update student homework status in database; (just write teacher-side, no student-side)
func UpdateStudentHomeworkStatus(homeworkInfo []module.HomeworkInfo) error {



	return nil
}

