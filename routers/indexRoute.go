package routers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"../models"
)

// logRequest - логгер/принтер в консоль данных о запросе
func logRequest(r *http.Request) {
	const dtFormat = "2006-01-02 03:04:05"
	correntTime := time.Now().Format(dtFormat)
	fmt.Println(r.Method, r.RemoteAddr, correntTime, r.RequestURI, r.Header["User-Agent"][0])
}

// IndexRoute = is '/'
// Передача параметров в template
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	tmp, _ := template.ParseFiles("./templates/index.html")
	tmp.Execute(w, "Hello")
}

// AjaxUsers - 1 query
func AjaxUsers(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	w.Header().Set("Content-Type", "application/json")
	Data := models.User{
		Name:  "Max",
		Money: 2332.33,
		Langs: []string{"RU", "ENG"},
	}
	respJSON, _ := json.Marshal(Data)
	fmt.Fprint(w, string(respJSON))
}

// AjaxDepartment - 2 query
func AjaxDepartment(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	w.Header().Set("Content-Type", "application/json")
	Data := models.Department{
		Name:    "Бухгалтерия",
		Number:  2,
		Country: "Russia",
	}
	respJSON, _ := json.Marshal(Data)
	fmt.Fprint(w, string(respJSON))
}

// AjaxDB - запрос к БД
func AjaxDB(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	Data := models.GetValues()
	respJSON, _ := json.Marshal(Data)
	fmt.Fprint(w, string(respJSON))
}
