package maps

import (
	"errors"
	"maps"
	"sort"
	"unicode"
	"strings"
)

// 1
func CountInts(nums []int) map[int]int {
	result := make(map[int]int)

	for _, v := range nums {
		result[v]++
	}

	return result
}


// 2
func LastIndex(words []string) map[string]int {
	result := make(map[string]int)

	for i, v := range words {
		result[v] = i
	}

	return result
}


// 3
func Unique(nums []int) []int {
	check :=  make(map[int]struct{})
	result := make([]int, 0, len(nums))

	for _, v := range nums {
		if _, ok := check[v]; !ok {
			check[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}


// 4
func ContainsAll(have, need []string) bool {
	check := make(map[string]struct{})
	for _, v := range have {
		check[v] = struct{}{}
	}

	for _, v := range need {
		if _, ok := check[v]; !ok {
			return false
		} 
	}
	return true
}


// 5
func SortedKeys(m map[string]int) []string {
	result := make([]string, 0, len(m))

	for key := range m {
		result = append(result, key)
	}

	sort.Strings(result)
	return result
}


// 6
func SumValues(m map[string]int) int {
	if len(m) == 0 {
		return 0
	}

	result := 0
	for _, value := range m {
		result += value
	}

	return result
}


// 7
func Invert(m map[string]int) (map[int]string, error) {
	result := make(map[int]string, len(m))

	for k, v := range m {
		if _, ok := result[v]; ok {
			return nil, errors.New("Ошибка! Значение дублированно.")
		}
		result[v] = k
	}

	return result, nil
}


// 8
func MergeCounts(a, b map[string]int) map[string]int {
	result := make(map[string]int)

	for k, v := range a {
		result[k] = v
	}

	for k, v := range b {
		result[k] += v
	}

	return result
}


// 9
func  DeleteZeroValues(m map[string]int) {
	for k, v := range m {
		if v == 0 {
			delete(m, k)
		}
	}
}


// 10
func Clone(m map[string][]int) map[string][]int {
	result := make(map[string][]int, len(m))

	for key, slice := range m {
		newSlice := make([]int, len(slice))
		copy(newSlice, slice)

		result[key] = newSlice
	}

	return result
}


// 11
func Equal(a, b map[string]int) bool {
	return maps.Equal(a, b)
}


// 12
func Difference(a, b []int) []int {
	result := make([]int, len(a))
	checkA := make(map[int]struct{}, len(a))
	checkB := make(map[int]struct{}, len(b))

	for _, v := range b {
		checkB[v] = struct{}{}
	}

	for _, v := range a {
		if _, ok := checkB[v]; ok {
			continue
		}

		if _, ok := checkA[v]; ok {
			continue
		}
		checkA[v] = struct{}{}
		result = append(result, v)
	}

	return result
}


// 13
func Intersection(a, b []string) []string {
	result := make([]string, 0, len(a))

	checkA := make(map[string]struct{})
	checkB := make(map[string]struct{})

	for _, v := range b {
		checkB[v] = struct{}{}
	}

	for _, v := range a {
		if _, ok := checkB[v]; !ok {
			continue
		}
		if _, ok := checkA[v]; ok {
			continue
		}

		checkA[v] = struct{}{}
		result = append(result, v)
	}

	return result
}


// 14
func SymmetricDifference(a, b []int) []int {
	result := make([]int, 0, len(a))

	checkA := make(map[int]struct{})

	for _, v := range a {
		checkA[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := checkA[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}


// 15
func UniqueFold(words []string) []string {
	result := []string{}
	mapa := make(map[string]struct{})
	
	for _, v := range words {
		runes := make([]rune, 0, len(words))
		for _, char := range v {
			runes = append(runes, unicode.ToLower(char))
		}
		mapa[string(runes)] = struct{}{}
	}

	for k, _ := range mapa {
		result = append(result, k)
	}

	return result
}


// 16
func  MostFrequent(nums []int) (value int, count int, ok bool) {
	if len(nums) == 0 {
		return 0, 0, false
	}

	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	maxcount := -1
	bestvalue := 0

	for num, cnt := range freq {
		if cnt > maxcount || (cnt == maxcount && num < bestvalue) {
			maxcount = cnt
			bestvalue = num
		}
	}

	return bestvalue, maxcount, true
}


// 17
func FirstUniqueRune(s string) (rune, bool) {
	if s == "" {
		return 0, false
	}

	freq := make(map[rune]int)
	for _, v := range s {
		freq[v]++
	}

	for _, v := range s {
		if freq[v] == 1 {
			return v, true
		}
	}

	return 0, false
}


// 18
func  AreAnagrams(a, b string) bool {
	if a == "" || b == "" {
		return false
	}

	freq := make(map[rune]int)

	for _, v := range strings.ToLower(a) {
		if unicode.IsSpace(v) {
			continue
		}
		freq[v]++
	}

	for _, v := range strings.ToLower(b) {
		if unicode.IsSpace(v) {
			continue
		}
		freq[v]--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}


// 19
func WordFrequency(text string) map[string]int {
	freq := make(map[string]int)
	var word strings.Builder

	flush := func() {
		if word.Len() > 0 {
			freq[word.String()]++
			word.Reset()
		}
	}

	for _, v := range text {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			word.WriteRune(unicode.ToLower(v))
		} else {
			flush()
		}
	}
	flush()

	return freq
}


// 20
func TopKWords(text string, k int) []string {
	words := strings.Fields(strings.ToLower(text))
	count := make(map[string]int)
	unique := make([]string, 0, len(count))

	for _, v := range words {
		count[v]++
	}

	for v := range count {
		unique = append(unique, v)
	}

	sort.Slice(unique, func(i, j int) bool {
		if count[unique[i]] != count[unique[j]] {
			return count[unique[i]] > count[unique[j]]
		}
		return unique[i] < unique[j]
	})

	if k > len(unique) {
		k = len(unique)
	}
	return unique[:k]
}