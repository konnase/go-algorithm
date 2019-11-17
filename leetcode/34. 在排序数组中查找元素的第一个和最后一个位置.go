package main

import "fmt"

func searchResult(nums []int, l, r, target int) []int {
	if l > r {
		return []int{-1, -1}
	}
	mid := (l + r) / 2
	begin, end := -1, -1
	if nums[mid] == target {
		if mid-1 >= l && nums[mid-1] == target {
			t := searchResult(nums, l, mid-1, target)
			begin = t[0]
		} else {
			begin = mid
		}
		if mid+1 <= r && nums[mid+1] == target {
			t := searchResult(nums, mid+1, r, target)
			end = t[1]
		} else {
			end = mid
		}
	} else if nums[mid] > target {
		t := searchResult(nums, l, mid-1, target)
		begin = t[0]
		end = t[1]
	} else {
		t := searchResult(nums, mid+1, r, target)
		begin = t[0]
		end = t[1]
	}
	return []int{begin, end}
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return searchResult(nums, 0, len(nums)-1, target)
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	fmt.Println(searchRange(nums, target))
}
