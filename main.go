package main

import (
	"fmt"
	// "slices"
	// "time"
	// "packages/maps"
	// "packages/strings"
	"packages/slices"
	// "packages/concurrency"
)

func main() {
	spisok := []int{1,2,3,3,4,5,5}
	check := slices.UniqueInPlace(spisok)
	fmt.Println(check)
}