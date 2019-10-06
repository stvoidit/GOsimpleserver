package routers

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

type User struct {
	Role          string
	UserID        int
	Department    int
	Authenticated bool
}

func (u User) checkRole(roles []string, w http.ResponseWriter, r *http.Request) {
	var inArray bool
	for _, val := range roles {
		if u.Role == val {
			inArray = true
			break
		} else {
			inArray = false
		}
	}
	if !inArray {
		ref := fmt.Sprintf("?ref=%s", r.URL.Path)
		http.Redirect(w, r, "/login"+ref, http.StatusFound)
	}
}

const secretKey = "MySecretKey"

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte(secretKey))
	store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}
	gob.Register(User{})
}

func getUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}
	user, ok := val.(User)
	if !ok {
		user = User{}
	}
	return user
}

// Login - типа авторизация
func Login(w http.ResponseWriter, r *http.Request) {
	ref := r.URL.Query().Get("ref")
	fmt.Println(ref)
	ses, _ := store.Get(r, "user")
	user := &User{
		Role:          "admin",
		UserID:        2,
		Department:    1,
		Authenticated: true,
	}
	ses.Values["user"] = user
	_ = ses.Save(r, w)
	http.Redirect(w, r, ref, http.StatusFound)
}

// LogOut - сброс сессии
func LogOut(w http.ResponseWriter, r *http.Request) {
	ses, _ := store.Get(r, "user")
	ses.Values["user"] = User{}
	ses.Options.MaxAge = -1
	store.Save(r, w, ses)
	w.Write([]byte("You are logout"))
}
