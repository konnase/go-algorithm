package main

import "fmt"

func romanToInt(s string) int {
	var result int
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		if s[i] == 'I' {
			if i+1 < lenS {
				if s[i+1] == 'V' {
					result += 4
					i += 1
				} else if s[i+1] == 'X' {
					result += 9
					i += 1
				} else {
					result += 1
				}
			} else {
				result += 1
			}
		} else if s[i] == 'X' {
			if i+1 < lenS {
				if s[i+1] == 'L' {
					result += 40
					i += 1
				} else if s[i+1] == 'C' {
					result += 90
					i += 1
				} else {
					result += 10
				}
			} else {
				result += 10
			}
		} else if s[i] == 'C' {
			if i+1 < lenS {
				if s[i+1] == 'D' {
					result += 400
					i += 1
				} else if s[i+1] == 'M' {
					result += 900
					i += 1
				} else {
					result += 100
				}
			} else {
				result += 100
			}
		} else if s[i] == 'V' {
			result += 5
		} else if s[i] == 'L' {
			result += 50
		} else if s[i] == 'D' {
			result += 500
		} else if s[i] == 'M' {
			result += 1000
		}
	}
	return result
}

func main() {
	//s := "MCMXCIV"
	s := "LVIII"
	fmt.Println(romanToInt(s))
}
