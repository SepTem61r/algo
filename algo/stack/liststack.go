package main

import (
	"container/list"
)

type linkedListStack struct {
	data *list.List
}

func NewLinkerListStack() *linkedListStack {
	return &linkedListStack{data: list.New()}
}
func main() {

}

// push
func (lls *linkedListStack) push(value int) {
	lls.data.PushBack(value)
}

func (lls *linkedListStack) pop() any {
	if lls.empty() {
		return nil
	}
	e := lls.data.Back()
	lls.data.Remove(e)
	return e.Value
}
func (lls *linkedListStack) empty() bool {
	return lls.data.Len() == 0
}
func (s *linkedListStack) size() int {
	return s.data.Len()
}

func (s *linkedListStack) isEmpty() bool {
	return s.data.Len() == 0
}

func (s *linkedListStack) toList() *list.List {
	return s.data
}
