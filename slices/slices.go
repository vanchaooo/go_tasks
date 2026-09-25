package slices

import (
	"slices"
	"errors"
	"fmt"
)


// 1
func Sum(s []int) int {
	res := 0
	if len(s) != 0 {
		for _, v := range s {
			res += v
		}
		return res
	} else {
		return 0
	}
}


// 2
func Min(s []int) (int, bool) {
	if len(s) != 0 {
		min := s[0]
		for _, v := range s {
			if min > v {
				min = v
			}
		}
		return min, true
	} else {
		return 0, false
	}
}

// 3
func Max(s []int) (int, bool) {
	if len(s) != 0 {
		min := s[0]
		for _, v := range s {
			if min < v {
				min = v
			}
		}
		return min, true
	} else {
		return 0, false
	}
}


// 4
func CountEven(s []int) int {
	count := 0
	for _, v := range s {
		if v % 2 == 0 {
			count++
		}
	}
	return count
}


// 5
func Contains(s []int, target int) bool {
	var flag bool
	for _, v := range s {
		if v == target {
			flag = true
		}
	}
	return flag
}


// 6
func CountOccurrences(s []int, target int) int {
	res := 0
	for _, v := range s {
		if v == target {
			res++
		}
	}
	return res
}


// 7
func IndexOf(s []int, target int) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}


// 8
func IndexesOf(s []int, target int) []int {
	list := []int{}
	for i, v := range s {
		if v == target {
			list = append(list,i)
		}
	}
	return list
}


// 9
func ReplaceAll(s []int, old, new int) {
	for i, v := range s {
		if v == old {
			s[i] = new
		}
	}
}


// 10
func Reverse(s []int) {
	slices.Reverse(s)
}


// 11
func SwapPairs(s []int) {
	for i := 0; i+1 < len(s); i += 2 {
		s[i], s[i+1] = s[i+1], s[i]
	}
}


// 12
func ReverseRange(s []int, left, right int) error {
	fmt.Println(left, right, len(s))
	if left < 0 || right >= len(s) || left > right {
		return errors.New("Invalid range.")
	}
	
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	// slices.Reverse(s[left:right+1])

	return nil
}


// 13
func RemoveAt(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return nil, errors.New("Index out of range.")
	}
	s = append(s[:i], s[i+1:]...)
	// slices.Delete(s, i, i+1)
	return s, nil
}


// 14
func RemoveAtFast(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return nil, errors.New("Index out of range.")
	}

	last := len(s) - 1
	s[i] = s[last]
	s = s[:last]

	return s, nil
}


// 15
func RemoveAll(s []int, target int) []int {
	newSlice := []int{}
	for _, v := range s {
		if v != target {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}


// 16
func RemoveAllInPlace(s []int, target int) []int {
	write := 0
	for i, v := range s {
		if v != target {
			s[write] = v
			write++
		}
		fmt.Println(i, v, target, write)
	}
	return s[:write]
}


// 17
func KeepEvenInPlace(s []int) []int {
	write := 0
	for _, v := range s {
		if v % 2 == 0 {
			s[write] = v
			write++
		}
	}
	return s[:write]
}


// 18
func MoveZerosToEnd(s []int) {
	write := 0
	for _, v := range s {
		if v != 0 {
			s[write] = v
			write++
		}
	}


	for i, _ := range s[write:] {
		s[write+i] = 0
		fmt.Println(i)
		fmt.Println(s)
	}
}


// 19
func CompactNonZero(s []int) []int {
	// list := []int{}

	// for _, v := range s {
	// 	if v != 0 {
	// 		list = append(list, v)
	// 	}
	// }
	// return list

	write := 0
    for i := range s {
        if s[i] != 0 {
            s[write] = s[i]
            write++
        }
    }
    return s[:write]
}


// 20
func InsertAt(s *[]int, i, value int) ([]int, error) {
	if i > len(*s) {
		return nil, errors.New("Index out of range")
	}

	*s = slices.Insert(*s, i, value)

	return *s, nil
}
