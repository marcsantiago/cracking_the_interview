package main

import (
	"sort"
	"strings"
)

func isPermutation(a, b string) bool {
	a = strings.ToLower(strings.TrimSpace(a))
	b = strings.ToLower(strings.TrimSpace(b))
	if len(a) != len(b) {
		return false
	}

	sa := strings.Split(a, "")
	sb := strings.Split(b, "")
	sort.Strings(sa)
	sort.Strings(sb)

	for i := range sa {
		if sa[i] != sb[i] {
			return false
		}
	}

	return true
}

func main() {

}

// in thoery I could build out a sorting type for strings
// type sortRunes []rune

// func (s sortRunes) Less(i, j int) bool {
//     return s[i] < s[j]
// }

// func (s sortRunes) Swap(i, j int) {
//     s[i], s[j] = s[j], s[i]
// }

// func (s sortRunes) Len() int {
//     return len(s)
// }

// func SortString(s string) string {
//     r := []rune(s)
//     sort.Sort(sortRunes(r))
//     return string(r)
// }
