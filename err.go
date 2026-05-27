package main

import "fmt"

func errorf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}
