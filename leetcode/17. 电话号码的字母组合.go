package main

import "fmt"

//var phone []string
//var indice []int

func combine(indice []int, phone []string, index int) []string {
	if index >= len(indice) {
		return []string{}
	}
	num := indice[index]
	former := combine(indice, phone, index + 1)
	var result []string
	if len(former) == 0 {
		for _, c := range phone[num] {
			result = append(result, string(c))
		}
	} else {
		for _, c := range phone[num] {
			for _, str := range former {
				temp := string(c) + str
				result = append(result, temp)
			}
		}
	}
	return result
}

func letterCombinations(digits string) []string {
	var phone []string
	var indice []int
	phone = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if len(digits) == 0 {
		return []string{}
	}
	for _, c := range digits {
		index := c - '0'
		indice = append(indice, int(index))
	}
	//fmt.Println(indice)
	return combine(indice, phone, 0)
}

func main() {
	nums := "2"
	fmt.Println(letterCombinations(nums))
}
