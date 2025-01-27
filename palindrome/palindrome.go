package main

import (
	"fmt"
)

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func reverseList(head *EduLinkedListNode) *EduLinkedListNode {
	var prev, next, cur *EduLinkedListNode = nil, nil, head

	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	return prev
}

func palindrome(head *EduLinkedListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = slow.next.next
	}

	cur := reverseList(slow)

	for head != nil && cur != nil {
		if head.data != cur.data {
			return false
		}
		cur = cur.next
		head = head.next
	}

	return true
}

func main() {

	node := EduLinkedListNode{data: 2}
	node1 := EduLinkedListNode{data: 4}
	node2 := EduLinkedListNode{data: 6}
	node3 := EduLinkedListNode{data: 4}
	node4 := EduLinkedListNode{data: 2}

	node.next = &node1
	node1.next = &node2
	node2.next = &node3
	node3.next = &node4

	fmt.Println(palindrome(&node))
}
