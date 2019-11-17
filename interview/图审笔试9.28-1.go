package main

import "fmt"

func main() {
	MAXI := int64(100003)
	base := int64(2014)
	var y, k int64
	fmt.Scanf("%d %d", &y, &k)
	if y == 2014 {
		fmt.Println(0)
	} else if y == 2015 {
		fmt.Println(1)
	} else {
		n := y - base
		//a := make([]int64, n+1)
		//a[0] = 0
		//a[1] = 1
		var a_i_2, a_i_1, a int64;
		a_i_2 = 0
		a_i_1 = 1
		for i := int64(2); i <= n; i++ {
			//a[i] = a[i-2] * a[i-2] + k * a[i-1]
			a = a_i_2 * a_i_2 + k * a_i_1
			if a > MAXI {
				a %= MAXI
			}
			a_i_2 = a_i_1
			a_i_1 = a
		}
		//fmt.Println(a[n])
		fmt.Println(a)
	}

}