package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x >= 0 && x <= 9{
		return true
	}
	if x < 0 && x >= -9 {
		return false
	}
	str := strconv.FormatInt(int64(x), 10)
	i, j := 0, len(str)-1
	for {
		if i >= j {
			break
		}
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	a := 12241
	fmt.Println(isPalindrome(a))
}
