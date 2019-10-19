package routers

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretKey = "MySecretKey"

// TokenClaim - token custom type
type TokenClaim struct {
	user
	jwt.StandardClaims
}

// GetTokenHandler - get api token
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	claim := TokenClaim{user{"admin", 1, 1, true}, jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(15 * time.Minute)).Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	w.Write([]byte(tokenString))
}

// TokenHandler - api validation middleware
func TokenHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := strings.ContainsAny(r.Header.Get("Authorization"), "Bearer")
		if bearer {
			bearer := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
			profile, err := checkToken(bearer)
			if err != nil {
				w.WriteHeader(403)
				w.Write([]byte(err.Error()))
				return
			}
			ctx := context.WithValue(r.Context(), keyContext, profile)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(403)
			w.Write([]byte("Not validate token!"))
			return
		}
	})
}

// checkToken - token validation func
func checkToken(tokenString string) (*user, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad")
		}
		return []byte(secretKey), nil
	})
	if claim, ok := token.Claims.(*TokenClaim); ok && token.Valid {
		return &claim.user, nil
	}
	return &user{}, err
}
