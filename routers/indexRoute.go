package routers

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"../models"
	"github.com/gorilla/sessions"
)

var cookieStore = sessions.NewCookieStore([]byte("MySecretKey"))

const cookieName = "Role"

type sesKey int

const (
	sesKeyLogin sesKey = iota
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
	// logRequest(r)
	canEntry := []string{"Admin", "Moderator"}
	check := RoleVerify(w, r, canEntry)
	if check {
		tmp, _ := template.ParseFiles("./templates/index.html")
		tmp.Execute(w, "Hello")
	}
}

// AjaxUsers - 1 query
func AjaxUsers(w http.ResponseWriter, r *http.Request) {
	// logRequest(r)
	w.Header().Set("Content-Type", "application/json")
	Data := models.User{
		Name:  "Max",
		Money: 2332.33,
		Langs: []string{"RU", "ENG"},
	}

	respJSON, _ := json.Marshal(Data)
	w.Write(respJSON)
}

// AjaxDepartment - 2 query
func AjaxDepartment(w http.ResponseWriter, r *http.Request) {
	// logRequest(r)
	w.Header().Set("Content-Type", "application/json")
	Data := models.Department{
		Name:    `Отдел "Бухгалтерия"`,
		Number:  2,
		Country: "Russia",
	}
	respJSON, _ := json.Marshal(Data)
	fmt.Fprint(w, string(respJSON))
}

// AjaxDB - запрос к БД
func AjaxDB(w http.ResponseWriter, r *http.Request) {
	// logRequest(r)
	q := models.DirectoryValue{}
	Data := q.GetAll()
	respJSON, _ := json.Marshal(Data)
	fmt.Fprint(w, string(respJSON))
}

// AjaxGetOne - выбрать одно значение
func AjaxGetOne(w http.ResponseWriter, r *http.Request) {
	// logRequest(r)
	q := models.DirectoryValue{ID: 2816}
	Data := q.SelectOne()
	respJSON, _ := json.Marshal(Data)
	fmt.Fprint(w, string(respJSON))
}

// RoleVerify - проверка роли
func RoleVerify(w http.ResponseWriter, r *http.Request, roles []string) bool {
	type ChCh func() bool // Объявление типа функции, используем как подобие лямбды
	ses, _ := cookieStore.Get(r, cookieName)
	data := ses.Values[sesKeyLogin]
	var valid bool
	if data == nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("bad verify"))
	} else {
		// check - типа лямбда, возвращет bool, есть ли роль в списке перечисленных у роута
		var check ChCh = func() bool {
			var roleRequered bool
			for _, v := range roles {
				if data == v {
					roleRequered = true
					break
				} else {
					roleRequered = false
				}
			}
			return roleRequered
		}
		// Вызов лямбды, присваивание флага верификации
		// Если salse, то возвращаем так же ошибку
		valid = check()
		if !valid {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("bad verify"))
		}
	}
	return valid
}

// Login - типа авторизация
func Login(w http.ResponseWriter, r *http.Request) {
	gob.Register(sesKey(0))
	ses, _ := cookieStore.Get(r, cookieName)
	ses.Values[sesKeyLogin] = "Admin"
	ses.Options.HttpOnly = false
	ses.Options.Secure = false
	ses.Options.MaxAge = 86400 // 1 day exp
	// ses.Options.Domain = "localhost"
	cookieStore.Save(r, w, ses)
	w.Write([]byte("You are loggin"))
}
