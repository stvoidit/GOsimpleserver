package store

// User - ..
type User struct {
	ID       int64  `json:"id"`
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

// Video - ...
type Video struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Active bool   `json:"active"`
}

// InsertVideo - ...
func (v *Video) InsertVideo() {
	_, err := DB.Session.Exec(`INSERT INTO videos (id, url)
	SELECT $1::VARCHAR, $2 WHERE NOT EXISTS
	(SELECT 1 FROM videos WHERE id = $1)`, v.ID, v.URL)
	if err != nil {
		panic(err)
	}
}
