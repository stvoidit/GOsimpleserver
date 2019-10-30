package routers

import (
	"net/http"
	"sync"
	"time"

	services "../services/youtube"
	"../store"
)

// UserVideos - ...
func UserVideos(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	result := store.GetStat(q.Get("chanId"))
	Jsonify(w, result, 200)
}

// UserChannels - ..
func UserChannels(w http.ResponseWriter, r *http.Request) {
	result := store.GetAllChanels()
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
