package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	lenN := len(nums)
	if lenN < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	//fmt.Println(nums)
	result := [][]int{}
	for k := 0; k < lenN; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i := k+1
		j := lenN - 1
		for {
			if i >= j {
				break
			}
			if i > k+1 && nums[i] == nums[i-1] {
				i += 1
				continue
			}
			if j < lenN-1 && nums[j] == nums[j+1] {
				j -= 1
				continue
			}
			sum := nums[k] + nums[i] + nums[j]
			if sum < 0 {
				i += 1
			} else if sum > 0 {
				j -= 1
			} else {
				temp := []int{nums[k], nums[i], nums[j]}
				result = append(result, temp)
				i += 1
				j -= 1
			}
		}
	}
	return result
}

func main()  {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{0,0,0}
	//nums := []int{-2,0,0,2,2}
	nums := []int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0}
	fmt.Println(threeSum(nums))
}
