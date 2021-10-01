package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/kkdai/youtube/v2"
)

func main() {
	resp, err := soup.Get("https://gameranx.com/updates/id/261320/article/halo-5-guardians-is-still-not-in-development-for-pc-players/")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("iframe")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["src"])
		src := link.Attrs()["src"]
		if src == "" {
			src = link.Attrs()["data-src"]
		}
		if strings.Contains(src, "youtube.com") {

			iframeResp, err := soup.Get(src)
			if err != nil {
				os.Exit(1)
			}
			iframeDoc := soup.HTMLParse(iframeResp)

			youtubeLink := iframeDoc.Find("link", "rel", "canonical").Attrs()["href"]
			videoId := strings.Split(youtubeLink, "?v=")[1]
			client := youtube.Client{}

			video, err := client.GetVideo(videoId)
			if err != nil {
				panic(err)
			}

			stream, _, err := client.GetStream(video, &video.Formats[0])
			if err != nil {
				panic(err)
			}

			file, err := os.Create(video.Title + ".mp4")
			if err != nil {
				panic(err)
			}
			defer file.Close()

			_, err = io.Copy(file, stream)
			if err != nil {
				panic(err)
			}
			fmt.Println("Successfully saved video: ", video.Title)
			os.Exit(0)
		}
	}
	fmt.Println("No videos found")
}
