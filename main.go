package main

func main() {
	feeds := parseFeeds()
	reports := crawl(feeds)
	deliver(reports)
}
