package routers

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"../store"
	"github.com/gorilla/sessions"
)

// Cookie - link on CookieStore, init on server
var Cookie *sessions.CookieStore
var keyContext interface{} = "profile"

func init() {
	gob.Register(&store.User{})
	Cookie = sessions.NewCookieStore(store.Config.Secret)
}

// Login - authentication
func Login(w http.ResponseWriter, r *http.Request) {
	ref := r.URL.Query().Get("ref")
	if ref == "" {
		ref = "/"
	}
	ses, err := Cookie.Get(r, "authentication-profile")
	if err != nil {
		log.Println("login:", err)
	}
	var au store.User
	JSONLoad(r, &au)
	verify := au.CheckPassword()
	if !verify {
		m := map[string]string{"status": "incorrect login or password"}
		Jsonify(w, m, 200)
		return
	}
	ses.Values["Profile"] = au
	_ = ses.Save(r, w)
	m := map[string]string{"status": "ok", "goto": ref}
	Jsonify(w, m, 200)
}

// LogOut - reset session cookie
func LogOut(w http.ResponseWriter, r *http.Request) {
	ses, _ := Cookie.Get(r, "authentication-profile")
	ses.Values["Profile"] = store.User{}
	ses.Options.MaxAge = -1
	Cookie.Save(r, w, ses)
	http.Redirect(w, r, "/login", 302)
}

// CookiesHandler - cookie validation handler
func CookiesHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ses, err := Cookie.Get(r, "authentication-profile")
		if err != nil {
			log.Println("middleware", err)
		}
		profile := ses.Values["Profile"]
		if profile == nil {
			ref := fmt.Sprintf("?ref=%s", r.URL.Path)
			http.Redirect(w, r, "/login"+ref, 302)
			return
		}
		ctx := context.WithValue(r.Context(), keyContext, profile)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
