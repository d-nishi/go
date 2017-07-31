package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	prev  *Node
	next  *Node
}

var min int

func (l *List) StackInsertBack(value int) {
	n := &Node{
		value: value,
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

func (l *List) StackMin() {
	min = l.head.value
	for l.head != nil {
		fmt.Printf("%v\n", l.head.value)
		if l.head.value < min {
			min = l.head.value
		}
		l.head = l.head.next
	}
	fmt.Println("Min in the Stack:", min)
}

func main() {

	l := List{}

	l.StackInsertBack(2)
	l.StackInsertBack(2)
	l.StackInsertBack(1)
	l.StackInsertBack(6)
	l.StackInsertBack(3)
	l.StackInsertBack(9)
	l.StackInsertBack(5)

	l.StackMin()
}
