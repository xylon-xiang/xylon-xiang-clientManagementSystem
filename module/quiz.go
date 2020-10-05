package module

type QuizPostBody struct {
	StudentName string `json:"student_name"`
	StudentId   string `json:"student_id"`

	SeatPosition string `json:"seat_position"`

	QuestionContent string `json:"question_content"`
}
