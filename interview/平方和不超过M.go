// 给 N 个整数，和一个 M（不超过 10^9），每次可以将一个整数 s 拆分成两个整数 a、b，但是必须满足 s = a + b。问最少拆分多少次之后，可以满足所有数的平方和不超过 M。（来自幻方面试题）

/*
解题思路
首先，假如我们确定了每个数最终会被拆分成多少个数，那么肯定是对其进行均匀拆分，才能使拆分结果的平方和最小。
在这个基础上，有一种贪心的思路是：记录每一个整数以及这个整数拆分当前拆分次数，计算出这个整数以当前拆分次数 c 和 c+1 分别进行拆分之后，对平方和的影响是多少；然后每次选择影响最大的一个（用堆来维护）将其拆分次数 +1，一直操作到满足 M 的限制为止。
这种贪心的正确性大概是可以得到证明的（没有仔细推导过，但是跑了一些数据看起来是对的），但是复杂度还是太高了，因为每次只进行一次拆分，而答案的上限可以达到 M-1。
对于上面贪心的过程，其实有很多操作是重复的，首先对于一开始就是相同的数字的部分，可以先并在一起；其次，对于一个数字，中间很多次 c+1 的操作，对平方和的影响是相同的，例如一个数字 100，将其拆分成 25 份的话，拆分方案为 [4, 4, 4, ..., 4]，拆分成 33 份的话，方案为 [4, 3, 3, 3, ..., 3]，在 25 份到 33 份之间，每多拆分一份，都是将 3 个 4 变成 4 个 3，对于整体平方和的影响都是 -12，如果我们将这样的区间统一处理的话，最终每个数字 s 需要进行处理的拆分方案数量级就是 sqrt(s)（这样的区间的数量不会超过 2sqrt(s)）。
通过这样的合并，就可以在限定时间内完成贪心的过程了。
*/
package main

import (
	"container/heap"
	"fmt"
)

var (
	N = 10
	M = 100
)

// 最大堆
type PriorityQueue []int

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p *PriorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue) Push(a any) {
	(*p) = append((*p), a.(int))
}

func (p *PriorityQueue) Pop() any {
	old := *p
	n := len(old)
	a := old[n-1]
	*p = old[:n-1]
	return a
}

func catNum(n int) (int, int) {
	half := n / 2
	if n%2 == 0 {
		return half, half
	}
	return half + 1, half
}

func main() {
	a := []int{5, 7, 4, 6}
	p := &PriorityQueue{}
	heap.Init(p)

	for _, v := range a {
		heap.Push(p, v)
	}

	squareSum := 0
	for _, v := range a {
		squareSum += v * v
	}

	println(squareSum)

	count := 0
	for squareSum > M {
		top := heap.Pop(p).(int)
		fmt.Printf("top: %d\n", top)
		a, b := catNum(top)
		delta := top*top - a*a - b*b

		heap.Push(p, a)
		heap.Push(p, b)
		squareSum -= delta
		count++
	}
	println(count)
}
