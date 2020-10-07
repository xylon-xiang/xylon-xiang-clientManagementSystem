package module

type InputValue struct {

	// student info
	StudentId       string
	StudentName     string
	StudentPassword string

	ClassName string

	// the input time format, such as 2019-11-12 12:25
	ClassDate string

	// the unix timestamp
	ClassStartDate int64
}
