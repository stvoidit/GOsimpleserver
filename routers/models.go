package routers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Jsonify - отправка json response
func Jsonify(w http.ResponseWriter, i interface{}, code int) {
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	encoder.Encode(i)
}

// JSONLoad - преобразование json в структуру
func JSONLoad(req *http.Request, i interface{}) (interface{}, error) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// User - cookies объект
type User struct {
	Role          string
	UserID        int
	Department    int
	Authenticated bool
}

func (u User) checkRole(roles []string) error {
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
		return errors.New("not validate")
	}
	return nil
}

// SKUD - something data
type SKUD struct {
	ID           int
	VocationUser int
	Date         time.Time
}
