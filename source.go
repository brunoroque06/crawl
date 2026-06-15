package main

import (
	"time"
)

type Item struct {
	Title string
	Url   string
	Pub   time.Time
}

type Source interface {
	GetUrl() string
	Parse(body string) ([]Item, error)
}
