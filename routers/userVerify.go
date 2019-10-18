package routers

import (
	"net/http"
)

// User - cookies объект
type user struct {
	Role          string
	UserID        int
	Department    int
	Authenticated bool
}

// Login - авторизация
func Login(w http.ResponseWriter, r *http.Request) {
	// ref := r.URL.Query().Get("ref")
	// ses, _ := store.Get(r, "user")
	// if r.Method == "POST" {
	// 	var Data struct {
	// 		Username string `json:"username"`
	// 		Password string `json:"password"`
	// 	}
	// 	_, err := JSONLoad(r, &Data)
	// 	if err != nil {
	// 		res := map[string]string{
	// 			"message": "Wrong login or password!",
	// 		}
	// 		Jsonify(w, res, 401)
	// 		return
	// 	}
	// 	user := &User{
	// 		Role:          "admin",
	// 		UserID:        2,
	// 		Department:    1,
	// 		Authenticated: true,
	// 	}
	// 	ses.Values["user"] = user
	// 	_ = ses.Save(r, w)
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

// LogOut - сброс сессии
// func LogOut(w http.ResponseWriter, r *http.Request) {
// 	ses, _ := store.Get(r, "user")
// 	ses.Values["user"] = User{}
// 	ses.Options.MaxAge = -1
// 	store.Save(r, w, ses)
// 	w.Write([]byte("You are logout"))
// }

// GetTokenHandler - получение токена
// func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"Role":          "admin",
// 		"UserID":        1,
// 		"Department":    1,
// 		"Authenticated": true,
// 	})

// 	tokenString, _ := token.SignedString(secretKey)
// 	w.Write([]byte(tokenString))
// }

// checkToken - валидация
// func checkToken(tokenString string) (User, error) {
// 	u := User{}
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("bad")
// 		}
// 		return secretKey, nil
// 	})
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		u.Role = claims["Role"].(string)
// 		u.UserID = int(claims["UserID"].(float64))
// 		u.Department = int(claims["Department"].(float64))
// 		u.Authenticated = claims["Authenticated"].(bool)
// 		return u, nil
// 	}
// 	return u, err
// }

// ValidateToken - валидация токена api
// func ValidateToken(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		bearer := strings.ContainsAny(r.Header.Get("Authorization"), "Bearer")
// 		if bearer {
// 			bearer := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
// 			_, err := checkToken(bearer)
// 			if err != nil {
// 				w.Write([]byte(err.Error()))
// 				return
// 			}
// 		} else {
// 			w.WriteHeader(403)
// 			w.Write([]byte("Not validate token!"))
// 			return
// 		}
// 		h.ServeHTTP(w, r)
// 	})
// }

// TokenHandler - декоратор для api
// func TokenHandler(h http.Handler, adapters ...func(http.Handler) http.Handler) http.Handler {
// 	for _, adapter := range adapters {
// 		h = adapter(h)
// 	}
// 	return h
// }

// CookiesHandler - Валидация по кукам
func CookiesHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
