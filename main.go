package main

import (
	"fmt"
	// "slices"
	// "time"
	// "packages/maps"
	// "packages/strings"
	// "packages/slices"
	// "packages/concurrency"
)

func main() {
	// slice := 5
	// check := concurrency.BufferedSequence(slice)
	// fmt.Println(check)

	ch := make(chan int)

    go func() {
		ch <- 10
	}()
    fmt.Println(<-ch)
}

