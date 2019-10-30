package store

import (
	"log"
	"time"

	"github.com/lib/pq"
)

const datetimetz = "2006-01-02 15:04:05.99-07"

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
	DB.QueryRow(`select exists(select 1 FROM users WHERE "username" = $1 and "password" = hashpassword($2))`, u.Username, u.Password).Scan(&valid)
	if valid {
		DB.QueryRow(`select id, "role", email
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
	Created   time.Time
}

// InsertVideo - ...
func (v *Video) InsertVideo() {
	_, err := DB.Exec(`INSERT INTO videos (id, url)
	SELECT $1::VARCHAR, $2 WHERE NOT EXISTS
	(SELECT 1 FROM videos WHERE id = $1)`, v.ID, v.URL)
	if err != nil {
		panic(err)
	}
}

// GetAllUrls - ...
func GetAllUrls() []Video {
	rows, err := DB.Query(`SELECT id, url, uploaddate, channel, title, created FROM VIDEOS`)
	if err != nil {
		log.Printf(err.Error())
	}
	var videos []Video
	for rows.Next() {
		var v Video
		rows.Scan(&v.ID, &v.URL, &v.Active, &v.Uploaded, &v.ChannelID, &v.Created)
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
	_, err := DB.Exec(`INSERT INTO public.statistic
	("views", likes, dislikes, channel, channelname, followers, video)
	VALUES($1::int, $2::int, $3::int, $4, $5, $6, $7);`, s.Views, s.Likes, s.Dislikes, s.ChannelID, s.ChannelName, s.Followers, s.Video)
	if err != nil {
		log.Printf(err.Error())
	}
}

// InsertVideo - ...
func (s *Statistic) InsertVideo(url string) bool {
	result, err := DB.Exec(`INSERT INTO public.videos
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

// StatisticSlice - ...
type StatisticSlice struct {
	ID        string `sql:"id"`
	URL       string `sql:"url"`
	Title     string `sql:"title"`
	DateSlice []time.Time
	Views     []int64
	Likes     []int64
	Dislikes  []int64
	Created   time.Time
}

// GetStat - ...
func GetStat(chanID string) []StatisticSlice {
	var stat []StatisticSlice
	rows, err := DB.Query(`select
		v.id, v.created, v.url, v.title, array_agg(s.updated), array_agg("views"), array_agg(likes), array_agg(dislikes)
		from videos as v
		join statistic as s on s.video = v.id
		where v.channel = $1
		group by 1,2,3,4
		`, chanID)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var rss StatisticSlice
		var dtslice []string
		rows.Scan(&rss.ID, &rss.Created, &rss.URL, &rss.Title, pq.Array(&dtslice), pq.Array(&rss.Views), pq.Array(&rss.Likes), pq.Array(&rss.Dislikes))
		for _, v := range dtslice {
			t, _ := time.Parse(datetimetz, v)
			rss.DateSlice = append(rss.DateSlice, t.UTC())
		}
		stat = append(stat, rss)
	}
	return stat
}

// Channel - ...
type Channel struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// GetAllChanels - ...
func GetAllChanels() []Channel {
	var channels []Channel
	rows, err := DB.Query(`select * from all_channels`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var ch Channel
		rows.Scan(&ch.Name, &ch.ID)
		channels = append(channels, ch)
	}
	return channels
}
