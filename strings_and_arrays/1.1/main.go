package main

import (
	"sort"
	"strings"
)

func isUniqueWithHashMap(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	uniMap := make(map[rune]struct{})
	for _, r := range strings.ToLower(s) {
		if _, ok := uniMap[r]; !ok {
			uniMap[r] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func isUniqueWithOutHashMap(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	sl := strings.Split(strings.ToLower(s), "")
	sort.Strings(sl)
	for i := range sl {
		f := i + 1
		if f < len(sl) {
			if sl[i] == sl[f] {
				return false
			}
		}
	}

	return true
}

func main() {

}
