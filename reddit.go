package main

import (
	"encoding/json"
	"fmt"
)

type Reddit struct{ Sub string }

func (r Reddit) GetUrl() string {
	return fmt.Sprintf("https://www.reddit.com/r/%s/hot.json?limit=%d", r.Sub, last)
}
func (r Reddit) Parse(body string) ([]Item, error) { return reddit(body) }

type RedditResponse struct {
	Data struct {
		Children []struct {
			Data struct {
				Title     string `json:"title"`
				Permalink string `json:"permalink"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func reddit(body string) ([]Item, error) {
	var r RedditResponse
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}
	var items []Item
	for _, child := range r.Data.Children {
		items = append(items, Item{
			Title: child.Data.Title,
			Url:   "https://www.reddit.com" + child.Data.Permalink,
		})
	}
	return items, nil
}
