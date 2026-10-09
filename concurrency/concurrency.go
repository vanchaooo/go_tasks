package concurrency

import (
	// "sync"
	"fmt"
)

// 1
func FirstTask() <- chan string{
	ch := make(chan string)
	go func() {
		ch <- "hello"
	}()
	return ch
}

// 2
func SendNumber(n int, ch chan int) {
	ch <- n
}

// 3
func Double(n int) int {
	ch := make(chan int)
	go func(n int) {
		ch <- n*2
	}(n)
	result := <- ch
	return result
}

// 4
func Greet(name string) string {
	ch := make(chan string)
	go func() {
		ch <- "Привет, "+name
	}()
	result := <- ch
	return result
}

// 5
func SumAsync(nums []int) int {
	ch := make(chan int)
	go func() {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		ch <- sum
	}()
	result := <- ch
	return result
}

// 6
func RuneCountAsync(s string) int {
	ch := make(chan int)
	go func() {
		runes := []rune{}
		for char := range s {
			runes = append(runes, rune(char))
		}
		ch <- len(runes)
	}()
	result := <- ch
	return result
}

// 7
func PrintSignal() chan bool {
	done := make(chan bool)
	go func() {
		fmt.Println("Готово")
		done <- true
	}()
	return done
}

// 8
type Rectangle struct {Width, Height int}
func (r Rectangle) SendArea(ch chan int) {
	go func ()	{
		ch <- r.Height*r.Width
	}()
}

// 9
type PairInfo struct { Sum int; Equal bool }
func AnalyzePair(a, b int) PairInfo {
	ch := make(chan PairInfo)
	go func() {
		pair := PairInfo {
			Sum: a+b,
			Equal: a == b,
		}
		ch <- pair
	}()
	result := <- ch
	return result
}

// 10
func SumNonNegative(nums []int) int {
	ch := make(chan int)
	go func() {
		sum := 0
		for _, v := range nums {
			if v >= 0 {
				sum += v
			}
		}
		ch <- sum
	}()
	result := <- ch
	return result
}

// 11
func ThreeMessages() <- chan int {
	ch := make(chan int)
	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
	}()
	return ch
}

// 12
func CopyThroughChannel(nums []int) []int {
	ch := make(chan int)
	res := make([]int, 0, len(nums))
	go func() {
		for _, v := range nums {
			ch <- v
		}
	}()
	for i:=0; i<len(nums); i++ {
		res = append(res, <- ch)
	}
	return res
}

// 13
func ReverseThroughChannel(nums []int) []int {
	ch := make(chan int)
	res := make([]int, 0, len(nums))
	go func() {
		for i:=len(nums)-1; i>=0; i-- {
			ch <- nums[i]
		}
	}()
	for i:=0; i<len(nums); i++ {
		res = append(res, <- ch)
	}
	return res
}

// 14
func FiveDone(ch chan struct{}) {
	go func() {
		for i:=1; i<6; i++ {
			fmt.Println(i)
		}
		ch <- struct{}{}
	}()
}

// 15
func TwoResults(a, b int) []int {
	ch := make(chan int)
	res := []int{}
	go func() {
		ch <- a
	}()
	go func() {
		ch <- b
	}()
	res = append(res, <- ch)
	res = append(res, <- ch)
	return res
}

// 16
func CollectIDs(n int) []int {
	ch := make(chan int)
	res := []int{}
	for i:=0; i<n; i++ {
		go func(id int) {
			ch <- id
		}(i)
	}
	for i:=0; i<n; i++ {
		res = append(res, <- ch)
	}
	return res
}

// 17
func SumHalves(nums []int) int {
	ch := make(chan int)
	result := 0
	a := nums[:len(nums)/2]
	b := nums[len(nums)/2:]
	go func() {
		total := 0
		for _, v := range a {
			total += v
		}
		ch <- total
	}()
	go func() {
		total := 0
		for _, v := range b {
			total += v
		}
		ch <- total
	}()
	result += <- ch
	result += <- ch
	return result
}

// 18
func Echo(x int) int {
	ch := make(chan int)
	go func () {
		ch <- x	
	}()
	return <- ch
}

// 19
func ReceiveText(text string) string {
	var ch chan string
	ch = make(chan string)
	go func() {
		ch <- text
	}()
	return <- ch
}

// 20
func CountSteps(steps []int) int {
	if len(steps) > 100 {
		return 0
	}
	ch := make(chan int)
	result := 0
	for _, v := range steps {
		go func(num int) {
			count := 0
			for i:=0; i<num; i++ {
				count++
			}
			ch <- count
		}(v)
	}
	for i:=0; i<len(steps); i++ {
		result += <- ch
	}
	return result
}