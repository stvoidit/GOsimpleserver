package models

type User struct {
	Name  string
	Money float64
	Langs []string
}

type Department struct {
	Name    string
	Number  int32
	Country string
}
