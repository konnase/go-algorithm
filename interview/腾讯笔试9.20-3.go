package main

import (
	"fmt"
	"sort"
)

/*
7 5
5 8 10 3 6 10 8
*/

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	var a []int
	for i := 0; i < n; i++ {
		var value int
		fmt.Scanf("%d", &value)
		a = append(a, value)
	}
	sort.Ints(a)
	//fmt.Println(a)
	sum := 0
	i := 0
	for {
		if i == len(a) {
			break
		}
		//fmt.Printf("print: %d %d\n",a[i], sum)
		if a[i]-sum > 0 {
			fmt.Println(a[i] - sum)
			sum = a[i]
			i++
		} else {
			i++
		}
	}
}
