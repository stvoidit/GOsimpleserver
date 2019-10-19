package store

// User - ..
type User struct {
	ID           int64  `json:"user_id"`
	FirstName    string `json:"firstnme"`
	MiddleName   string `json:"lastname"`
	LastName     string `json:"middlename"`
	DepartmentID int64  `json:"department"`
	Position     string `json:"position,omitempty"`
}

// Department - ...
type Department struct {
	ID             int64  `json:"departments_id"`
	Name           string `json:"name"`
	IsGeneral      bool   `json:"usgeneral"`
	Maindepartment int    `json:"maindepartment"`
}

// UserDepartments - ...
type UserDepartments struct {
	User
	Department
}
