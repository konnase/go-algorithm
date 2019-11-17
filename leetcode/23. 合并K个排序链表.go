package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (l ListNodeHeap)Len() int {
	return len(l)
}

func (l ListNodeHeap) Less(i, j int) bool {
	return l[i].Val < l[j].Val
}

func (l ListNodeHeap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *ListNodeHeap) Push(x interface{}) {
	item := x.(*ListNode)
	*l = append(*l, item)
}

func (l *ListNodeHeap) Pop() interface{} {
	n := len(*l)
	item := (*l)[n-1]
	*l = (*l)[:n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var lst ListNodeHeap
	isNil := true
	for _, l := range lists{
		if l != nil {
			isNil = false
			lst = append(lst, l)
		}
	}
	if isNil {
		return nil
	}
	var result, current *ListNode
	heap.Init(&lst)
	for ; lst.Len() > 0; {
		node := heap.Pop(&lst)
		if node.(*ListNode).Next != nil {
			heap.Push(&lst, node.(*ListNode).Next)
		}
		if result == nil {
			result = node.(*ListNode)
			current = result
		} else {
			current.Next = node.(*ListNode)
			current = current.Next
		}
	}
	return result
}

func main() {
	var lists []*ListNode
	var h *ListNode
	h = &ListNode{Val:5, Next:nil}
	h = &ListNode{Val:4, Next:h}
	h = &ListNode{Val:1, Next:h}
	lists = append(lists, h)
	h = &ListNode{Val:4, Next:nil}
	h = &ListNode{Val:3, Next:h}
	h = &ListNode{Val:1, Next:h}
	lists = append(lists, h)
	h = &ListNode{Val:6, Next:nil}
	h = &ListNode{Val:2, Next:h}
	lists = append(lists, h)
	result := mergeKLists(lists)
	for ; result != nil; {
		fmt.Println(result.Val)
		result = result.Next
	}
}