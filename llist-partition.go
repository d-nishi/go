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

func (l *List) Show() {
	n := l.head
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
	fmt.Println()
}

func (l *List) InsertBack(value int) {
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

func (l *List) Partition(n int) *List {
	oldhead := l.head
	for l.tail != nil {
		if l.tail.value < n {
			l.tail = l.tail.prev
		}
		if l.tail.value >= n {
			cur := l.tail
			l.tail = l.tail.prev
			l.tail.next = cur.next
			fmt.Println("2nd", l.tail.value)

			cur1 := l.head
			l.head = cur
			l.head.next = cur1
			cur1.prev = l.head
			fmt.Println("3rd", l.head.value, l.head.prev.value)
		}
		if l.tail == oldhead {
			break
		}
	}
	return l
}

func main() {

	l := List{}

	l.InsertBack(3)
	l.InsertBack(5)
	l.InsertBack(8)
	l.InsertBack(5)
	l.InsertBack(10)
	l.InsertBack(2)
	l.InsertBack(1)

	new := l.Partition(5)
	new.Show()

}
