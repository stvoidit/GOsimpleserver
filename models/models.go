package models

import (
	"database/sql"
	"errors"
	"fmt"

	_ "gopkg.in/goracle.v2"
	"gopkg.in/ini.v1"
)

type SQLquery struct {
	ID    int    `json:"id"`
	VALUE string `json:"value,omitempty"`
}

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

// DataBaseCfg - соадинение с БД, чтение конфига config.ini
func DataBaseCfg() string {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		myerrorMessage := errors.New("Нет файла конфигураций")
		panic(myerrorMessage)
	}
	serverSection := cfg.Section("server")
	host := serverSection.Key("host")
	port := serverSection.Key("port")
	user := serverSection.Key("user")
	password := serverSection.Key("password")
	dbname := serverSection.Key("dbname")
	conn := fmt.Sprintf("%s/%s@%s:%s/%s", user, password, host, port, dbname)
	return conn
}

// GetValues - пример запроса в БД Oracle
func GetValues() []SQLquery {
	connString := DataBaseCfg()
	db, err := sql.Open("goracle", connString)
	if err != nil {
		fmt.Println(err.Error())
	}
	rows, err := db.Query("select ID, VALUE from UNIO.T_DIRECTORY_VALUE")
	if err != nil {
		fmt.Println(err.Error())
		// return
	}

	listQuery := []SQLquery{}
	for rows.Next() {
		rowdata := SQLquery{}
		rows.Scan(&rowdata.ID, &rowdata.VALUE)
		listQuery = append(listQuery, rowdata)
	}
	return listQuery
}
