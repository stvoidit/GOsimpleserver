package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"./routers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routers.IndexRoute)
	r.HandleFunc("/user", routers.AjaxUsers)
	r.HandleFunc("/dep", routers.AjaxDepartment)
	r.HandleFunc("/login", routers.Login)
	// http.HandleFunc("/db", routers.AjaxDB)
	// http.HandleFunc("/getone", routers.AjaxGetOne)
	http.ListenAndServe("0.0.0.0:9000", r)
}
