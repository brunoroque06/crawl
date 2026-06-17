package main

import (
	"encoding/xml"
)

type atom struct{ url string }

func (a atom) getUrl() string {
	return a.url
}
func (a atom) parse(body string) ([]item, error) { return atomParse(body) }

type atomFeed struct {
	Entries []struct {
		Title string `xml:"title"`
		Link  struct {
			Href string `xml:"href,attr"`
		} `xml:"link"`
		Updated desTime `xml:"updated"`
	} `xml:"entry"`
}

func atomParse(body string) ([]item, error) {
	var feed atomFeed
	err := xml.Unmarshal([]byte(body), &feed)
	if err != nil {
		return nil, err
	}
	var items []item
	for _, entry := range feed.Entries {
		items = append(items, item{pub: entry.Updated.Time, title: entry.Title, url: entry.Link.Href})
	}
	return items, nil
}
