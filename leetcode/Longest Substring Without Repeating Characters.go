package main

import (
	"fmt"
)

// O(n^2)
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	a := make([]int, len(s))
	a[0] = 1
	for i := 1; i < len(s); i++ {
		changed := false
		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				if i-j < a[i-1]+1 {
					a[i] = i - j
				} else {
					a[i] = a[i-1] + 1
				}
				changed = true
			}
		}
		if !changed {
			a[i] = a[i-1] + 1
		}

	}
	maxI := 0
	fmt.Println(a)
	for _, value := range a {
		if value > maxI {
			maxI = value
		}
	}
	return maxI
}

func lengthOfLongestSubstringOn(s string) int {
	a := make([]int, 256)
	for i := range a {
		a[i] = -1
	}
	maxLen := 0
	start := -1
	for i := range s {
		if a[int(s[i])] > start {
			start = a[int(s[i])]
		}
		a[int(s[i])] = i
		maxLen = max(maxLen, i-start)
	}
	return maxLen
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	//s := "abba"
	s := "abbcabcbb"
	//fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstringOn(s))
}
