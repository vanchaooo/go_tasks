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
	slice := []int{1, 4, 2, 5, 3}
	check := concurrency.BufferedCopy(slice)
	fmt.Println(check)
}