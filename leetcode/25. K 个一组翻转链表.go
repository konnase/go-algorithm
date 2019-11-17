package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1{
		return head
	}
	h := head
	var pre *ListNode
	p := h.Next
	count := 1
	for ; p != nil; {
		count += 1
		if count == k {
			h1 := h
			nextK := p.Next
			n := h.Next
			for ; n != p; {
				temp := n.Next
				n.Next = h
				h = n
				n = temp
			}
			n.Next = h
			if pre != nil {
				pre.Next = p
			} else {
				head = p
			}
			pre = h1
			pre.Next = nextK
			h = nextK
			if h == nil {
				break
			}
			p = h.Next
			count = 1
			continue
		} else {
			p = p.Next
		}
	}
	return head
}

func main() {
	h := &ListNode{Val: 5, Next: nil}
	h = &ListNode{Val: 4, Next: h}
	h = &ListNode{Val: 3, Next: h}
	h = &ListNode{Val: 2, Next: h}
	head := &ListNode{Val: 1, Next: h}
	head = reverseKGroup(head, 2)
	for ; head != nil; {
		fmt.Println(head.Val)
		head = head.Next
	}
}