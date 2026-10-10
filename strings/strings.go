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
	for _, r := range s {
		fmt.Printf("символ: %c, числовое значение: %d\n", r, r)
		if r > 127 {
			return false
		}
	}
	return true
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
	runes := []rune(s)

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

	for i := 0; i < len(check)-len([]rune(sub))+1; i++ {
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

// 14
func CountOverlapping(s, sub string) (int, error) {
	if len(s) == 0 || len(sub) == 0 {
		return 0, errors.New("Строка и подстрока не могут быть пустыми")
	}
	if len(sub) > len(s) {
		return 0, errors.New("Подстрока не может быть больше основной строки")
	}

	count := 0
	for i:=0; i<=len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			count++
		}
	}
	return count, nil
}

// 15
func CommonPrefix(a, b string) string {
	runesA := []rune(a)
	runesB := []rune(b)

	minLen := len(runesA)
	if len(runesB) < minLen {
		minLen = len(runesB)
	}

	prefix := []rune{}
	for i:=0; i<minLen; i++ {
		if runesA[i] == runesB[i] {
			prefix = append(prefix, runesA[i])
		} else {
			break
		}
	}
	return string(prefix)
}

// 16
func LongestCommonPrefix(items []string) string {
	if len(items) == 0 {
		return ""
	}

	prefixRunes := []rune(items[0])
	for i:=1; i < len(items); i++ {
		currentRunes := []rune(items[i])
		j := 0
		for j < len(prefixRunes) && j < len(currentRunes) && prefixRunes[j] == currentRunes[j] {
			j++
		}

		prefixRunes = prefixRunes[:j]
		if len(prefixRunes) == 0 {
			return ""
		}
	}
	return string(prefixRunes)
}

// 17
func Split(s, sep string) ([]string, error) {
	if s == "" {
		return []string{}, errors.New("Строка не может быть пустой")
	}

	var result []string
	start := 0
	sepLen := len(sep)
	for i := 0; i <= len(s)-sepLen; {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			i += sepLen
			start = i
		} else {
			i++
		}
	}
	result = append(result, s[start:])
	return result, nil
}

// 18
func Join(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	if len(parts) == 1 {
		return parts[0]
	}

	var builder strings.Builder
	builder.WriteString(parts[0])
	for _, part := range parts[1:] {
		builder.WriteString(sep)
		builder.WriteString(part)
	}
	return builder.String()
}

// 19
func TrimSet(s, cutset string) string {
	if len(s) == 0 || len(cutset) == 0 {
		return s
	}

	runes := []rune(s)
	start := 0
	end := len(runes) - 1
	for start <= end && strings.ContainsRune(cutset, runes[start]) {
		start++
	}
	for end >= start && strings.ContainsRune(cutset, runes[end]) {
		end--
	}
	return string(runes[start : end+1])
}

// 20
func CollapseRuns(s string) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	result := make([]rune, 0, len(runes))
	result = append(result, runes[0])
	for i := 1; i < len(runes); i++ {
		if runes[i] != result[len(result)-1] {
			result = append(result, runes[i])
		}
	}
	return string(result)
}

// 21
func EncodeRLE(s string) string {
	if s == "" {
		return ""
	}

	check := []rune(s)
	current := check[0]
	count := 1
	var builder strings.Builder
	for i:=0; i<len(check); i++ {
		if check[i] == current {
			count++
		} else {
			builder.WriteString(fmt.Sprintf("%d:%c;", count, current))
			current = check[i]
			count = 1
		}
	}
	builder.WriteString(fmt.Sprintf("%d:%c;", count, current))
	return builder.String()
}

// 22
func DecodeRLE(encoded string) (string, error) {
	if encoded == "" {
		return "", errors.New("Пустая строка")
	}

	var builder strings.Builder
	totalLen := 0
	parts, err := Split(encoded, ";")
	if err != nil {
		return "", err
	}

	for _, part := range parts {
		var count int
		var r rune
		n, err := fmt.Sscanf(part, "%d:%c", &count, &r)
		
		if err != nil || n != 2 {
			return "", errors.New("повреждённая запись")
		}
		if count <= 0 {
			return "", errors.New("нулевое или отрицательное количество")
		}

		totalLen += count
		for i := 0; i < count; i++ {
			builder.WriteRune(r)
		}
	}
	return builder.String(), nil
}

// 23
func WordCount(s string) int {
	return len(strings.Fields(s))
}

// 24
func WordFrequency(s string) map[string]int {
	if s == "" {
		return nil
	}

	result := make(map[string]int)
	stroka := strings.ToLower(s)
	words := strings.FieldsFunc(stroka, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	for _, word := range words {
		result[word]++
	}
	return result
}

// 25
func FirstUniqueRune(s string) (rune, bool) {
	if s == "" {
		return 0, false
	}

	check := make(map[rune]int)
	for _, v := range s {
		check[v]++
	}
	for _, v := range s {
		if check[v] == 1 {
			return v, true
		}
	}
	return 0, false
}