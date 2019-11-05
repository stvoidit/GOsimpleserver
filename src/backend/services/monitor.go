package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"../store"
	services "./youtube"
)

var logger log.Logger

func init() {
	logfile, err := os.OpenFile("monitor.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(logfile)
}

func getURL(link chan store.Video, results chan<- store.Statistic, wg *sync.WaitGroup) {
	defer wg.Done()
	tr := &http.Transport{DisableKeepAlives: false}
	client := &http.Client{Timeout: 600 * time.Second, Transport: tr}
	for url := range link {
		req, err := http.NewRequest("GET", url.URL, nil)
		if err != nil {
			logger.Println(url.URL, err.Error())
		}
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 YaBrowser/19.10.2.195 Yowser/2.5 Safari/537.36")
		response, err := client.Do(req)
		if err != nil {
			continue
		}
		b, err := ioutil.ReadAll(response.Body)
		if strings.Contains(string(b), "action_recaptcha_verify2") {
			logger.Println("recapcha detected!")
			continue
		}
		if err != nil {
			continue
		}
		defer response.Body.Close()
		stat, err := services.ParseYoutube(b, url.ID)
		if err != nil {
			logger.Println(url.URL, err.Error())
		}
		results <- stat
	}
}

func worker(urls []store.Video, threads int) []store.Statistic {
	links := make(chan store.Video, len(urls))
	results := make(chan store.Statistic, len(urls))

	for _, url := range urls {
		links <- url
	}
	close(links)

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go getURL(links, results, &wg)
	}

	wg.Wait()
	close(results)

	var data []store.Statistic
	for {
		l, ok := <-results
		if !ok {
			break
		}
		data = append(data, l)
	}
	return data
}

func main() {
	all := store.GetAllUrls()
	data := worker(all, 4)
	for _, s := range data {
		s.Insert()
	}
	logger.Println("done")
}
