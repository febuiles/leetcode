package main

import "fmt"

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func removeNthLastNode(head *EduLinkedListNode, n int) *EduLinkedListNode {
	left := head
	right := head
	for i := 0; i < n; i++ {
		right = right.next
	}

	if right == nil {
		return head.next
	}

	for right.next != nil {
		right = right.next
		left = left.next
	}
	left.next = left.next.next
	return head
}

func main() {

	node := EduLinkedListNode{data: 23}
	node1 := EduLinkedListNode{data: 89}
	node2 := EduLinkedListNode{data: 10}
	node3 := EduLinkedListNode{data: 5}
	node4 := EduLinkedListNode{data: 67}
	node5 := EduLinkedListNode{data: 39}
	node6 := EduLinkedListNode{data: 70}
	node7 := EduLinkedListNode{data: 28}

	node.next = &node1
	node1.next = &node2
	node2.next = &node3
	node3.next = &node4
	node4.next = &node5
	node5.next = &node6
	node6.next = &node7
	node7.next = nil

	removeNthLastNode(&node, 2)
	fmt.Print("> ")
	cur_node := node
	for true {
		fmt.Printf("%d ", cur_node.data)
		if cur_node.next == nil {
			break
		}
		cur_node = *cur_node.next
	}
	fmt.Print("|")
}
