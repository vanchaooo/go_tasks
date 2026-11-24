package concurrency

import (
	// "sync"
	"fmt"
	"unicode/utf8"
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

// 21
func Bufer() []int {
	ch := make(chan int, 3)

	ch <- 4
	ch <- 5
	ch <- 6

	a := <- ch
	b := <- ch
	c := <- ch

	return []int{a, b, c}
}

// 22
func ReBufer() []int {
	ch := make(chan int, 2)
	result := make([]int, 0, 3)

	ch <- 1
	ch <- 2

	a := <- ch
	result = append(result, a)

	ch <- 3
	b := <- ch
	c := <- ch
	result = append(result, b, c)

	fmt.Printf("Len после чтений: %d\n", len(ch))
	return result
}

// 23
func BufferedCopy(nums []int) []int {
	ch := make(chan int, len(nums))
	result := make([]int, 0, len(nums))

	for _, v := range nums {
		ch <- v
	}

	for range len(nums) {
		result = append(result, <- ch)
	}

	return result
}

// 24
func BufferedSequence(n int) []int {
	if n == 0 {
		return []int{}
	}

	ch := make(chan int, 2)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()

	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-ch)
	}
	return result
}

// 25
type Point3 struct { X, Y, Z int }

func EncodePoint(p Point3) [3]int {
	ch := make(chan int, 1)
	go func() {
		ch <- p.X
		ch <- p.Y
		ch <- p.Z
	}()

	result := [3]int{}
	for i := 0; i < 3; i++ {
		result[i] = <-ch
	}

	return result
}

// 26
func StartSquare(n int) chan int {
	ch := make(chan int, 1)

	go func() {
		ch <- n*n
	}()

	return ch
}

// 27
func DrainPrepared(ch chan int, count int) []int {
	if count == 0 {
		return []int{}
	}

	result := make([]int, 0, count)
	for range count {
		result = append(result, <- ch)
	}
	
	return result
}

// 28
func SumWithBuffer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ch := make(chan int, 1)
	go func() {
		for _, v := range nums {
			ch <- v
		}
	}()

	sum := 0
	for range nums {
		sum += <- ch
	}

	return sum
}

// 29
func WordBatch(words []string) int {
	if len(words) == 0 {
		return 0
	}

	ch := make(chan string, len(words))
	for _, v := range words {
		ch <- v
	}

	sum := 0
	for range words {
		sum += utf8.RuneCountInString(<- ch)
	}

	return sum
}

// 30
func BufferedPair(a, b string) map[string]int {
	ch := make(chan string, 2)
	result := make(map[string]int)

	go func() {
		ch <- a
	}()
	go func() {
		ch <- b
	}()

	for range 2 {
		result[<-ch]++
	}
	return result
}

// 31
func SendGreeting(name string, out chan<- string) {
	out <- "Привет, " + name
}

// 32
func ReceiveN(in <-chan int, n int) []int {
	if n == 0 {
		return []int{}
	}

	result := make([]int, 0, n)
	for range n {
		result = append(result, <- in)
	}
	return result
}

// 33
func SendAndClose(value int, out chan<- int) {
	out <- value
	close(out)
}

// 34
func main() {
	ch := make(chan int)

	go func() {
		ch <- 0
		ch <- 5
		ch <- 0
	}()

	var result []int
	for {
		v, ok := <- ch
		if !ok {
			break
		}
		result = append(result, v)
	}

	fmt.Println(result)
}

// 35
func CollectRange(in <-chan int) []int {
	result := make([]int, 0, len(in))
	for v := range in {
		result = append(result, v)
	}
	
	return result
}