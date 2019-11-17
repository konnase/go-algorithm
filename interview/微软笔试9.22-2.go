package main

import (
	"fmt"
	"math"
)

var MAXINT = math.MaxInt32
func isPrime(x int) bool {
	if x == 2 || x == 3 || x == 5 {
		return true
	}
	for i := 2; i < x/2; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func kPrime(k int) []int {
	if k == 1 {
		return []int{2}
	}
	if k == 2 {
		return []int{2, 3}
	}
	if k == 3 {
		return []int{2, 3, 5}
	}
	primes := []int{2, 3, 5}
	k -= 3
	i := 6
	for {
		if isPrime(i) {
			primes = append(primes, i)
			k--
		}
		if k == 0 {
			break
		}
		i++
	}
	return primes
}

func dp(d int, primes []int) int {
	if d == 1 {
		return -1
	}
	if d == 2 || d == 3 || d == 5 {
		return 1
	}
	if d == 4 {
		return 2
	}
	a := make([]int, d+1)
	a[0] = 0
	a[1] = MAXINT
	a[2] = 1
	a[3] = 1
	a[4] = 2
	a[5] = 1
	for i := 6; i <= d; i++ {
		minD := a[i-2] + 1
		for j := 1; j < len(primes); j++ {
			if i < primes[j] {
				continue
			}
			if a[i - primes[j]] + 1 < minD{
				minD = a[i - primes[j]] + 1
			}
		}
		a[i] = minD
	}
	return a[d]
}

func minDays(input1 int, input2 int) int {
	d := input1
	k := input2
	if d <= 1 {
		return -1
	}
	if k <= 0 {
		return -1
	}
	primes := kPrime(k)
	fmt.Println(primes)
	return dp(d, primes)
}

func main() {
	//fmt.Println(minDays(10, 1))
	//fmt.Println(minDays(11, 3))
	var d, k int
	fmt.Scanf("%d", &d)
	fmt.Scanf("%d", &k)
	fmt.Println(minDays(d, k))
}
