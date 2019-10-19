package store

import (
	"database/sql"
	"fmt"
	"os"

	// driver
	_ "github.com/lib/pq"
	"gopkg.in/ini.v1"
)

// DB - ...
var DB = Store{}

func init() {
	DB.Connect()
}

// Store - ...
type Store struct {
	Session *sql.DB
}

// Connect - ....
func (s *Store) Connect() {
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	database := cfg.Section("database")
	host, _ := database.GetKey("host")
	port, _ := database.GetKey("port")
	login, _ := database.GetKey("login")
	password, _ := database.GetKey("password")
	dbname, _ := database.GetKey("dbname")
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, login, password, dbname)
	db, _ := sql.Open("postgres", conn)
	s.Session = db
}