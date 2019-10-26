package store

import (
	"log"
)

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
	DB.Session.QueryRow(`select exists(select 1 FROM users WHERE "username" = $1 and "password" = hashpassword($2))`, u.Username, u.Password).Scan(&valid)
	if valid {
		DB.Session.QueryRow(`select id, "role", email
		FROM users
		WHERE "username" = $1 and "password" = hashpassword($2)`,
			u.Username, u.Password).Scan(&u.ID, &u.Role, &u.Email)
	}
	return valid
}

// Video - ...
type Video struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Active    bool   `json:"active"`
	ChannelID string `json:"channel"`
	Uploaded  string `json:"uploaded"`
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
		rows.Scan(&v.ID, &v.URL, &v.Active, &v.Uploaded, &v.ChannelID)
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
	("views", likes, dislikes, channel, channelname, followers, video)
	VALUES($1::int, $2::int, $3::int, $4, $5, $6, $7);`, s.Views, s.Likes, s.Dislikes, s.ChannelID, s.ChannelName, s.Followers, s.Video)
	if err != nil {
		log.Printf(err.Error())
	}
}

// InsertVideo - ...
func (s *Statistic) InsertVideo(url string) bool {
	result, err := DB.Session.Exec(`INSERT INTO public.videos
	(id, url, uploaddate, channel, title)
	select $1::varchar, $2, $3, $4, $5
	where not exists(select 1 from videos where id = $1::varchar);`, s.ID, url, s.UploadDate, s.ChannelID, s.Title)
	if err != nil {
		log.Printf(err.Error())
	}
	if ok, _ := result.RowsAffected(); ok == 0 {
		return false
	}
	return true

}
