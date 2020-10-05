package object_operation

import (
	"clientManagementSystem/module"
	"clientManagementSystem/teacher-side/constant"
	"clientManagementSystem/teacher-side/util"
	"encoding/json"
	"io/ioutil"
	log2 "log"
	"os"
)

func importClassInfo(classInfo module.ClassInfo) (success bool, err error) {

	studentStatus := make([]module.StudentStatus, len(classInfo.StudentsInfo))

	for key, value := range classInfo.StudentsInfo{
		studentStatus[key].ClassName = classInfo.ClassName
		studentStatus[key].ClassStartDate = classInfo.ClassStartDate
		studentStatus[key].ClassOverDate = classInfo.ClassOverDate


		// init the sign status as absent
		studentStatus[key].SignStatus = constant.ABSENT

		studentStatus[key].StudentInfo = value
		studentStatus[key].HomeworksInfo = nil
	}

	// save the student status
	err  = util.Save(util.STUDENTSTATUS, studentStatus)
	if err != nil{
		return false, err
	}

	// save the student info
	err = util.Save(util.STUDENTINFO, classInfo.StudentsInfo)
	if err != nil{
		return false, err
	}

	return true, nil
}


// this function is just for test
func ImportClassTestCase() {

	// the file uri should be specified by teacher
	// the test.json is just for test
	file, err := os.Open("test/test.json")

	if err != nil {
		log2.Printf("file open error: %v \n", err)
		return
	}

	byteStream, err := ioutil.ReadAll(file)
	if err != nil {
		log2.Printf("file read error: %v \n", err)
		return
	}

	var classInfo module.ClassInfo

	err = json.Unmarshal(byteStream, &classInfo)
	if err != nil {
		log2.Printf("json unmarshal error: %v \n", err)
		return
	}

	success, err := importClassInfo(classInfo)
	if err != nil {
		log2.Printf("import into database error: %v \n", err)
		return
	}

	if !success {
		log2.Println("unknown error")
		return
	}

	log2.Println("import success")
}
