package maps

import (
	"errors"
	"fmt"
	"maps"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Row struct {
	ID int
	Name string
	Score int
}

type ChangeKind string

const (
    Added ChangeKind = "added"
    Deleted ChangeKind = "deleted"
    Changed ChangeKind = "changed"
)

type Change struct {
    Key string
    Kind ChangeKind
    OldValue string
    NewValue string
}

type PatchValue struct {
    Value string
    Delete bool
}

type BiMap struct {
    stringToInt map[string]int
    intToString map[int]string
}





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
	result := make([]int, 0, len(a))
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

	fmt.Println(cap(result))
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
	checkB := make(map[int]struct{})

	for _, v := range a {
		checkA[v] = struct{}{}
	}
	for _, v := range b {
		checkB[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := checkA[v]; !ok {
			result = append(result, v)
		}
	}
	for _, v := range a {
		if _, ok := checkB[v]; !ok {
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

// 21
func GroupByParity(nums []int) map[string][]int {
	if len(nums) == 0 {
		return nil
	}

	result := make(map[string][]int)
	for _, v := range nums {
		if v % 2 == 0 {
			result["even"] = append(result["even"], v)

		} else {
			result["odd"] = append(result["odd"], v)
		}
	}
	return result
}

// 22
func GroupByLength(words []string) map[int][]string {
	if len(words) == 0 {
		return nil
	}

	result := make(map[int][]string)
	for _, v := range words {
		lenght := utf8.RuneCountInString(v)
		result[lenght] = append(result[lenght], v)
	}
	return result
}

// 23
func GroupAnagrams(words []string) [][]string {
	if len(words) == 0 {
		return nil
	}

	result := [][]string{}
	groups := make(map[string][]string)
    keys := []string{}
    for _, v := range words {
        letters := []rune(v)
        sort.Slice(letters, func(a, b int) bool {
            return letters[a] < letters[b]
        })
        key := string(letters)
        if _, ok := groups[key]; !ok {
            keys = append(keys, key)
        }
        groups[key] = append(groups[key], v)
    }

    for _, key := range keys {
        result = append(result, groups[key])
    }
    return result
}

// 24
func MapValues(m map[string]int, f func(int) int) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		result[k] = f(v)
	}
	return result
}

// 25
func FilterMap(m map[string]int, keep func(string, int) bool) map[string]int {
	result := make(map[string]int)

    for key, value := range m {
        if keep(key, value) {
            result[key] = value
        }
    }
    return result
}

// 26
func Compose(first map[string]string, second map[string]int) map[string]int {
	result := make(map[string]int)

	for k, _ := range first {
		if _, ok := second[first[k]]; ok {
			result[k] = second[k]
		}
	}
	return result
}

// 27
func InnerJoin(left map[int]string, right map[int]int) []Row {
	result := []Row{}
    for id, name := range left {
        score, ok := right[id]

        if ok {
            result = append(result, Row{
                ID: id,
                Name: name,
                Score: score,
            })
        }
    }
    sort.Slice(result, func(i, j int) bool {
        return result[i].ID < result[j].ID
    })
    return result
}

// 28
func Diff(old, current map[string]string) []Change {
	result := []Change{}
    keys := make([]string, 0, len(old)+len(current))
    for key := range old {
        keys = append(keys, key)
    }
    for key := range current {
        if _, ok := old[key]; !ok {
            keys = append(keys, key)
        }
    }

    sort.Strings(keys)
	for _, key := range keys {
		oldValue, oldOK := old[key]
		newValue, newOK := current[key]
		switch {
		case !newOK:
			result = append(result, Change{
				Key: key,
				Kind: Deleted,
				OldValue: oldValue,
				NewValue: "",
			})
		case !oldOK:
			result = append(result, Change{
				Key: key,
				Kind: Added,
				OldValue: "",
				NewValue: newValue,
			})
		case oldValue != newValue:
			result = append(result, Change{
				Key: key,
				Kind: Changed,
				OldValue: oldValue,
				NewValue: newValue,
			})
		}
	}
    return result
}

// 29
func ApplyPatch(dst map[string]string, patch map[string]PatchValue) {
    for key, do := range patch {
        if do.Delete {
			delete(dst, key)
		} else {
			dst[key] = do.Value
		}
    }
}

// 30
func NewBiMap() *BiMap {
    return &BiMap{
        stringToInt: make(map[string]int),
        intToString: make(map[int]string),
    }
}

func (b *BiMap) Put(s string, n int) bool {
    if v, ok := b.stringToInt[s]; ok && v != n {
        return false
    }
    if v, ok := b.intToString[n]; ok && v != s {
        return false
    }

    b.stringToInt[s] = n
    b.intToString[n] = s
    return true
}

func (b *BiMap) GetByString(s string) (int, bool) {
    v, ok := b.stringToInt[s]
    return v, ok
}

func (b *BiMap) GetByInt(n int) (string, bool) {
    v, ok := b.intToString[n]
    return v, ok
}

func (b *BiMap) DeleteByString(s string) bool {
    n, ok := b.stringToInt[s]
    if !ok {
        return false
    }

    delete(b.stringToInt, s)
    delete(b.intToString, n)

    return true
}