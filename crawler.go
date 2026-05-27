package main

import (
	"sync"
)

type Feed struct {
	Name string
	Src  Source
}

type Item struct {
	Title string
	Url   string
}

type Source interface {
	GetUrl() string
	Parse(body string) ([]Item, error)
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
	return items, nil
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
