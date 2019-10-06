package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"./routers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routers.IndexRoute)
	r.HandleFunc("/test", routers.IndexRoute)
	r.HandleFunc("/login", routers.Login)
	r.HandleFunc("/logout", routers.LogOut)
	http.ListenAndServe("0.0.0.0:9000", r)
}
