package gui

import (
	"clientManagementSystem/module"
	"fyne.io/fyne/widget"
	//"time"
)

const longForm = "Jan 2, 2006 3:04 (MST)"

// afterSubmit is a function list used after clicking the submit button
func (windows *Windows) AddLoginForm(windowName string, input *module.InputValue, afterSubmit ...func()) {


	timeLabel := widget.NewLabel("Please attention: time should be such a format : \n" + longForm)

	studentNameTextBox := newEnterEntry()
	studentNameFormItem := widget.NewFormItem("your name: ", studentNameTextBox)

	studentIdTextBox := newEnterEntry()
	studentIdFormItem := widget.NewFormItem("your student ID: ", studentIdTextBox)

	studentPasswordTextBox := widget.NewPasswordEntry()
	studentPasswordFormItem := widget.NewFormItem("your password: ", studentPasswordTextBox)

	classNameTextBox := newEnterEntry()
	classNameFormItem := widget.NewFormItem("this class name: ", classNameTextBox)

	classStartDateTextBox := newEnterEntry()
	classStartDateFormItem := widget.NewFormItem(
		"this class start time: ", classStartDateTextBox)

	loginForm := widget.NewForm(
		studentNameFormItem, studentIdFormItem, studentPasswordFormItem,
		classNameFormItem, classStartDateFormItem,
	)

	loginForm.OnSubmit = func() {
		(*input).StudentId = studentIdTextBox.Text
		(*input).StudentName = studentNameTextBox.Text
		(*input).StudentPassword = studentPasswordTextBox.Text
		(*input).ClassName = classNameTextBox.Text
		(*input).ClassDate = classStartDateTextBox.Text

		// parse time string
		// this is the standard method

		// loc, _ := time.LoadLocation("asia/shanghai")
		// t, _ := time.ParseInLocation(longForm, (*input).ClassDate, loc)
		//(*input).ClassStartDate = t.Unix()

		// FIXME: this time is just for test is just for test
		(*input).ClassStartDate = int64(1601790639)


		for _, value := range afterSubmit{
			value()
		}

		// init the textBox value and refresh
		loginForm.OnCancel()
	}

	loginForm.OnCancel = func() {
		studentIdTextBox.SetText("")
		studentNameTextBox.SetText("")
		studentPasswordTextBox.SetText("")
		classNameTextBox.SetText("")
		classStartDateTextBox.SetText("")
		loginForm.Refresh()
	}

	loginGroup := widget.NewGroup("login input", timeLabel ,loginForm)

	w := (*windows).WindowsMap[windowName]
	(*w).SetContent(widget.NewVBox(
		loginGroup,
	))
}
