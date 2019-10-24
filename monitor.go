package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"./services"
	"./store"
)

func getURL(link chan store.Video, results chan<- store.Statistic) {
	tr := &http.Transport{DisableKeepAlives: false}
	client := &http.Client{Timeout: 600 * time.Second, Transport: tr}
	for {
		select {
		case url := <-link:
			req, err := http.NewRequest("GET", url.URL, nil)
			if err != nil {
				log.Fatalln(url, err)
			}
			req.Header.Set("user-agent", "Chrome/78.0.3904.70")
			response, _ := client.Do(req)
			b, _ := ioutil.ReadAll(response.Body)
			stat, err := services.ParseYoutube(b)
			stat.Video = url.ID
			if err != nil {
				log.Fatalln(err.Error())
			}
			results <- stat
		default:
			continue
		}
	}

}

func worker(urls []store.Video, threads int) []store.Statistic {
	links := make(chan store.Video, len(urls))
	results := make(chan store.Statistic)

	for i := 0; i < threads; i++ {
		go getURL(links, results)
	}

	for _, url := range urls {
		links <- url
	}

	var data []store.Statistic
loop:
	for {
		select {
		case l := <-results:
			data = append(data, l)
			if len(data) == len(urls) {
				break loop
			}
		default:
			continue
		}
	}
	return data
}

func main() {
	all := store.GetAllUrls()
	data := worker(all, 10)
	for _, s := range data {
		fmt.Println(s)
		s.Insert()
	}
	// var result string
	// fmt.Fscan(os.Stdin, &result)
}
