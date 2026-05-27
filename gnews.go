package main

import "net/url"

type Gnews struct{ Query string }

func (g Gnews) GetUrl() string {
	query := url.QueryEscape(g.Query)
	return "https://news.google.com/rss/search?hl=en-US&gl=US&ceid=US%3Aen&q=" + query
}
func (g Gnews) Parse(body string) ([]Item, error) { return rss(body) }
