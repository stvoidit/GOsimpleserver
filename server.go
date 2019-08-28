package main

import (
	"net/http"

	"./routers"
)

func main() {
	http.HandleFunc("/", routers.IndexRoute)
	http.HandleFunc("/user", routers.AjaxUsers)
	http.HandleFunc("/dep", routers.AjaxDepartment)
	http.ListenAndServe("0.0.0.0:9000", nil)
}
