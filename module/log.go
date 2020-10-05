package module

type StudentLogPost struct {
	StudentPassword string `json:"student_password"`
	ClassName       string `json:"class_name"`
	ClassStartDate  int64  `json:"class_start_date"`
}
