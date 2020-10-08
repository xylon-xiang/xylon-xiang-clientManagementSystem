package gui

import (
	"clientManagementSystem/module"
	"fyne.io/fyne/widget"
)

func (windows *Windows) StudentSideMainWindows(windowsName string,
	value *module.InputValue, homework *chan string) {

	quizGroup := addQuizGroup(value)
	homeworkGroup := addHomeworkGroup(value, homework)

	w := windows.WindowsMap[windowsName]
	(*w).SetContent(widget.NewVBox(
		quizGroup,
		homeworkGroup,
		))

}