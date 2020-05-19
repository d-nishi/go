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
	prev  *Node
	next  *Node
}

func main() {

	l := List{}

	l.QueueElem(10)
	l.QueueElem(40)
	l.QueueElem(1)
	l.QueueElem(15)
	l.QueueElem(80)

	l.Dequeue(3)

}

func (l *List) QueueElem(x int) {
	n := &Node{
		value: x,
		prev:  l.tail,
	}

	if l.tail != nil {
		l.tail.next = n
	}
	l.tail = n

	if l.head == nil {
		l.head = n
	}

}

func (l *List) Dequeue(x int) {
	if l.head == nil {
		fmt.Println("Nil list")
		return
	}

	for l.head != nil {
		fmt.Printf("value: %d\n", l.head.value)
		l.head = l.head.next
	}

}
