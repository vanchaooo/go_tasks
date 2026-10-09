package main

import (
	"fmt"
	// "time"
	// "packages/maps"
	// "packages/strings"
	// "packages/slices"
	"packages/concurrency"
)

func main() {
	nums := []int{2, 9, 3}
	check := concurrency.CountSteps(nums)
	fmt.Println(check)
}