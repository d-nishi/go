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

var count int
var start *Node

func (l *List) Insert_Front(value string) {
	n := &Node{
		value: value,
		next:  l.head,
	}
	if l.head != nil {
		l.head.prev = n
	}
	l.head = n

	if l.head == nil {
		l.head = n
	}

}

func (l *List) return_elemN(n int) {
	count = 1
	for l.head != nil {
		l.head = l.head.next
		count = count + 1
		if count == n {
			break
		}
	}
	l.tail = l.head
	for l.tail != nil {
		fmt.Println(l.tail.value)
		l.tail = l.tail.prev
	}
}

func main() {

	l := List{}

	l.Insert_Front("11")
	l.Insert_Front("12")
	l.Insert_Front("13")
	l.Insert_Front("14")
	l.Insert_Front("15")
	l.Insert_Front("16")

	l.return_elemN(3)
}
