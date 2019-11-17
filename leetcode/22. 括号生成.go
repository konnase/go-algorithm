package main

import "fmt"

var result []string

func dfs(l, r, n int, str string) {
	if l < r {
		return
	}
	if l == n && r == n {
		result = append(result, str)
		return
	}
	if l < n {
		dfs(l+1, r, n, str + "(")
	}
	if r < n {
		dfs(l, r+1, n, str + ")")
	}
}

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	result = []string{}
	dfs(0, 0, n, "")
	return result
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}


