package main

import (
	"fmt"
	"os"
)

func fail(msg string, args ...any) {
	fmt.Printf(msg, args...)
	os.Exit(1)
}
