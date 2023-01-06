package main

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	var ans string
	ans = strs[0]
	for _, item := range strs {
		for i := 0; i < len(item) && i < len(ans); i++ {
			if ans[i] != item[i] {
				ans = item[0:i]
				break
			}
		}
	}
	return ans
}
