package routers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretKey = "MySecretKey"

// TokenClaim - token custom type
type TokenClaim struct {
	Role          string
	UserID        int64
	DeparmentID   int64
	Authenticated bool
	jwt.StandardClaims
}

// GetTokenHandler - get api token
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	claim := TokenClaim{"admin", 1, 1, true, jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(15 * time.Minute)).Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	fmt.Println(secretKey)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	w.Write([]byte(tokenString))
}

// TokenHandler - api validation middleware
func TokenHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := strings.ContainsAny(r.Header.Get("Authorization"), "Bearer")
		if bearer {
			bearer := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
			_, err := checkToken(bearer)
			if err != nil {
				w.WriteHeader(403)
				w.Write([]byte(err.Error()))
				return
			}
		} else {
			w.WriteHeader(403)
			w.Write([]byte("Not validate token!"))
			return
		}
		h.ServeHTTP(w, r)
	})
}

// checkToken - token validation func
func checkToken(tokenString string) (TokenClaim, error) {
	u := TokenClaim{}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return u, err
	}
	check := token.Claims.Valid()
	if check != nil {
		return u, check
	}
	return u, err
}
