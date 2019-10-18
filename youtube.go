// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"regexp"
// 	"strings"
// )

// type MetadataRenderer struct {
// 	ID    string `json:"videoId"`
// 	Title struct {
// 		Runs []struct {
// 			Text string `json:"text"`
// 		} `json:"runs"`
// 	} `json:"title"`
// 	Views struct {
// 		VCR struct {
// 			ViewCount struct {
// 				Count string `json:"simpleText"`
// 			} `json:"viewCount"`
// 		} `json:"videoViewCountRenderer"`
// 	} `json:"viewCount"`
// 	Likes struct {
// 		LBR struct {
// 			LikeCount    int `json:"likeCount"`
// 			DislikeCount int `json:"dislikeCount"`
// 		} `json:"likeButtonRenderer"`
// 	} `json:"likeButton"`
// }

// type YouTube struct {
// 	AutoPlay int    `json:"autoplay_count"`
// 	RVC      string `json:"rvs"`
// 	RawWNR   string `json:"watch_next_response"`
// 	WNR      struct {
// 		RContext struct {
// 			TCWR struct {
// 				Res1 struct {
// 					Res2 struct {
// 						Contents []struct {
// 							SectionRenderer struct {
// 								Contents []struct {
// 									MetadataRenderer *MetadataRenderer `json:"videoMetadataRenderer"`
// 								} `json:"contents"`
// 							} `json:"itemSectionRenderer"`
// 						} `json:"contents"`
// 					} `json:"results"`
// 				} `json:"results"`
// 			} `json:"twoColumnWatchNextResults"`
// 		} `json:"contents"`
// 	}
// }

// func main() {
// 	url := "..."
// 	client := &http.Client{}
// 	r, err := client.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	pattern := regexp.MustCompile(`'RELATED_PLAYER_ARGS': (.*),\n`)
// 	data := pattern.FindSubmatch(b)
// 	replacer := strings.NewReplacer(`//`, ``, `""`, `"`)
// 	if len(data) > 1 {
// 		js := replacer.Replace(string(data[1]))
// 		youtube := new(YouTube)
// 		json.Unmarshal([]byte(js), &youtube)
// 		json.Unmarshal([]byte(youtube.RawWNR), &youtube.WNR)
// 		fmt.Println(len(youtube.WNR.RContext.TCWR.Res1.Res2.Contents))
// 		for _, v := range youtube.WNR.RContext.TCWR.Res1.Res2.Contents {
// 			for _, i := range v.SectionRenderer.Contents {
// 				log.Println(i.MetadataRenderer)
// 			}
// 			fmt.Println()
// 		}
// 	}
// }
