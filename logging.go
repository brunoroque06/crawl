package main

import (
	"fmt"
	"os"
)

func stderr(s ...any) {
	fmt.Fprintln(os.Stderr, s...)
}

func stdout(s ...any) {
	fmt.Println(s...)
}

func printReport(report []Report) {
	for _, feed := range report {
		stdout(feed.Name)
		if feed.Error != nil {
			stdout("  error: ", feed.Error.Error())
			stdout()
			continue
		}
		for _, item := range feed.Items {
			stdout("  ", item.Title, cleanUrl(item.Url))
		}
		stdout()
	}
}
