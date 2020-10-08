package gui

import (
	"clientManagementSystem/module"
	"clientManagementSystem/student-side/q_a"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	log2 "log"
	"strconv"
)

func addHomeworkGroup(value *module.InputValue, broadcastHomework *chan string) *widget.Group {

	var homeworks []module.HomeworkInfo

	var formList []*widget.Form

	var propertiesMap []map[string]string

	homeworkContentLabel := widget.NewLabel("this will show the teacher's broadcast broadcastHomework")

	// set the broadcastHomework content when teacher broadcast the broadcastHomework message
	go func() {
		for {
			homeworkContent := <-(*broadcastHomework)
			homeworkContentLabel.SetText(homeworkContent)
			homeworkContentLabel.Refresh()
		}
	}()

	// a group containing the broadcastHomework forms
	homeworkSubmitGroup := widget.NewGroup("broadcastHomework submit group")
	homeworkSubmitScroller := widget.NewScrollContainer(homeworkSubmitGroup)
	homeworkSubmitScroller.SetMinSize(fyne.NewSize(800, 400))

	// when clicking the button, add a broadcastHomework submit form
	AddHomeworkItemButton := widget.NewButton("add broadcastHomework submit form", func() {

		propertyMap := make(map[string]string)

		homeworkSubmitForm := homeworkSubmitForm(&propertyMap)

		propertiesMap = append(propertiesMap, propertyMap)

		formList = append(formList, homeworkSubmitForm)
		homeworkSubmitGroup.Append(homeworkSubmitForm)
	})

	// when clicking the button, send the homework to teacher
	submitButton := widget.NewButton("submit broadcastHomework: ", func() {

		for key, _ := range formList {

			//add the homework
			homeworkType, _ := strconv.Atoi(propertiesMap[key]["homeworkType"])

			// FIXME: fix the homework info when user enter
			homework := module.HomeworkInfo{
				HomeworkTitle:  propertiesMap[key]["homeworkTitle"],
				HomeworkType:   homeworkType,
				HomeworkAnswer: propertiesMap[key]["homeworkAnswer"],
			}

			homeworks = append(homeworks, homework)
		}

		UploadHomeworkController(value, &homeworks)
	})

	homeworkGroup := widget.NewGroup("broadcastHomework group",
		homeworkContentLabel,
		homeworkSubmitScroller, AddHomeworkItemButton,
		submitButton,
	)

	return homeworkGroup

}

func homeworkSubmitForm(propertyMap *map[string]string) *widget.Form {

	homeworkTitleTextBox := newEnterEntry()
	homeworkTitleFormItem := widget.NewFormItem("homework title: ", homeworkTitleTextBox)

	homeworkTypeTextBox := newEnterEntry()
	homeworkTypeFormItem := widget.NewFormItem("homework type: ", homeworkTypeTextBox)

	homeworkAnswerTextBox := newEnterEntry()
	homeworkAnswerFormItem := widget.NewFormItem("homework answer: ", homeworkAnswerTextBox)

	homeworkForm := widget.NewForm(homeworkTitleFormItem, homeworkTypeFormItem, homeworkAnswerFormItem)

	homeworkForm.OnSubmit = func() {
		(*propertyMap)["homeworkTitle"] = homeworkTitleTextBox.Text
		(*propertyMap)["homeworkType"] = homeworkTypeTextBox.Text
		(*propertyMap)["homeworkAnswer"] = homeworkAnswerTextBox.Text

	}

	return homeworkForm
}

func UploadHomeworkController(value *module.InputValue, homeworks *[]module.HomeworkInfo) {

	status := module.StudentStatus{
		Class: module.Class{
			ClassName:      (*value).ClassName,
			ClassStartDate: (*value).ClassStartDate,
		},
		StudentInfo: module.StudentInfo{
			StudentId:   (*value).StudentId,
			StudentName: (*value).StudentName,
		},
		HomeworksInfo: *homeworks,
	}

	response, err := q_a.UploadHomework(status)
	if err != nil {
		log2.Printf("upload homework err: %v\n", err)
		return
	}

	log2.Println(response)

}
