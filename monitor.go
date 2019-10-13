package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type instagramJSON struct {
	Context     string `json:"@context"`
	Type        string `json:"@type"`
	Name        string `json:"name"`
	AltName     string `json:"alternateName"`
	Descripting string `json:"description"`
	URL         string `json:"url"`
	Image       string `json:"image"`
	PageData    struct {
		Type      string `json:"@type"`
		ID        string `json:"@id"`
		Statistic struct {
			Type     string `json:"@type"`
			SubType  string `json:"interactionType"`
			SubCount string `json:"userInteractionCount"`
		} `json:"interactionStatistic"`
	} `json:"mainEntityofPage"`
}

// followersCount - количество подписчиков
func (j *instagramJSON) followersCount() int {
	value, _ := strconv.ParseInt(j.PageData.Statistic.SubCount, 10, 64)
	return int(value)
}

func (j *instagramJSON) userID() string {
	return j.PageData.ID
}

// parseJSON - запись данных
func (j *instagramJSON) parseJSON(data []byte) {
	json.Unmarshal(data, &j)
}

func findJSON(r *http.Response) ([]byte, error) {
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	pattern := `<script type="application/ld.json">\n(.*)`
	rgx := regexp.MustCompile(pattern)
	js := rgx.FindSubmatch(responseData)
	if len(js) != 2 {
		return nil, errors.New("no math data")
	}
	return js[1], nil
}

func getURL(link chan string, results chan<- instagramJSON) {
	tr := &http.Transport{DisableKeepAlives: false}
	client := http.Client{Timeout: 600 * time.Second, Transport: tr}
	for {
		select {
		case url := <-link:
			var jsData instagramJSON
			req, err := http.NewRequest("GET", strings.TrimSpace(url), nil)
			if err != nil {
				log.Println(err)
				log.Fatalln(url)
			}
			req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.92 YaBrowser/19.10.0.1401 Yowser/2.5 Safari/537.36")
			response, err := client.Do(req)
			if err != nil || response.StatusCode != 200 {
				log.Println("Error code", response.StatusCode, url)
				if response.StatusCode == 429 {
					link <- url
					time.Sleep(20 * time.Second)
					continue
				}
			}

			js, err := findJSON(response)
			if err != nil {
				log.Println(err, url)
				results <- jsData
				time.Sleep(10 * time.Second)
			}

			jsData.parseJSON(js)
			// log.Println(jsData.followersCount())
			results <- jsData
			time.Sleep(10 * time.Second)
		default:
			continue
		}
	}

}

func readTXT() []string {
	b, _ := ioutil.ReadFile("links.txt")
	list := strings.Split(string(b), "\n")
	return list
}

func worker(urls []string, threads int) {
	links := make(chan string, len(urls))
	results := make(chan instagramJSON)

	for i := 0; i < threads; i++ {
		go getURL(links, results)
	}

	for _, url := range urls {
		links <- url
	}

	data := []instagramJSON{}
loop:
	for {
		select {
		case l := <-results:
			fmt.Println(l.followersCount(), l.userID())
			data = append(data, l)
			if len(data) == len(urls) {
				break loop
			}
		default:
			continue
		}
	}
	fmt.Println("\n", len(data))
	var result string
	fmt.Fscan(os.Stdin, &result)
}

func main() {
	urls := readTXT()
	worker(urls, 10)
}
