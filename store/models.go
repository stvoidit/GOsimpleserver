package store

// User - ..
type User struct {
	ID           int64
	FirstName    string
	MiddleName   string
	LastName     string
	DepartmentID int64
	Position     string
}

// Department - ...
type Department struct {
	ID             int64
	Name           string
	IsGeneral      bool
	Maindepartment int
}
