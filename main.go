package main

const (
	last        = 8
	cutOffHours = 36
)

func main() {
	feeds := []Feed{}

	reports := crawl(feeds)

	print(reports)
}
