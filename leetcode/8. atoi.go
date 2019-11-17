package main

import (
	"fmt"
)

func myAtoi(str string) int {
	result := 0
	signed := 1
	//base := 1
	start := false
	for _, c := range str {
		if !start {
			if c == ' ' {
				continue
			}
			if c == '-' {
				signed = -1
				start = true
			} else if c >= '0' && c <= '9' {
				result = result * 10 + int(c - '0')
				//base *= 10
				start = true
			} else if c == '+' {
				start = true
			}else {
				return 0
			}
		} else {
			if c >= '0' && c <= '9' {
				result = result * 10 + int(c - '0')
				//base *= 10
			} else {
				return result * signed
			}
			if result * signed > (1 << 31) - 1{
				return (1 << 31) - 1
			} else if result * signed < -(1 << 31) {
				return -(1 << 31)
			}
		}
	}
	return result * signed
}

func main()  {
	a := myAtoi("2344")
	fmt.Println(a)
}