package main

import (
	"encoding/xml"
)

type Atom struct{ Url string }

func (a Atom) GetUrl() string {
	return a.Url
}
func (r Atom) Parse(body string) ([]Item, error) { return atom(body) }

type AtomFeed struct {
	Entries []struct {
		Title string `xml:"title"`
		Link  struct {
			Href string `xml:"href,attr"`
		} `xml:"link"`
		Updated DesTime `xml:"updated"`
	} `xml:"entry"`
}

func atom(body string) ([]Item, error) {
	var feed AtomFeed
	err := xml.Unmarshal([]byte(body), &feed)
	if err != nil {
		return nil, err
	}
	var items []Item
	for _, entry := range feed.Entries {
		items = append(items, Item{Pub: entry.Updated.Time, Title: entry.Title, Url: entry.Link.Href})
	}
	return items, nil
}
