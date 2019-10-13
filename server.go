package main

import (
	"net/http"
	"os"

	"./routers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	var validateFunc = routers.CookiesHandler // cookies Handler decorator
	var chekingFunc = routers.ValidateCookies // cookies Validating func

	r.HandleFunc("/test", routers.IndexRoute)
	r.PathPrefix("/api/").Handler(http.StripPrefix("/api", routers.TokenHandler(r, routers.ValidateToken)))

	r.HandleFunc("/get-token", routers.GetTokenHandler)
	r.Handle("/", validateFunc(routers.IndexRoute, chekingFunc, []string{"moderator", "admin"}))
	r.HandleFunc("/something", routers.Something)
	r.HandleFunc("/login", routers.Login)
	r.HandleFunc("/logout", routers.LogOut)

	logfile, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	defer logfile.Close()
	http.ListenAndServe("0.0.0.0:9000", handlers.LoggingHandler(logfile, r))
}
