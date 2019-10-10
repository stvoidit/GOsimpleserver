package routers

import (
	"database/sql"
	"fmt"
	"os"

	// driver
	_ "github.com/lib/pq"

	"gopkg.in/ini.v1"
)

var db *sql.DB

func init() {
	engine := configDB()
	db = engine.Conn()
}

// DB - session object
type DB struct {
	ConnStr string
}

func configDB() DB {
	cfg, err := ini.Load("config.ini")
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
	d := DB{ConnStr: conn}
	return d
}

// Conn - connestion
func (c DB) Conn() *sql.DB {
	db, _ := sql.Open("postgres", c.ConnStr)
	return db
}
