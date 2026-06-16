package main

import "net/url"

type gnews struct{ query string }

func (g gnews) getUrl() string {
	query := url.QueryEscape(g.query)
	return "https://news.google.com/rss/search?hl=en-US&gl=US&ceid=US%3Aen&q=" + query
}
func (g gnews) parse(body string) ([]item, error) { return rssParse(body) }
