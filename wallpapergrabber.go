package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

var userAgentString = "go:wallpapergrabber:v0.0.1 (by /u/DishFromZanzibar)"
var redditWallpaperEndpoint = "https://old.reddit.com/r/wallpapers/"

func main() {
	soup.Header("User-Agent", userAgentString)

	resp, err := soup.Get(redditWallpaperEndpoint)
	if err != nil {
		fmt.Println("Error making request")
		return
	}

	doc := soup.HTMLParse(resp)
	wallpaperPath := doc.Find("div", "class", "thing").Attrs()["data-url"]

	DownloadFile(wallpaperPath)
}

// DownloadFile downloads a file to the disk
// Kudos to https://golangcode.com/download-a-file-from-a-url/
func DownloadFile(url string) error {
	splitURL := strings.Split(url, "/")
	filepath := splitURL[len(splitURL)-1]

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
