package main

import (
	"fmt"
	"math"
)

func longestValidParentheses(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	l := len(s)
	d := make([]int, l)
	d[0] = 0
	if s[0] == '(' && s[1] == ')' {
		d[1] = 2
	}
	for i := 2; i < l; i++ {
		if s[i] == ')' && s[i-1] == '(' {
			d[i] = d[i-2] +2
		} else if s[i] == ')' && s[i-1] == ')' {
			index := i-d[i-1]-1
			if index < 0 {
				continue
			}
			if s[index] == '(' {
				if index-1 < 0 {
					d[i] = d[i-1] + 2
				} else {
					d[i] = d[index-1] + d[i-1] + 2
				}
			}
		}
	}
	M := math.MinInt32
	for _, v := range d {
		if v > M {
			M = v
		}
	}
	return M
}

func main() {
	//s := ")()())"
	s := "(()))()()()()("
	fmt.Println(longestValidParentheses(s))
}
