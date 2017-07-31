package main

import "fmt"

var C int

type Node struct {
	value string
	prev  *Node
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Sort(C int) {
	for ; C > 0; C-- {
		n := l.head
		for n != nil {
			if n.next == nil {
				break
			}
			if n.value > n.next.value {
				tmp := n.next.value
				n.next.value = n.value
				n.value = tmp
			}
			n = n.next
		}
	}
}

func (l *List) Show() {
	n := l.tail
	C = 0
	for n != nil {
		fmt.Println(n.value)
		C = C + 1
		n = n.prev
	}
}

func (l *List) Insertfront(value string) {
	n := &Node{
		value: value,
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

func (l *List) Insertback(value string) {
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
	l.Insertback("11")
	l.Insertback("13")
	l.Insertfront("09")
	l.Insertfront("07")
	l.Insertfront("17")
	l.Insertfront("29")
	fmt.Println("Original List")
	l.Show()
	l.Sort(C)
	fmt.Println("Bubble Sorted List")
	l.Show()
}
