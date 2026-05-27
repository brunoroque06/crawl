package main

type Gnews struct{ Query string }

func (g Gnews) Url() string {
	return "https://news.google.com/rss/search?hl=en-US&gl=US&ceid=US%3Aen&q=" + g.Query
}
func (g Gnews) Parse(body string) ([]Item, error) { return rss(body) }
