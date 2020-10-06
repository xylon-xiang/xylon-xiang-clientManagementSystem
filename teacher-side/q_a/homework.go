package q_a

import (
	"clientManagementSystem/module"
	"clientManagementSystem/teacher-side/constant"
	"clientManagementSystem/teacher-side/util"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

// correct homework by exact match the int or string
// pass the homework whose type is text/file
func AutoCorrectHomeWork(homeworkInfo *[]module.HomeworkInfo) error {

	results, err := util.FindAll(util.HOMEWORKINFO, "")
	if err != nil {
		return err
	}

	standardHomeworkInfo := results.([]module.HomeworkInfo)

	// auto correct the homework whose type is select or fill blank
	for _, answer := range standardHomeworkInfo {

		for key, homework := range *homeworkInfo {

			// if over ddl, then , no score
			if time.Now().Unix() > answer.HomeworkDDL {
				(*homeworkInfo)[key].HomeworkScore = constant.HOMEWORKSCOREZERO
				continue
			}

			// if it is the homework type is text/file, pass it
			if homework.HomeworkType == constant.TEXT ||
				homework.HomeworkType == constant.FILE {
				continue
			}

			// if all matched, then they are the same questions
			if answer.HomeworkType == homework.HomeworkType &&
				answer.HomeworkTitle == homework.HomeworkTitle {

				// auto correct it
				if answer.HomeworkAnswer == homework.HomeworkAnswer {
					(*homeworkInfo)[key].HomeworkScore = constant.HOMEWORKSCOREMAX
				} else {
					(*homeworkInfo)[key].HomeworkScore = constant.HOMEWORKSCOREZERO
				}

			}
		}
	}

	return nil
}

// this function is used to update homeworkInfo in studentStatus,
// always used after correct homework
func UpdateStudentHomeworkStatus(studentStatus module.StudentStatus) error {

	startTime := strconv.FormatInt(studentStatus.ClassStartDate, 10)
	filter := map[string]string{
		"StudentId":      studentStatus.StudentId,
		"ClassName":      studentStatus.ClassName,
		"ClassStartDate": startTime,
	}
	results, err := util.FindOne(util.STUDENTSTATUS, filter)
	if err != nil {
		return err
	}

	status := results.(module.StudentStatus)

	status.HomeworksInfo = studentStatus.HomeworksInfo

	err = util.UpdateOne(util.STUDENTSTATUS, status)
	if err != nil {
		return err
	}

	return nil
}

// this function is used to pretreated the homework,
func SaveFileQuestionAsFile(studentStatus module.StudentStatus,
	questionTitle string, file *multipart.FileHeader) error {

	uri := util.GetFilePath(studentStatus.StudentId, studentStatus.StudentName, file.Filename)

	src, err := file.Open()
	if err != nil{
		return err
	}
	dst, err := os.Create(uri)
	if err != nil{
		return err
	}

	// save the file
	if _, err = io.Copy(dst,src); err != nil{
		return err
	}

	startDate := strconv.FormatInt(studentStatus.ClassStartDate, 10)
	filter := map[string]string{
		"StudentId": studentStatus.StudentId,
		"ClassName": studentStatus.ClassName,
		"ClassStartDate": startDate,
	}
	results, err := util.FindOne(util.STUDENTSTATUS, filter)
	if err != nil{
		return err
	}

	status := results.(module.StudentStatus)

	for _,  homework := range status.HomeworksInfo{
		if homework.HomeworkTitle == questionTitle {
			homework.HomeworkAnswer = uri
		}
	}

	return nil
}
