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

func (l *List) Show() {
	n := l.head
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
	fmt.Println()
}

func (l *List) InsertBack(value string) {
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

func main() {

	l := List{}

	l.InsertBack("(")
	l.InsertBack(")")
	l.InsertBack("[")
	l.InsertBack("{")
	l.InsertBack("{")
	l.InsertBack("{")
	l.InsertBack("[")
	l.InsertBack("]")
	l.InsertBack("}")
	l.InsertBack("}")
	l.InsertBack("}")
	l.InsertBack("]")

	fmt.Printf("head: %v\n", l.head.value)
	fmt.Printf("tail: %v\n", l.tail.value)

	l.Show()
	fmt.Printf("head: %v\n", l.head.value)
	fmt.Printf("tail: %v\n", l.tail.value)
}
