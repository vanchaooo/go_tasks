package main

import (
	"fmt"
	// "packages/maps"
	"packages/strings"
	// "packages/slices"
)

func main() {
	text := "римировой"
	check := "мир"

	fmt.Println(strings.Contains(text, check))
}