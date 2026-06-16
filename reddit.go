package main

import (
	"fmt"
)

type reddit struct{ sub string }

func (r reddit) getUrl() string {
	return fmt.Sprintf("https://www.reddit.com/r/%s/.rss", r.sub)
}
func (r reddit) parse(body string) ([]item, error) { return atomParse(body) }
