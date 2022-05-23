package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}
func NewQueue()

func main() {
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)

	v.InsertBefore(3, e4)
	v.InsertAfter(2, e1)

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, "->")
	}
}
