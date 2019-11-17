package main

import "fmt"

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}
	i, j := -1, -1
	for k := len(nums) - 1; k >= 1; k-- {
		if nums[k] > nums[k-1] {
			i = k - 1
			break
		}
	}
	if i == -1 {
		m, n := 0, len(nums)-1
		for ; ; {
			if m >= n {
				return
			}
			swap(nums, m, n)
			m++
			n--
		}
	}
	for k := len(nums) - 1; k > i; k--{
		if nums[k] > nums[i] {
			j = k
			break
		}
	}
	swap(nums, i, j)
	m, n := i+1, len(nums)-1
	for ; ; {
		if m >= n {
			return
		}
		swap(nums, m, n)
		m++
		n--
	}
}

func main() {
	nums := []int{1,2,3}
	nextPermutation(nums)
	fmt.Println(nums)
}
