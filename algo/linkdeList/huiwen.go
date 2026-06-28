package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	curr := dummy

	for head != nil {
		curr.Next = head
		head = head.Next
		fmt.Printf("%d\n", head.Val)
		curr = curr.Next
	}
	return dummy.Next
}
