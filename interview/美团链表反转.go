package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	id int
	next *Node
}

func reverse(root *Node, k int) *Node {
	count := 0
	if root == nil {
		return nil
	}
	if root.next ==nil {
		return root
	}
	var result *Node
	head := root

	var secondHead *Node
	next := root.next
	for ;head.next != nil; {
		count += 1
		if count == k {
		if next.next == nil {
		head.next = nil
		break
	}
		if secondHead == nil {
		result = next
	}
		secondHead = head
		head = next.next
		count = 0
		secondHead.next = head
		continue
	}
		temp := next.next
		next.next = head
		if secondHead != nil {
		secondHead.next = next
	}
		if next.next == nil {
		head.next = nil
		break
	}
		next = temp

	}
	return result
}

func readLine1() string {
	var temp byte
	var result []byte
	fmt.Scanf("%c", &temp)
	for ;temp != '\n'; {
		result = append(result, temp)
		fmt.Scanf("%c", &temp)
	}
	return string(result)
}

func main() {
	var k int
	linkListStr := readLine1()
	fmt.Scanf("%d", &k)
	// linkListList := make([]int, 0)
	linkListList := strings.Split(linkListStr, " ")
	root := &Node{}
	parent := &Node{}
	for _, value := range linkListList {
		intValue, _ := strconv.Atoi(value)
		item := &Node{
			id: intValue,
			next : nil,
		}
		if root == nil {
			root = item
			parent = item
		} else {
			parent.next = item
			parent = item
		}

	}
	result := reverse(root, k)
	for ;result.next != nil; {
		fmt.Println(result.id)
		result = result.next
	}
}