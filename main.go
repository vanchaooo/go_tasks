package main

import (
	"fmt"
	"packages/maps"
	// "packages/strings"
	// "packages/slices"
)

func main() {


	fmt.Println(maps.SymmetricDifference(
		[]int{1, 1, 2},
		[]int{2, 3, 3},
	))
}