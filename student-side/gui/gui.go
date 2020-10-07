package gui

import (
	"clientManagementSystem/module"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type Windows struct {
	Application fyne.App
	WindowsMap map[string]*fyne.Window
}

type enterEntry struct {
	widget.Entry
}

func newEnterEntry() *enterEntry {
	entry := &enterEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func NewGUI(input *module.InputValue) *Windows {

	windows := Windows{
		Application: app.New(),
		WindowsMap: make(map[string]*fyne.Window),
	}

	windows.Application.Run()

	return &windows
}

func (windows *Windows) SetLoginWindows() {

	loginWindowsName := "login"

	windows.registerWindows(loginWindowsName, 400, 400)

	// set the login windows size




}




func (windows *Windows) registerWindows(windowsName string, dx, dy int) {

	win := (*windows).Application.NewWindow(windowsName)
	win.Resize(fyne.NewSize(dx, dy))

	(*windows).WindowsMap[windowsName] = &win
}

func (windows *Windows) addLoginForm(windowName string, input *module.InputValue) {

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

		//(*input).ClassStartDate = int64(3)
	}

	loginForm.OnCancel = func() {
		studentIdTextBox.SetText("")
		studentNameTextBox.SetText("")
		studentPasswordTextBox.SetText("")
		classNameTextBox.SetText("")
		classStartDateTextBox.SetText("")
		loginForm.Refresh()
	}

	loginGroup := widget.NewGroup("login input", loginForm)

	w := (*windows).WindowsMap[windowName]
	(*w).SetContent(widget.NewVBox(
		loginGroup,
		))
}

