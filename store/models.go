package store

// User - ..
type User struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	Role     int32  `json:"role"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// CheckPassword - ...
func (u *User) CheckPassword() bool {
	var valid bool
	DB.Session.QueryRow(`select exists(select 1 FROM users where "username" = $1 and "password" = hashpassword($2))`, u.Username, u.Password).Scan(&valid)
	return valid
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
