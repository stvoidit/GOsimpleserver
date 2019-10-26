package application

import (
	"../routers"
	"github.com/gorilla/mux"
)

// App - ...
var App NewApp

func init() {
	App.GetConfig()
	routers.Cookie = App.Session
	App.routers()
}

func (app *NewApp) routers() {
	public := app.Router
	public.HandleFunc("/login", routers.Login)
	public.HandleFunc("/logout", routers.LogOut)
	public.HandleFunc("/get-token", routers.GetTokenHandler)
	public.HandleFunc("/AddVideo", routers.AddVideo).Methods("POST")

	api := app.apiRouter()
	api.HandleFunc("/departments", routers.Departments)
	api.HandleFunc("/usersdeps", routers.UsersDepartments)
	api.HandleFunc("/users", routers.Users)
	api.HandleFunc("/AddVideo", routers.AddVideo).Methods("POST")

	private := app.cookieRouter()
	private.HandleFunc("/departments", routers.Departments)
	private.HandleFunc("/usersdeps", routers.UsersDepartments)
	private.HandleFunc("/users", routers.Users)
	private.HandleFunc("/AddVideo", routers.AddVideo).Methods("POST")
}

func (app *NewApp) cookieRouter() *mux.Router {
	private := app.Router.NewRoute().Subrouter()
	private.Use(routers.CookiesHandler)
	return private
}

func (app *NewApp) apiRouter() *mux.Router {
	api := app.Router.NewRoute().PathPrefix("/api/").Subrouter()
	api.Use(routers.TokenHandler)
	return api
}
