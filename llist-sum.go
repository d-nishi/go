package main

import "fmt"
import "math"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value float64
	prev  *Node
	next  *Node
}

var count int
var size int
var intlist float64

func (l *List) FindSize() int {
	count = 0
	for l.head != nil {
		count = count + 1
		l.head = l.head.next
	}
	return count
}

func (l *List) Show() {
	n := l.head
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
}

func (l *List) InsertFront(value float64) {
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

func (l *List) ConvertToInteger(size int) float64 {
	n := l.tail
	s := 0
	intlist = 0
	for n != nil && s <= size {
		intlist = intlist + ((n.value) * (math.Pow10(s)))
		n = n.prev
		s = s + 1
	}

	return intlist

}

func main() {

	l1 := List{}
	l2 := List{}

	l1.InsertFront(9)
	l1.InsertFront(6)
	l1.InsertFront(8)

	l2.InsertFront(6)
	l2.InsertFront(5)
	l2.InsertFront(4)
	l2.InsertFront(9)

	l1.Show()
	l2.Show()

	s1 := l1.FindSize()
	s2 := l2.FindSize()

	if s1 >= s2 {
		size = s1
	}
	size = s2

	intlist1 := l1.ConvertToInteger(size)
	intlist2 := l2.ConvertToInteger(size)
	sum := intlist1 + intlist2
	fmt.Printf("Sum of %v + %v = %v\n", intlist1,intlist2, sum)

}
