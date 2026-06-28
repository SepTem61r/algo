package main

type listNode struct {
	Val  int
	Next *listNode
}

func newNode(val int) *listNode {
	return &listNode{
		Val:  val,
		Next: nil,
	}
}

func main() {
	n1 := newNode(1)
	n2 := newNode(2)
	n3 := newNode(3)
	n1.Next = n2
	n2.Next = n3
}

/* 在链表的节点 n0 之后插入节点 P */
func InsertNode(n0 *listNode, P *listNode) {
	n1 := n0.Next
	P.Next = n1
	n0.Next = P
}

/* 在链表中查找值为 target 的首个节点 */
func findNode(head *listNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}
