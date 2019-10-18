package application

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// NewApp - ...
type NewApp struct {
	Router  *mux.Router
	Session *sessions.CookieStore
}

// GetConfig - ...
func (a *NewApp) GetConfig() {
	ses := sessions.NewCookieStore([]byte("MySecretKey"))
	a.Router = mux.NewRouter()
	a.Session = ses
}
