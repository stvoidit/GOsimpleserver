package routers

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
)

var secretKey = []byte("MySecretKey")

var store *sessions.CookieStore

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

func init() {
	store = sessions.NewCookieStore(secretKey)
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

// Login - авторизация
func Login(w http.ResponseWriter, r *http.Request) {
	ref := r.URL.Query().Get("ref")
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

// GetTokenHandler - получение токена
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Role":          "admin",
		"UserID":        1,
		"Department":    1,
		"Authenticated": true,
	})

	tokenString, _ := token.SignedString(secretKey)
	w.Write([]byte(tokenString))
}

// APIValidate - валидация токена
func APIValidate(w http.ResponseWriter, r *http.Request) {
	bearer := strings.ContainsAny(r.Header.Get("Authorization"), "Bearer")
	if bearer {
		bearer := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
		_, err := ValidateToken(bearer)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	}
}

// ValidateToken - валидация и возврат типа Юзер
func ValidateToken(tokenString string) (User, error) {
	u := User{}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad")
		}
		return secretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		u.Role = claims["Role"].(string)
		u.UserID = int(claims["UserID"].(float64))
		u.Department = int(claims["Department"].(float64))
		u.Authenticated = claims["Authenticated"].(bool)
		return u, nil
	}
	return u, err

}
