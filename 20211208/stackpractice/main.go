package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back == nil {
		return nil
	} else {
		return s.v.Remove(back)
	}
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func main() {
	Stack := NewStack()
	for i := 0; i < 5; i++ {
		Stack.Push(i)
	}
	v := Stack.Pop()
	for v != nil {

		fmt.Printf("%v->", v)
		v = Stack.Pop()
	}
}
