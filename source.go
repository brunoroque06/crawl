package main

import (
	"time"
)

type item struct {
	title string
	url   string
	pub   time.Time
}

type source interface {
	getUrl() string
	parse(body string) ([]item, error)
}
