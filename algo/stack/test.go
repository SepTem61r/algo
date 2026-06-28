package main

import "container/list"

type MyQueue struct {
	data *list.List
}

func Constructor() MyQueue {
	return MyQueue{data: list.New()}
}

func (this *MyQueue) Push(x int) {
	this.data.PushBack(x)
}

func (this *MyQueue) Pop() any {
	e := this.data.Front()
	this.data.Remove(e)
	return e.Value
}

func (this *MyQueue) Peek() any {
	e := this.data.Front()
	return e.Value
}

func (this *MyQueue) Empty() bool {
	return this.data.Len() == 0
}
