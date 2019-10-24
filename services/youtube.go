package services

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"../store"
)

// MetadataRenderer - ...
type MetadataRenderer struct {
	ID    string `json:"videoId"`
	Title struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
	} `json:"title"`
	Views struct {
		VCR struct {
			ViewCount struct {
				Count string `json:"simpleText"`
			} `json:"viewCount"`
		} `json:"videoViewCountRenderer"`
	} `json:"viewCount"`
	Likes struct {
		LBR struct {
			LikeCount    int64 `json:"likeCount"`
			DislikeCount int64 `json:"dislikeCount"`
		} `json:"likeButtonRenderer"`
	} `json:"likeButton"`
	DateText struct {
		SimpleText string `json:"simpleText"`
	} `json:"dateText"`
	Owner struct {
		VOR struct {
			SubscriberCountText struct {
				Runs []struct {
					Followers string `json:"text"`
				} `json:"runs"`
			} `json:"subscriberCountText"`
			ChanelName struct {
				Runs []struct {
					Text               string `json:"text"`
					NavigationEndpoint struct {
						BrowseEndpoint struct {
							ChannelID string `json:"browseId"`
						} `json:"browseEndpoint"`
					} `json:"navigationEndpoint"`
				} `json:"runs"`
			} `json:"title"`
		} `json:"videoOwnerRenderer"`
	} `json:"owner"`
}

// YouTube - ...
type YouTube struct {
	// AutoPlay int64 `json:"autoplay_count"`
	// RVC      string `json:"rvs"`
	RawWNR string `json:"watch_next_response"`
	WNR    struct {
		RContext struct {
			TCWR struct {
				Res1 struct {
					Res2 struct {
						Contents []struct {
							SectionRenderer struct {
								Contents []struct {
									MetadataRenderer *MetadataRenderer `json:"videoMetadataRenderer"`
								} `json:"contents"`
							} `json:"itemSectionRenderer"`
						} `json:"contents"`
					} `json:"results"`
				} `json:"results"`
			} `json:"twoColumnWatchNextResults"`
		} `json:"contents"`
	}
}

// ParseYoutube - ...
func ParseYoutube(html []byte) (store.Statistic, error) {
	pattern := regexp.MustCompile(`'RELATED_PLAYER_ARGS': (.*),\n`)
	data := pattern.FindSubmatch(html)
	replacer := strings.NewReplacer(`//`, ``, `""`, `"`)
	cv := new(store.Statistic)
	if len(data) > 1 {
		youtube := new(YouTube)
		js := replacer.Replace(string(data[1]))
		// ioutil.WriteFile("___raw.json", []byte(js), 0666)
		json.Unmarshal([]byte(js), &youtube)
		json.Unmarshal([]byte(youtube.RawWNR), &youtube.WNR)
		youtube.RawWNR = ""

		// bjson, _ := json.Marshal(youtube)
		// ioutil.WriteFile("___parse.json", bjson, 0666)

		clearPatternVideo := regexp.MustCompile(`[^\d]+`)
		clearViews := clearPatternVideo.ReplaceAll([]byte(youtube.WNR.RContext.TCWR.Res1.Res2.Contents[0].SectionRenderer.Contents[0].MetadataRenderer.Views.VCR.ViewCount.Count), []byte(""))
		views, _ := strconv.ParseInt(string(clearViews), 10, 64)

		cv.ID = youtube.WNR.RContext.TCWR.Res1.Res2.Contents[0].SectionRenderer.Contents[0].MetadataRenderer.ID
		cv.Views = views
		cv.Likes = youtube.WNR.RContext.TCWR.Res1.Res2.Contents[0].SectionRenderer.Contents[0].MetadataRenderer.Likes.LBR.LikeCount
		cv.Dislikes = youtube.WNR.RContext.TCWR.Res1.Res2.Contents[0].SectionRenderer.Contents[0].MetadataRenderer.Likes.LBR.DislikeCount
		cv.Title = youtube.WNR.RContext.TCWR.Res1.Res2.Contents[0].SectionRenderer.Contents[0].MetadataRenderer.Title.Runs[0].Text
		cv.ChannelName = youtube.WNR.RContext.TCWR.Res1.Res2.Contents[0].SectionRenderer.Contents[0].MetadataRenderer.Owner.VOR.ChanelName.Runs[0].Text
		cv.ChannelID = youtube.WNR.RContext.TCWR.Res1.Res2.Contents[0].SectionRenderer.Contents[0].MetadataRenderer.Owner.VOR.ChanelName.Runs[0].NavigationEndpoint.BrowseEndpoint.ChannelID
		cv.Followers = youtube.WNR.RContext.TCWR.Res1.Res2.Contents[0].SectionRenderer.Contents[0].MetadataRenderer.Owner.VOR.SubscriberCountText.Runs[0].Followers
		cv.UploadDate = youtube.WNR.RContext.TCWR.Res1.Res2.Contents[0].SectionRenderer.Contents[0].MetadataRenderer.DateText.SimpleText
		return *cv, nil
	}
	return *cv, errors.New("can't parse")
}

// func main() {
// 	url := "https://www.youtube.com/watch?v=nH2qi4FoJ7M"
// 	client := &http.Client{}
// 	r, err := client.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	youtube, err := ParseYoutube(b)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%+v", youtube)
// }
