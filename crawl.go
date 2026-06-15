package main

import (
	"sync"
	"time"
)

type Feed struct {
	Name string
	Src  Source
}

type Report struct {
	Name  string
	Items []Item
	Error error
}

func fetch(source Source) ([]Item, error) {
	body, err := get(source.GetUrl(), nil)
	if err != nil {
		return nil, err
	}
	items, err := source.Parse(body)
	if err != nil {
		return nil, err
	}
	var filtered []Item
	now := time.Now()
	for _, i := range items {
		if now.Sub(i.Pub) < cutOffHours*time.Hour {
			filtered = append(filtered, i)
			if len(filtered) == last {
				break
			}
		}
	}
	return filtered, nil
}

func crawl(feeds []Feed) []Report {
	var mu sync.Mutex
	var reports []Report
	var wg sync.WaitGroup

	for _, f := range feeds {
		wg.Go(func() {
			items, err := fetch(f.Src)
			mu.Lock()
			reports = append(reports, Report{Name: f.Name, Items: items, Error: err})
			mu.Unlock()
		})
	}

	wg.Wait()
	return reports
}

func report(reports []Report) {
	for _, feed := range reports {
		stdout(feed.Name)
		if feed.Error != nil {
			stdout("  error:", feed.Error.Error())
			stdout()
			continue
		}
		for _, item := range feed.Items {
			stdout(" ", item.Title, cleanUrl(item.Url))
		}
		stdout()
	}
}
