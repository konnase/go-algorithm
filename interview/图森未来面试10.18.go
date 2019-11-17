package main

/*
将k个升序数组合并成一个降序数组
*/

import (
	"fmt"
	"math"
)

func sortWithCompare(a [][]int) []int {
	var pointer []int
	lenA := 0
	i := 0
	for i = 0; i < len(a); i++ {
		pointer = append(pointer, len(a[i])-1)
		lenA += len(a[i])
	}
	var result []int
	for {
		index := -1
		maxV := math.MinInt32
		for i := 0; i < len(pointer); i++ {
			if pointer[i] >= 0 && a[i][pointer[i]] > maxV {
				maxV = a[i][pointer[i]]
				index = i
			}
		}
		if index == -1 {
			break
		}
		pointer[index] -= 1
		result = append(result, maxV)
	}
	return result
}

func main() {
	a := [][]int{
		{1, 4, 5, 7, 12},
		{1, 2, 6, 7, 11, 20},
		{2, 3, 3, 3},
	}
	fmt.Println(sortWithCompare(a))
}
