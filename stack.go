package main

import (
	"fmt"
)

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func main() {

	l := List{}

	l.PushElem(12)
	l.PushElem(1)
	l.PushElem(52)
	l.PushElem(14)
	l.PushElem(6)
	l.PushElem(33)

	l.PopElem()
}

func (l *List) PushElem(x int) {
	n := &Node{
		value: x,
		next:  l.head,
	}

	if l.head != nil {
		l.head.prev = n
	}
	l.head = n

	if l.tail == nil {
		l.tail = n
	}

}

func (l *List) PopElem() {

	for l.head != nil {
		fmt.Printf("value: %d\n", l.head.value)
		l.head = l.head.next
	}

	if l.head == nil {
		fmt.Println("Nil")
	}

}
