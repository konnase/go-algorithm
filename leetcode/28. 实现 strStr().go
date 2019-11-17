package main

import (
	"fmt"
	"strings"
)

// violent algorithm
func strStr1(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	lenH := len(haystack)
	for i := 0; i < lenH; i++ {
		if haystack[i] != needle[0] {
			continue
		}
		j := i + 1
		k := 1
		matched := true
		for {
			if k >= len(needle) {
				break
			}
			if j >= lenH || haystack[j] != needle[k] {
				matched = false
				break
			}
			j++
			k++
		}
		if matched {
			return i
		}
	}
	return -1
}

// call library
func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	index := strings.Index(haystack, needle)
	return index
}

// sunday algorithm
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	if len(haystack) < len(needle) {
		return -1
	}
	a := make([]int, 256)
	count := 1
	for i := len(needle) - 1; i >= 0; i-- {
		c := needle[i]
		if a[c] == 0 {
			a[c] = count
			count++
		}
	}
	lenH := len(haystack)
	idx := 0
	for {
		if idx >= lenH || idx+len(needle) > lenH{
			return -1
		}
		match := true
		for j := 0; j < len(needle); j++ {
			if haystack[idx+j] == needle[j] {
				continue
			}
			match = false
			break
		}
		if match {
			return idx
		}
		if idx+len(needle) >= lenH {
			return -1
		}
		c := haystack[idx+len(needle)]
		in := false
		for j := 0; j < len(needle); j++ {
			if c == needle[j] {
				idx += a[c]
				in = true
				break
			}
		}
		if !in {
			idx += len(needle)+1
		}
	}
}

func main() {
	haystack := "mississipi"
	needle := "issi"
	fmt.Println(strStr(haystack, needle))
}
