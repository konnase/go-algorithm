package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func threeSumClosest(nums []int, target int) int {
	lenN := len(nums)
	if lenN < 3 {
		return 0
	}
	sort.Ints(nums)
	//fmt.Println(nums)
	s := 1<<31 - 1
	//result := math.MaxInt32
	for k := 0; k < lenN; k++ {
		//if nums[k] > 0 {
		//	break
		//}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i := k + 1
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
			if sum-target < 0 {
				i += 1
			} else if sum-target > 0 {
				j -= 1
			} else {
				return sum
			}
			if abs(sum-target) < abs(s-target) {
				s = sum
			}
		}
	}
	return s
}

func main() {
	//nums := []int{-1, 2, 1, -4}
	//t := 1
	//nums := []int{1, 1, 1, 1}
	//t := 0
	nums := []int{0, 2, 1, -3}
	t := 1
	fmt.Println(threeSumClosest(nums, t))
}
