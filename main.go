package main

import (
	"fmt"
	// "time"
	"packages/maps"
	// "packages/strings"
	// "packages/slices"
	// "packages/concurrency"
)

func main() {
	str := []string{"dog", "cat", "asdasdasd", "chipicahapa", "dog", "asd"}
	check:= maps.GroupAnagrams(str)
	fmt.Println(check)
}