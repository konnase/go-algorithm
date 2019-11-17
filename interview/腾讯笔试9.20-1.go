package main

import (
	"fmt"
	"strconv"
)

func main() {
	var t, n int
	needChange := false
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var a []int
		if needChange {
			var temp byte
			fmt.Scanf("%c", &temp)
		}
		fmt.Scanf("%d", &n)
		//fmt.Printf("n=%d\n", n)
		for {
			var temp byte
			//for {
			fmt.Scanf("%c", &temp)
			if temp != '\n' {
				value, _ := strconv.Atoi(string(temp))
				a = append(a, value)
			} else {
				needChange = true
				break
			}
			//}
		}
		//fmt.Println(a)
		lenA := len(a)
		if lenA < 11 {
			fmt.Println("NO")
			continue
		}

		yes := false
		for j := lenA - 1; j >= 0; j-- {
			if a[j] == 8 && lenA-j >= 11 {
				//fmt.Printf("j= %d\n", j)
				fmt.Println("YES")
				yes = true
				break
			}
		}
		if !yes {
			fmt.Println("NO")
		}
	}
}

/*
4
15
138842311245865
3
000
12
132141245865
13
1381412458654
 */