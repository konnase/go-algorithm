package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	p := head.Next
	h := head
	head = p
	for ; p != nil; {
		temp := p.Next
		p.Next = h
		h.Next = temp
		if pre != nil {
			pre.Next = p
		}
		pre = h
		h = temp
		if temp == nil {
			break
		}
		p = temp.Next
	}
	return head
}

func main() {
	h := &ListNode{Val: 2, Next: nil}
	h = &ListNode{Val: 3, Next: h}
	h = &ListNode{Val: 4, Next: h}
	head := &ListNode{Val: 1, Next: h}
	head = swapPairs(head)
	for ; head != nil; {
		fmt.Println(head.Val)
		head = head.Next
	}
}
