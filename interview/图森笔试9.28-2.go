package main

import (
	"fmt"
)

//var MAXY = int(10e7)

func robotTravel(n int) int {
	listY := make(map[int][]int, 0)
	for i := 0; i < n; i++ {
		//fmt.Println(i)
		var x, y int
		fmt.Scanf("%d", &x)
		fmt.Scanf("%d", &y)
		listY[y] = append(listY[y], x)
	}
	//fmt.Println("start traveling")
	sum := 0
	count := 0 // traveled dots
	x := 1
	y := 1
	for {
		if count >= n {
			break
		}
		//fmt.Printf("count = %d, y = %d", count, y)
		xHeap := listY[y]
		if len(xHeap) == 0 {
			y++
			sum += 1
			continue
		}
		minX := xHeap[0]
		maxX := minX
		for j := 0; j < len(xHeap); j++ {
			count++
			temp := xHeap[j]
			//fmt.Printf("count = %d, y = %d temp = %d\n", count, y, temp)
			if temp > maxX {
				maxX = temp
			}
			if temp < minX {
				minX = temp
			}
		}
		if x == minX && x == maxX {
			x = minX
		} else if x <= minX {
			sum += (maxX - x)
			x = maxX
		} else if x >= maxX {
			sum += (x - minX)
			x = minX
		} else {
			if x-minX < maxX-x {
				sum += (x - minX)
				x = maxX
			} else {
				sum += (maxX - x)
				x = minX
			}
			sum += (maxX - minX)
		}
		//fmt.Printf("sum = %d\n", sum)
		y++
		sum += 1
	}
	return sum - 1
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	//fmt.Println(n)
	sum := robotTravel(n)
	fmt.Println(sum)
}

/*
2
3 3
2 2

7
2 2
4 2
3 3
1 4
4 4
3 5
4 5
*/
