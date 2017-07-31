package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value string
	prev  *Node
	next  *Node
}

var x int

func (l *List) StackInsertBack(value string) {
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

func (l *List) PopStack(x int) {

	if l.head == nil && l.tail == nil {
		fmt.Println("Null")
	}
	for l.head != nil && x > 0 {
		fmt.Println(l.head.value)
		l.head = l.head.next
		x = x - 1
	}

}

func (l *List) PopQueue(x int) {
	fmt.Println("x:", x)
	if l.head == nil && l.tail == nil {
		fmt.Println("Null")
	}
	for l.tail != nil && x > 0 {
		fmt.Println(l.tail.value)
		l.tail = l.tail.prev
		x = x - 1
	}
}

func main() {

	l := List{}

	l.StackInsertBack("15")
	l.StackInsertBack("15")
	l.StackInsertBack("10")
	l.StackInsertBack("10")
	l.StackInsertBack("22")
	l.StackInsertBack("10")

	l.PopStack(3)
	l.PopQueue(3)
}
