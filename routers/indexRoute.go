package routers

import (
	"net/http"

	"../store"
)

var db = store.DB

// Users = is '/'
// Передача параметров в template
func Users(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()["user"]
	result := db.AllUsers(args)
	Jsonify(w, result, 200)

}

// Departments - ...
func Departments(w http.ResponseWriter, r *http.Request) {
	result := db.AllDepartments()
	Jsonify(w, result, 200)
}

// UsersDepartments - ...
func UsersDepartments(w http.ResponseWriter, r *http.Request) {
	result := db.AllUsersDepartments()
	Jsonify(w, result, 200)
}
