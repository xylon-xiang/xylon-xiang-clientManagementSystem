package object_operation

import (
	"clientManagementSystem/teacher-side/constant"
	"clientManagementSystem/teacher-side/module"
)

func ImportClassInfo(classInfo module.ClassInfo) (success bool, err error) {

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
	err  = module.Save(module.STUDENTSTATUS, studentStatus)
	if err != nil{
		return false, err
	}

	// save the student info
	err = module.Save(module.STUDENTINFO, classInfo.StudentsInfo)
	if err != nil{
		return false, err
	}

	return true, nil
}


