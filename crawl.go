package main

import (
	"sync"
	"time"
)

type feed struct {
	name string
	src  source
}

type report struct {
	name  string
	items []item
	error error
}

func fetch(source source) ([]item, error) {
	body, err := get(source.getUrl(), nil)
	if err != nil {
		return nil, err
	}
	items, err := source.parse(body)
	if err != nil {
		return nil, err
	}
	var filtered []item
	now := time.Now()
	for _, i := range items {
		if now.Sub(i.pub) < cutOffHours*time.Hour {
			filtered = append(filtered, i)
			if len(filtered) == last {
				break
			}
		}
	}
	return filtered, nil
}

func crawl(feeds []feed) []report {
	var mu sync.Mutex
	var reports []report
	var wg sync.WaitGroup

	for _, f := range feeds {
		wg.Go(func() {
			items, err := fetch(f.src)
			mu.Lock()
			reports = append(reports, report{name: f.name, items: items, error: err})
			mu.Unlock()
		})
	}

	wg.Wait()
	return reports
}

func deliver(reports []report) {
	for _, feed := range reports {
		stdout(feed.name)
		if feed.error != nil {
			stdout("  error:", feed.error.Error())
			stdout()
			continue
		}
		for _, item := range feed.items {
			stdout(" ", item.title, cleanUrl(item.url))
		}
		stdout()
	}
}
