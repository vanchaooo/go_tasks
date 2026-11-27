package main

import (
	"fmt"
	// "slices"
	// "time"
	// "packages/maps"
	// "packages/strings"
	// "packages/slices"
	"packages/concurrency"
)

func main() {
	count := 0
	for range concurrency.EmptyInts() {
		count++
	}
	fmt.Println(count)
}