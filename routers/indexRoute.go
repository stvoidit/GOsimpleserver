package routers

import (
	"net/http"

	"../store"
)

var db = store.DB

// IndexRoute = is '/'
// Передача параметров в template
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	result := db.AllUsers()
	Jsonify(w, result, 200)

}

// Something - ...
func Something(w http.ResponseWriter, r *http.Request) {
	Jsonify(w, nil, 200)
}
