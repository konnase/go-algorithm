package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}
	nilNode := &ListNode{Val:-1, Next:head}
	var prefix = nilNode
	slow := nilNode
	fast := head
	count := 0
	for ; fast != nil; {
		count += 1
		if count >= n {
			prefix = slow
			slow = slow.Next
		}
		fast = fast.Next
	}
	if slow == head {
		head = head.Next
	} else {
		prefix.Next = slow.Next
	}
	slow = nil
	nilNode = nil
	prefix = nil
	return head
}

func main() {
	//h := &ListNode{Val:5, Next:nil}
	//h = &ListNode{Val:4, Next:h}
	//h = &ListNode{Val:3, Next:h}
	h := &ListNode{Val: 2, Next: nil}
	head := &ListNode{Val: 1, Next: h}
	head = removeNthFromEnd(head, 1)
	for ; head != nil; {
		fmt.Println(head.Val)
		head = head.Next
	}
}
