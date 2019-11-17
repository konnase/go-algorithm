package main

import (
	"fmt"
	"sort"
)

/*
4
3 8
2 10
5 2
*/

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var emp []int
	empNum := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		var num, value int
		fmt.Scanf("%d %d", &num, &value)
		emp = append(emp, value)
		empNum[value] = num
	}
	sort.Ints(emp)
	//fmt.Println(emp)
	//fmt.Println(empNum)
	//fmt.Println(empNum[2])
	i := 0
	j := len(emp) - 1
	max := 0
	for {
		if i == j && empNum[emp[i]] == 0 {
			break
		}
		if empNum[emp[i]] != 0 {
			if empNum[emp[j]] != 0 {
				temp := emp[i] + emp[j]
				if temp > max {
					max = temp
				}
				empNum[emp[i]] -= 1
				empNum[emp[j]] -= 1
			} else {
				j -= 1
			}
		} else {
			i += 1
		}
	}
	fmt.Println(max)
}
