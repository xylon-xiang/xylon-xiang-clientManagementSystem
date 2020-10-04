package constant

// the attendance condition
const (
	NORMAL     = 4
	LATE       = 2
	LEAVEEARLY = 1
	ABSENT     = 0
)

// the homework type
const (
	SELECT    = 1
	FILLBLANK = 2
	TEXT      = 3
	FILE      = 4
)

// log status
const (
	ACCEPT  = "login:Authorization Accept"
	FAILURE = "login:Authorization Failure"
)

const (
	SIGNIN  = "sign in"
	SIGNOUT = "sign out"
)

// teacher's object_operation
const (
	IMPORT = "import"

	// for specific student
	CHANGESIGNSTATUS   = "changeStudentStatus"
	GETEACHCLASSSTATUS = "getEachClassStatus"
	GETATTENDANCERATE  = "getAttendanceRate"
	GETHOMEWORKSTATUS  = "getHomeworkStatus"
	GETCUMULATIVESCORE = "getCumulativeScore"
)
