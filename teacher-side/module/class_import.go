package module

// teacher upload data format
type ClassInfo struct {
	ClassName    string        `json:"class_name"`
	ClassDate    int64         `json:"class_date"`
	StudentsInfo []StudentInfo `json:"students_info"`
}

type StudentInfo struct {
	StudentId       string `json:"student_id"`
	StudentName     string `json:"student_name"`
	StudentPassword string `json:"student_password"`
}

// database saving format
type StudentStatus struct {
	// class info
	ClassName      string `json:"class_name"`
	ClassStartDate int64  `json:"class_start_date"`
	ClassOverDate  int64  `json:"class_over_date"`

	// student info
	StudentId       string `json:"student_id"`
	StudentName     string `json:"student_name"`
	StudentPassword string `json:"student_password"`

	// sign info
	SignInDate  int64 `json:"sign_in_date"`
	SignOutDate int64 `json:"sign_out_date"`
	SignStatus  int   `json:"sign_status"`

	// homework info
	HomeworksInfo []HomeworkInfo `json:"homeworks_info"`
}

type HomeworkInfo struct {
	HomeworkTitle  string `json:"homework_title"`
	HomeworkType   int    `json:"homework_type"`
	HomeworkAnswer string `json:"homework_answer"`
	HomeworkScore  int    `json:"homework_score"`
}
