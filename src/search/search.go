package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/salawhaaat/aws-go-yt-search/src/config"
)

func Search(query string) (string, error) {
	url := fmt.Sprintf("%s?part=id&type=video&q=%s&key=%s", config.Cfg.SearchAPI, strings.ReplaceAll(query, " ", "+"), config.Cfg.ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response YouTubeSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}
	if len(response.Items) == 0 {
		return "", fmt.Errorf("No video found for the query: %s", query)
	}
	link := fmt.Sprintf("https://www.youtube.com/watch?v=%s", response.Items[0].ID.VideoID)
	return link, nil
}

func ProcessLinks(links []string) map[string]string {
	results := make(map[string]string)
	for _, link := range links {
		response, err := Search(link)
		if err != nil {
			results[link] = "Error processing link: " + link
		} else {
			results[link] = response
		}
	}
	return results
}
