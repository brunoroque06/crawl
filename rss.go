package main

import (
	"encoding/xml"
	"time"
)

type Rss struct{ URL string }

func (r Rss) Url() string                       { return r.URL }
func (r Rss) Parse(body string) ([]Item, error) { return rss(body) }

type RssResponse struct {
	Channel struct {
		Items []struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func rss(body string) ([]Item, error) {
	var r RssResponse
	err := xml.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	var items []Item
	for _, entry := range r.Channel.Items {
		pubDate, err := time.Parse(time.RFC1123Z, entry.PubDate)
		if err != nil {
			pubDate, err = time.Parse(time.RFC1123, entry.PubDate)
			if err != nil {
				continue
			}
		}
		if now.Sub(pubDate) < cutOffHours*time.Hour {
			items = append(items, Item{Title: entry.Title, Url: entry.Link})
			if len(items) == last {
				break
			}
		}
	}
	return items, nil
}
