package main

import (
	"encoding/xml"
)

type rss struct{ url string }

func (r rss) getUrl() string                    { return r.url }
func (r rss) parse(body string) ([]item, error) { return rssParse(body) }

type rssResponse struct {
	Channel struct {
		Items []struct {
			Title   string  `xml:"title"`
			Link    string  `xml:"link"`
			PubDate desTime `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func rssParse(body string) ([]item, error) {
	var r rssResponse
	err := xml.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}
	var items []item
	for _, entry := range r.Channel.Items {
		items = append(items, item{pub: entry.PubDate.Time, title: entry.Title, url: entry.Link})
	}
	return items, nil
}
