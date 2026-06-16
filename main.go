package main

const (
	last        = 8
	cutOffHours = 25
)

func main() {
	feeds := parseFeeds()
	reports := crawl(feeds)
	deliver(reports)
}
