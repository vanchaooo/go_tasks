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
	ch := make(chan int)
	go func() {
		for i := 1; i < 6; i++ {
			ch <- i
		}
		close(ch)
	}()
	fmt.Println(concurrency.CollectRange(ch))
}