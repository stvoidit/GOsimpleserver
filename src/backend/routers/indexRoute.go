package routers

import (
	"net/http"
	"sync"
	"time"

	"../services"
	"../store"
)

var db = store.DB

// Users = ...
func Users(w http.ResponseWriter, r *http.Request) {
	result := "123"
	Jsonify(w, result, 200)

}

// Departments - ...
func Departments(w http.ResponseWriter, r *http.Request) {
	result := "123"
	Jsonify(w, result, 200)
}

// UsersDepartments - ...
func UsersDepartments(w http.ResponseWriter, r *http.Request) {
	result := "123"
	Jsonify(w, result, 200)
}

// AddVideo - ...
func AddVideo(w http.ResponseWriter, r *http.Request) {
	var NewVideos []struct {
		URL string `json:"url"`
	}
	JSONLoad(r, &NewVideos)
	tr := &http.Transport{DisableKeepAlives: false}
	client := &http.Client{Timeout: 600 * time.Second, Transport: tr}
	waiting := sync.WaitGroup{}
	for _, link := range NewVideos {
		go services.AddNew(link.URL, client, &waiting)
	}
	waiting.Wait()
	response := map[string]string{"status": "add videos"}
	Jsonify(w, response, 201)
}
