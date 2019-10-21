package routers

import (
	"net/http"

	"../store"
)

var db = store.DB

// Users = ...
func Users(w http.ResponseWriter, r *http.Request) {
	result := "123"
	Jsonify(w, result, 200)

}

// Departments - ...
func Departments(w http.ResponseWriter, r *http.Request) {
	result := "123"
	Jsonify(w, result, 200)
}

// UsersDepartments - ...
func UsersDepartments(w http.ResponseWriter, r *http.Request) {
	result := "123"
	Jsonify(w, result, 200)
}
