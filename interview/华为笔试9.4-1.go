package main

/*
商店有4盒装的钉子和9盒装的钉子，肖明想要买n颗钉子，请问最少需要买几盒。不能组成n颗钉子的情况输出-1
*/

import (
	"fmt"
)

func first() {
	var n int
	fmt.Scanln(&n)
	if n == 4 || n == 9 {
		fmt.Println(1)
		return
	}
	if n < 9 {
		fmt.Println(-1)
	}
	s := make([]int, n+1)
	for i := 0; i <= 9; i++ {
		s[i] = -1
	}
	s[4] = 1
	s[8] = 2
	s[9] = 1
	for i := 10; i <= n; i++ {
		if s[i-4] == -1 && s[i-9] == -1 {
			s[i] = -1
		} else if s[i-4] != -1 && s[i-9] != -1 {
			if s[i-4] > s[i-9] {
				s[i] = s[i-9] + 1
			} else {
				s[i] = s[i-4] + 1
			}
		} else if s[i-4] != -1 {
			s[i] = s[i-4] + 1
		} else if s[i-9] != -1 {
			s[i] = s[i-9] + 1
		}
	}
	//for i := 0; i <= n; i++{
	//	fmt.Printf("%d %d\n",i, s[i])
	//}
	fmt.Println(s[n])
}

func main() {
	first()
}
