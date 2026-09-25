package strings

import (
	"errors"
	"fmt"
	// "slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// 1
func StringSize(s string) (bytes int, runes int) {
	bytes = len(s)
	runes = utf8.RuneCountInString(s)
	return bytes, runes
}


// 2
func IsASCII(s string) bool {
	for i := range s {
		if i <= 127 {
			return true
		}
	}
	return false
}


// 3
func FirstRune(s string) (rune, bool) {
	if len(s) == 0 {
		fmt.Println(errors.New("Error: the string is empty"))
		return 0, false
	}
	
	r, _ := utf8.DecodeRuneInString(s)
	return r, true
}


// 4
func LastRune(s string) (rune, bool) {
	if len(s) == 0 {
		fmt.Println(errors.New("Error: the string is empty"))
		return 0, false
	}
	
	r, _ := utf8.DecodeLastRuneInString(s)
	return r, true
}


// 5
func Reverse(s string) string {
	runes := []rune(s)

	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	result := string(runes)
	return result
}


// 6
func RemoveRune(s string, target rune) string {
	runes := []rune(s)
	
	runa := 0 
	for _, v := range runes {
		if v != target {
			runes[runa] = v
			runa++
		}
	}

	result := string(runes[:runa])
	return result
}


// 7
func ReplaceRune(s string, old, new rune) string {
	runes := []rune(strings.ToLower(s))

	for i, v := range runes {
		if v == old {
			runes[i] = new
		}
	}
	
	result := string(runes)
	return result
}


// 8
func IsPalindrome(s string) bool {
	runes := []rune(strings.ToLower(s))
	filter := make([]rune, 0, len(runes))

	for _, v := range runes {
		if !unicode.IsSpace(v) {
			filter = append(filter, v)
		}
	}

	left, right := 0, len(filter)-1
	for left < right {
		if filter[left] != filter[right] {
			return false
		}
		left++
		right--
	}
	return true
}


// 9
func TrimUnicodeSpace(s string) string {
	runes := []rune(s)
    start, end := 0, len(runes)

    for start < end && unicode.IsSpace(runes[start]) {
        start++
    }

    for end > start && unicode.IsSpace(runes[end-1]) {
        end--
    }

    return string(runes[start:end])
}


// 10
func NormalizeSpaces(s string) string {
	runes := []rune(s)
    result := make([]rune, 0, len(runes))
    
    prevSpace := false
    for _, v := range runes {
        if unicode.IsSpace(v) {
            prevSpace = true
            continue
        }
        if prevSpace && len(result) > 0 {
            result = append(result, ' ')
        }
        result = append(result, v)
        prevSpace = false
    }
    
    return string(result)
}


// 11
func Contains(s, sub string) bool {
	check := []rune(s)

	for i := 0; i < len(check)-len([]rune(sub)); i++ {
		if strings.HasPrefix(string(check[i:]), sub) {
			return true
		}
	}
	return false
}


// 12
func IndexOf(s, sub string) int {
	if len(sub) == 0 {
		return 0
	}
	if len(s) < len(sub) {
		return -1
	}

	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}

	return -1
}


// 13
func CountOccurrences(s, sub string) (int, error) {
	if len(sub) == 0 {
		return 0, errors.New("подстрока не может быть пустой")
	}

	count := 0
	i := 0

	for i <= len(s)-len(sub) {
		if s[i:i+len(sub)] == sub {
			count++
			i += len(sub)
		} else {
			i++
		}
	}

	return count, nil
}