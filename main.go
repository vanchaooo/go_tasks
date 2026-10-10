package main

import (
	"fmt"
	// "time"
	// "packages/maps"
	"packages/strings"
	// "packages/slices"
	// "packages/concurrency"
)

func main() {
	str := "kjzdkasd as"
	check:= strings.WordFrequency(str)
	fmt.Println(check)
}