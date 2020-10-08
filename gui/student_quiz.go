package gui

import (
	"clientManagementSystem/module"
	"clientManagementSystem/student-side/quiz"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	log2 "log"
)

// TODO: some button for sending content about quiz, upload homework, logout


func addQuizGroup(value *module.InputValue) *widget.Group {

	answer := "this is the answer box"

	questionTextBox := newEnterEntry()
	questionScroll := widget.NewScrollContainer(questionTextBox)
	questionScroll.SetMinSize(fyne.NewSize(800, 80))
	questionFormItem := widget.NewFormItem("please enter your question: ", questionScroll)

	questionForm := widget.NewForm(questionFormItem)

	answerLabel := widget.NewLabel(answer)

	// clean the test box and refresh
	questionForm.OnCancel = func() {
		questionTextBox.SetText("")
		questionForm.Refresh()
	}

	questionForm.OnSubmit = func() {

		question := questionTextBox.Text

		answer := quizController(value, question)
		answerLabel.SetText(answer)
		answerLabel.Refresh()

		questionForm.OnCancel()
	}

	questionGroup := widget.NewGroup("quiz function", questionForm, answerLabel)

	return questionGroup
}

func quizController(inputValues *module.InputValue, question string) string {

	answer, err := quiz.HandUp((*inputValues).StudentId, (*inputValues).StudentName, question)

	if err != nil {
		log2.Printf("hand up err: %v\n", err)
		return ""
	}

	return answer
}

