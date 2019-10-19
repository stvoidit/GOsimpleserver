package routers

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// Cookie - link on CookieStore, init on server
var Cookie *sessions.CookieStore

func init() {
	gob.Register(&user{})
}

// User - cookies custom type
type user struct {
	Role          string
	UserID        int
	Department    int
	Authenticated bool
}

// Login - authentication
func Login(w http.ResponseWriter, r *http.Request) {
	// ref := r.URL.Query().Get("ref")
	ses, err := Cookie.Get(r, "authentication-profile")
	if err != nil {
		panic(err)
	}

	// TODO: check user in database
	profile := &user{
		Role:          "admin",
		UserID:        2,
		Department:    1,
		Authenticated: true,
	}
	ses.Values["Profile"] = profile
	_ = ses.Save(r, w)
	// 	http.Redirect(w, r, ref, http.StatusFound)
	// 	return
	// }
	// if r.Method == "GET" {
	// 	res := map[string]string{
	// 		"message": "Its login screen!",
	// 	}
	// 	Jsonify(w, res, 200)
	// }
	w.Write([]byte("login"))
}

// LogOut - reset session cookie
func LogOut(w http.ResponseWriter, r *http.Request) {
	ses, _ := Cookie.Get(r, "authentication-profile")
	ses.Values["Profile"] = user{}
	ses.Options.MaxAge = -1
	Cookie.Save(r, w, ses)
	w.Write([]byte("You are logout"))
}

// CookiesHandler - cookie validation handler
func CookiesHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ses, err := Cookie.Get(r, "authentication-profile")
		if err != nil {
			panic(err)
		}
		fmt.Println(ses.Values["Profile"])
		profile := ses.Values["Profile"]
		if profile == nil {
			w.WriteHeader(401)
			w.Write([]byte("not auth"))
			return
		}
		// w.Write([]byte("middleware\n"))
		// user := ses.GetUserData(r)
		// fmt.Println(user)
		// err := user.checkRole(filterRoles)
		// if err != nil {
		// 	ref := fmt.Sprintf("?ref=%s", r.URL.Path)
		// 	http.Redirect(w, r, "/login"+ref, 302)
		// 	return
		// }
		next.ServeHTTP(w, r)
	})
}
