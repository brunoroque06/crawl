package main

import (
	"fmt"
	"os"
)

func stderr(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}

func stdout(a ...any) {
	fmt.Println(a...)
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
