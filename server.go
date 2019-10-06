package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"./routers"
)

func main() {
	r := mux.NewRouter()
	var cH = routers.CookiesHandler  // cookies Handler decorator
	var cV = routers.ValidateCookies // cookies Validating func

	r.HandleFunc("/test", routers.IndexRoute)
	r.PathPrefix("/api/").Handler(http.StripPrefix("/api", routers.TokenHandler(r, routers.ValidateToken)))

	r.HandleFunc("/get-token", routers.GetTokenHandler)
	r.Handle("/", cH(routers.IndexRoute, cV, []string{"moderator", "admin"}))
	r.HandleFunc("/login", routers.Login)
	r.HandleFunc("/logout", routers.LogOut)

	// http.ListenAndServe("0.0.0.0:9000", r)
	http.ListenAndServe("0.0.0.0:9000", handlers.LoggingHandler(os.Stdout, r))
}
