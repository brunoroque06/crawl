package main

import (
	"fmt"
	"os"
)

func errorf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

func stderr(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}

func stdout(a ...any) {
	fmt.Println(a...)
}
