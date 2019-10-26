package routers

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"../store"

	"github.com/dgrijalva/jwt-go"
)

const secretKey = "MySecretKey"

// TokenClaim - token custom type
type TokenClaim struct {
	store.User
	jwt.StandardClaims
}

// GetTokenHandler - get api token
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	if username, password, ok := r.BasicAuth(); ok {
		anon := store.User{Username: username, Password: password}
		if !anon.CheckPassword() {
			message := map[string]string{"status": "incorrect passssword or login"}
			Jsonify(w, message, 401)
			return
		}
		claim := TokenClaim{anon, jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Duration(15 * time.Minute)).Unix(),
		}}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			panic(err)
		}
		w.Write([]byte(tokenString))
		return
	}
	Jsonify(w, map[string]string{"status": "need authorization Basic data"}, 401)
	return
}

// TokenHandler - api validation middleware
func TokenHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := strings.Contains(r.Header.Get("Authorization"), "Bearer")
		if bearer {
			token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
			profile, err := checkToken(token)
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			ctx := context.WithValue(r.Context(), keyContext, profile)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(401)
			w.Write([]byte("Need Bearer token in Authorization headers"))
			return
		}
	})
}

// checkToken - token validation func
func checkToken(tokenString string) (*store.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return &store.User{}, err
	}
	if claim, ok := token.Claims.(*TokenClaim); ok && token.Valid {
		return &claim.User, nil
	}
	return &store.User{}, err
}
