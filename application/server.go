package application

import (
	"../routers"
)

// App - ...
var App NewApp

func init() {
	App.GetConfig()
	App.middleware()
	App.routers()
}

func (app *NewApp) routers() {
	app.Router.HandleFunc("/", routers.IndexRoute)
	app.Router.HandleFunc("/login", routers.Login)
}

func (app *NewApp) middleware() {
	app.Router.Use(routers.CookiesHandler)
}
