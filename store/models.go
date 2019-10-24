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

// GetAllUrls - ...
func GetAllUrls() []Video {
	rows, _ := DB.Session.Query(`SELECT * FROM VIDEOS`)
	var videos []Video
	for rows.Next() {
		var v Video
		rows.Scan(&v.ID, &v.URL, &v.Active)
		videos = append(videos, v)
	}
	return videos
}

// Statistic - ...
type Statistic struct {
	ID          string
	Views       int64
	Likes       int64
	Dislikes    int64
	Title       string
	ChannelID   string
	ChannelName string
	Followers   string
	UploadDate  string
	Video       string
}

// Insert - ...
func (s *Statistic) Insert() {
	_, err := DB.Session.Exec(`INSERT INTO public.statistic
	("views", likes, dislikes, title, channel, channelname, followers, uploaddate, video)
	VALUES($1::int, $2::int, $3::int, $4, $5, $6, $7, $8, $9);`, s.Views, s.Likes, s.Dislikes, s.Title, s.ChannelID, s.ChannelName, s.Followers, s.UploadDate, s.Video)
	if err != nil {
		panic(err)
	}
}
