package main

import (
	"container/list"
	"fmt"
)

/* 初始化双向队列 */
// 在 Go 中，将 list 作为双向队列使用

func main() {
	deque := list.New()

	/* 元素入队 */
	deque.PushBack(2) // 添加至队尾
	deque.PushBack(5)
	deque.PushBack(4)
	deque.PushFront(3) // 添加至队首
	deque.PushFront(1)

	/* 访问元素 */
	front := deque.Front() // 队首元素
	rear := deque.Back()   // 队尾元素

	/* 元素出队 */
	deque.Remove(front) // 队首元素出队
	deque.Remove(rear)  // 队尾元素出队

	/* 获取双向队列的长度 */
	size := deque.Len()
	fmt.Printf("%d\n", size)
	/* 判断双向队列是否为空 */
	isEmpty := deque.Len() == 0
	fmt.Printf("%d\n", isEmpty)
}
