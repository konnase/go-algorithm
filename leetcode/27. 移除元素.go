package main

func removeElement(nums []int, val int) int {
	lenN := len(nums)
	if lenN == 0 {
		return 0
	} else if lenN == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}
	i, j := 0, 0
	for {
		if j > lenN-1 {
			break
		}
		if nums[j] != val {
			nums[i] = nums[j]
			i++
			j++
		} else {
			j++
		}
	}
	return i
}
