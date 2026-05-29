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
