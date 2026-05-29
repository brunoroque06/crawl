package main

import (
	"encoding/xml"
)

type Rss struct{ Url string }

func (r Rss) GetUrl() string                    { return r.Url }
func (r Rss) Parse(body string) ([]Item, error) { return rss(body) }

type RssResponse struct {
	Channel struct {
		Items []struct {
			Title   string  `xml:"title"`
			Link    string  `xml:"link"`
			PubDate DesTime `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func rss(body string) ([]Item, error) {
	var r RssResponse
	err := xml.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}
	var items []Item
	for _, entry := range r.Channel.Items {
		items = append(items, Item{Pub: entry.PubDate.Time, Title: entry.Title, Url: entry.Link})
	}
	return items, nil
}
