package main

const (
	last        = 8
	cutOffHours = 36
)

func main() {
	feeds := parseFeeds()
	reports := crawl(feeds)
	report(reports)
}
