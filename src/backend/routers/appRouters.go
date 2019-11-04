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
	q := r.URL.Query().Get("chanId")
	if q == "" {
		Jsonify(w, "Pleace choise channel ID", 200)
		return
	}
	result, err := store.Redis.GetCache(q)
	if err != nil {
		result := store.GetStat(q)
		store.Redis.SetJSON(q, result)
		Jsonify(w, result, 200)
		return
	}
	Jsonify(w, string(result), 200)
}

// UserChannels - ..
func UserChannels(w http.ResponseWriter, r *http.Request) {
	result, err := store.GetAllChanels()
	if err != nil {
		Jsonify(w, err.Error(), 509)
		return
	}
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
