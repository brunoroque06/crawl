package main

func print(report []Report) {
	for _, feed := range report {
		println(feed.Name)
		if feed.Error != nil {
			println("error: ", feed.Error.Error())
			continue
		}
		for _, item := range feed.Items {
			println("  ", item.Title, cleanUrl(item.Url))
		}
		println()
	}
}
