package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func maxArea(height []int) int {
	lenH := len(height)
	if lenH == 2 {
		return min(height[0], height[1])
	}
	i := 0
	j := lenH-1
	maxS := min(height[0], height[lenH-1]) * (lenH - 1)
	for {
		if i >= j {
			break
		}
		temp := min(height[i], height[j]) * (j-i)
		if temp > maxS {
			maxS = temp
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}
	return maxS
}

func main() {
	a := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(a))
}
