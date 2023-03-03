package main

import "fmt"

/*
将给出的链表中的节点每 k 个一组翻转，返回翻转后的链表
如果链表中的节点数不是 k 的倍数，将最后剩下的节点保持原样
你不能更改节点中的值，只能更改节点本身。

数据范围： $$\ 0 \le n \le 2000$$ ， $$1 \le k \le 2000$$ ，链表中每个元素都满足 $$0 \le val \le 1000$$
要求空间复杂度 $$O(1)$$，时间复杂度 $$O(n)$$

例如：

给定的链表是 $$1\to2\to3\to4\to5$$

对于 $$k = 2$$ , 你应该返回 $$2\to 1\to 4\to 3\to 5$$

对于 $$k = 3$$ , 你应该返回 $$3\to2 \to1 \
*/

type Node struct {
	val  int
	next *Node
}

func reverse(root *Node, k int) *Node {
	temp := root
	sum := 0
	for {
		if temp == nil {
			break
		}
		temp = temp.next
		sum++
	}
	cur := root
	temproot := root
	preroot := root
	var pre, result *Node
	for i := 0; i < sum/k; i++ {
		for j := 0; j < k; j++ {
			temp := cur.next
			cur.next = pre
			pre = cur
			cur = temp
		}
		if temproot == root {
			result = pre
		} else {
			preroot.next = pre
			preroot = temproot
		}
		temproot.next = cur
		pre = temproot
		temproot = cur
	}
	return result
}

func main() {
	root := &Node{
		val:  0,
		next: &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, &Node{8, &Node{9, &Node{10, nil}}}}}}}}}},
	}
	result := reverse(root, 3)
	for {
		if result == nil {
			break
		}
		fmt.Println(result.val)
		result = result.next
	}
}
