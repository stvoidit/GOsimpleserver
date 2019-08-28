package main

import (
	"net/http"

	"./routers"
)

func main() {
	http.HandleFunc("/", routers.IndexRoute)
	http.HandleFunc("/user", routers.AjaxUsers)
	http.HandleFunc("/dep", routers.AjaxDepartment)
	http.HandleFunc("/db", routers.AjaxDB)
	http.HandleFunc("/getone", routers.AjaxGetOne)
	http.ListenAndServe("0.0.0.0:9000", nil)
}
