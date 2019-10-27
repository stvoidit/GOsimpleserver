package application

import (
	"net/http"
	"path"

	"../routers"
	"github.com/gorilla/mux"
)

const staticPath = "static"

// App - ...
var App NewApp

func init() {
	App.GetConfig()
	routers.Cookie = App.Session
	App.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))
	App.Router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(path.Join(staticPath, "js")))))
	App.Router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(path.Join(staticPath, "css")))))
	App.routers()
}

func (app *NewApp) routers() {
	routers.RegistrateTemplates(path.Join(staticPath, "templates"))

	public := app.Router
	public.HandleFunc("/login", routers.Login)
	public.HandleFunc("/logout", routers.LogOut)
	public.HandleFunc("/get-token", routers.GetTokenHandler)
	public.HandleFunc("/UserVideos", routers.UserVideos).Methods("GET")
	public.HandleFunc("/MyVieos", routers.MyVieos).Methods("GET")

	api := app.apiRouter()
	api.HandleFunc("/AddVideo", routers.AddVideo).Methods("POST")

	private := app.cookieRouter()
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
