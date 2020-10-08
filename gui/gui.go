package gui

import (
	"clientManagementSystem/module"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type Windows struct {
	Application fyne.App
	WindowsMap  map[string]*fyne.Window
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
		WindowsMap:  make(map[string]*fyne.Window),
	}
	//
	//windows.Application.Run()

	return &windows
}

// create a new windows, and add it into `windows`
func (windows *Windows) RegisterWindows(windowsName string, dx, dy int) {

	win := (*windows).Application.NewWindow(windowsName)
	win.Resize(fyne.NewSize(dx, dy))

	(*windows).WindowsMap[windowsName] = &win
}

func (windows *Windows) AddAlarmItem(windowsName string, alarmInfo string) {

	alarm := widget.NewLabel(alarmInfo)

	w := (*windows).WindowsMap[windowsName]
	(*w).SetContent(widget.NewVBox(
		alarm,
		))


}
