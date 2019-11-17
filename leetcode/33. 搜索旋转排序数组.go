package main

import "fmt"

func searchTarget(nums []int, l, r, target int) int {
	if l > r {
		return -1
	}
	mid := (l+r)/2
	if nums[mid] > target {
		if nums[mid] >= nums[l] {
			if nums[l] > target {
				return searchTarget(nums, mid+1, r, target)
			} else {
				return searchTarget(nums, l, mid-1, target)
			}
		} else {
			return searchTarget(nums, l, mid-1, target)
		}
	} else if nums[mid] == target {
		return mid
	} else {
		if nums[mid] < nums[l] {
			if nums[l] > target {
				return searchTarget(nums, mid + 1, r, target)
			} else {
				return searchTarget(nums, l , mid-1, target)
			}
		} else {
			return searchTarget(nums, mid+1, r, target)
		}
	}
}

func search(nums []int, target int) int {
	L := len(nums)
	if L == 0 {
		return -1
	}
	return searchTarget(nums, 0, L-1, target)
}

func main() {
	nums := []int {4,5,6,7,0,1,2}
	target := 0
	fmt.Println(search(nums, target))
}
