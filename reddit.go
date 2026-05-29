package main

import (
	"fmt"
)

type Reddit struct{ Sub string }

func (r Reddit) GetUrl() string {
	return fmt.Sprintf("https://www.reddit.com/r/%s/.rss", r.Sub)
}
func (r Reddit) Parse(body string) ([]Item, error) { return atom(body) }
