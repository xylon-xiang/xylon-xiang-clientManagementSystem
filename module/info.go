package module

//////////////////////////////////////
// codes below are as import format //
//////////////////////////////////////

// teacher upload data format
type ClassInfo struct {
	ClassName      string        `json:"class_name" bson:"class_name"`
	ClassStartDate int64         `json:"class_start_date" bson:"class_start_date"`
	ClassOverDate  int64         `json:"class_over_date" bson:"class_over_date"`
	StudentsInfo   []StudentInfo `json:"students_info" bson:"students_info"`
}

type StudentInfo struct {
	StudentId       string `json:"student_id" bson:"student_id"`
	StudentName     string `json:"student_name" bson:"student_name"`
	StudentPassword string `json:"student_password" bson:"student_password"`
}

//////////////////////////////////////
// codes above are as import format //
//////////////////////////////////////

// database saving format

type StudentStatus struct {
	// class info
	Class

	// student info
	StudentInfo

	// sign info
	SignInDate  int64 `json:"sign_in_date" bson:"sign_in_date"`
	SignOutDate int64 `json:"sign_out_date" bson:"sign_out_date"`
	SignStatus  int   `json:"sign_status" bson:"sign_status"`

	// homework info
	HomeworksInfo []HomeworkInfo `json:"homeworks_info" bson:"homeworks_info"`
}

type Class struct {
	ClassName      string `json:"class_name" bson:"class_name"`
	ClassStartDate int64  `json:"class_start_date" bson:"class_start_date"`
	ClassOverDate  int64  `json:"class_over_date" bson:"class_over_date"`
}

// if type is file, then answer will be the file path
type HomeworkInfo struct {
	HomeworkTitle  string `json:"homework_title" bson:"homework_title"`
	HomeworkType   int    `json:"homework_type" bson:"homework_type"`
	HomeworkAnswer string `json:"homework_answer" bson:"homework_answer"`
	HomeworkScore  int    `json:"homework_score" bson:"homework_score"`
	HomeworkDDL    int64  `json:"homework_ddl" bson:"homework_ddl"`
}
