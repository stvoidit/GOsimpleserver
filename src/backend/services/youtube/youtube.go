package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"../../store"
)

// Video - ...
type Video struct {
	ID          string
	PrimaryInfo struct {
		Title struct {
			Runs []struct {
				Text string `json:"text"`
			} `json:"runs"`
		} `json:"title"`
		Views struct {
			Views struct {
				Count struct {
					Text string `json:"simpleText"`
				} `json:"viewCount"`
			} `json:"videoViewCountRenderer"`
		} `json:"viewCount"`
		LikesStatus struct {
			Renderer struct {
				Stat string `json:"tooltip"`
			} `json:"sentimentBarRenderer"`
		} `json:"sentimentBar"`
	} `json:"videoPrimaryInfoRenderer"`
}

// Channel - ...
type Channel struct {
	SecondaryInfo struct {
		Owner struct {
			OwnerRenderer struct {
				Title struct {
					Runs []struct {
						Text     string `json:"text"`
						Endpoint struct {
							Endpoint struct {
								ChannelID string `json:"browseId"`
							} `json:"browseEndpoint"`
						} `json:"navigationEndpoint"`
					} `json:"runs"`
				} `json:"title"`
				Subscribers struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
				} `json:"subscriberCountText"`
			} `json:"videoOwnerRenderer"`
		} `json:"owner"`
	} `json:"videoSecondaryInfoRenderer"`
}

// YouTube - ...
type YouTube struct {
	Contents struct {
		CWNR struct {
			Results struct {
				Results struct {
					Contents []struct {
						Video
						Channel
					} `json:"contents"`
				} `json:"results"`
			} `json:"results"`
		} `json:"twoColumnWatchNextResults"`
	} `json:"contents"`
}

// ParseYoutube - ...
func ParseYoutube(html []byte, ID string) (store.Statistic, error) {
	patternJSON := regexp.MustCompile(`window..ytInitialData.. = (.*);\n`)
	patternViews := regexp.MustCompile(`[^\d]+`)
	data := patternJSON.FindSubmatch(html)
	// ioutil.WriteFile("___peace.json", data[1], 0666)
	var cv store.Statistic
	cv.Video = ID
	cv.ID = ID
	if len(data) > 1 {
		// ioutil.WriteFile("__js.json", data[1], 666)
		var youtubeRenderData YouTube
		json.Unmarshal(data[1], &youtubeRenderData)
		// fmt.Println(youtubeRenderData.Contents.CWNR.Results.Results.Contents[0])

		if len(youtubeRenderData.Contents.CWNR.Results.Results.Contents[0].PrimaryInfo.Title.Runs) != 0 {
			cv.Title = youtubeRenderData.Contents.CWNR.Results.Results.Contents[0].PrimaryInfo.Title.Runs[0].Text
			// ioutil.WriteFile("___errLen.json", data[1], 0666)
			// defer os.Exit(1)
		} else {
			cv.Title = ""
		}

		if len(youtubeRenderData.Contents.CWNR.Results.Results.Contents[1].SecondaryInfo.Owner.OwnerRenderer.Title.Runs) != 0 {
			cv.ChannelName = youtubeRenderData.Contents.CWNR.Results.Results.Contents[1].SecondaryInfo.Owner.OwnerRenderer.Title.Runs[0].Text
			cv.ChannelID = youtubeRenderData.Contents.CWNR.Results.Results.Contents[1].SecondaryInfo.Owner.OwnerRenderer.Title.Runs[0].Endpoint.Endpoint.ChannelID
		} else {
			cv.ChannelName = ""
			cv.ChannelID = ""
			// ioutil.WriteFile("__js.json", data[1], 666)
			// ioutil.WriteFile("___errLen.json", data[1], 0666)
			// defer os.Exit(1)
		}

		if len(youtubeRenderData.Contents.CWNR.Results.Results.Contents[1].SecondaryInfo.Owner.OwnerRenderer.Subscribers.Runs) != 0 {
			cv.Followers = youtubeRenderData.Contents.CWNR.Results.Results.Contents[1].SecondaryInfo.Owner.OwnerRenderer.Subscribers.Runs[0].Text
		} else {
			cv.Followers = ""
			// ioutil.WriteFile("___errLen.json", data[1], 0666)
			// defer os.Exit(1)
		}

		// parse views in int64
		cv.Views = func() int64 {
			views := youtubeRenderData.Contents.CWNR.Results.Results.Contents[0].PrimaryInfo.Views.Views.Count.Text
			clearViews := patternViews.ReplaceAllString(views, "")
			intViews, err := strconv.ParseInt(clearViews, 10, 64)
			if err != nil {
				intViews = 0
			}
			return intViews
		}()

		// parse likes and dislkes
		cv.Likes, cv.Dislikes = func() (int64, int64) {
			stat := youtubeRenderData.Contents.CWNR.Results.Results.Contents[0].PrimaryInfo.LikesStatus.Renderer.Stat
			splitStat := strings.Split(stat, "/")
			if len(splitStat) != 2 {
				return 0, 0
			}
			likes := patternViews.ReplaceAllString(splitStat[0], "")
			dislikes := patternViews.ReplaceAllString(splitStat[1], "")

			intLikes, err := strconv.ParseInt(likes, 10, 64)
			if err != nil {
				intLikes = 0
			}
			intDislikes, err := strconv.ParseInt(dislikes, 10, 64)
			if err != nil {
				intDislikes = 0
			}
			return intLikes, intDislikes
		}()

		return cv, nil
	}
	// ioutil.WriteFile("___html.html", html, 0666)
	// defer os.Exit(1)
	return cv, errors.New("can't parse")
}

// AddNew - ...
func AddNew(url string, client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 YaBrowser/19.10.2.195 Yowser/2.5 Safari/537.36")
	response, _ := client.Do(req)
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	response.Body.Close()
	videoID := strings.Split(url, "v=")[1]
	s, err := ParseYoutube(b, videoID)
	if err != nil {
		fmt.Println(url, err.Error())
	}
	exist := s.InsertVideo(url)
	if exist {
		s.Insert()
	}
}
