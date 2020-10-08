package module

type StudentLogPost struct {
	StudentId       string `json:"student_id"`
	StudentName     string `json:"student_name"`
	StudentPassword string `json:"student_password"`
	ClassName       string `json:"class_name"`
	ClassStartDate  int64  `json:"class_start_date"`
}
