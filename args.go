package main

import (
	"flag"
	"os"
	"strings"
)

var sources = []struct {
	name  string
	usage string
	build func(value string) Source
}{
	{"gnews", "name=query", func(v string) Source { return Gnews{Query: v} }},
	{"reddit", "name=subreddit", func(v string) Source { return Reddit{Sub: v} }},
	{"rss", "name=url", func(v string) Source { return Rss{Url: v} }},
}

func parseFeeds() []Feed {
	var feeds []Feed

	for _, def := range sources {
		flag.Func(def.name, def.usage+" (repeatable)", func(s string) error {
			name, val, ok := strings.Cut(s, "=")
			if !ok {
				return errorf("expected %s, got %q", def.usage, s)
			}
			feeds = append(feeds, Feed{Name: name, Src: def.build(val)})
			return nil
		})
	}
	flag.Parse()

	if len(feeds) == 0 {
		stderr("no feeds specified")
		flag.Usage()
		os.Exit(1)
	}

	return feeds
}
