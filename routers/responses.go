package routers

import (
	"encoding/json"
	"net/http"
)

// Jsonify - отправка json response
func Jsonify(w http.ResponseWriter, i interface{}, code int) {
	switch i.(type) {
	default:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	case string:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	}
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
