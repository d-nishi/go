package main

import "fmt"

var m map[string]int
var count int
var ok bool
var elem int

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
	for l.head != nil {
		fmt.Println(l.head.value)
		l.head = l.head.next
	}
	fmt.Println()
}

func (l *List) Insert_Back(value string) {
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

func (l *List) Size() int {
	count = 0
	for l.head != nil {
		l.head = l.head.next
		count = count + 1
	}
	return count
}

func (l *List) NoDuplicates() *List {
	m := make(map[string]int)

	for l.tail != nil {
		elem, ok = m[l.tail.value]
		if ok == false {
			m[l.tail.value] = 1
			l.tail = l.tail.prev
		}
		if ok == true && l.tail.prev != nil {
			m[l.tail.value] = elem + 1
			cur := l.tail
			l.tail = l.tail.prev
			l.tail.next = cur.next
		}
		if ok == true && l.tail.prev == nil {
			m[l.tail.value] = elem + 1
			l.tail = l.tail.next
			l.tail.prev = nil
			l.head = l.tail
			break
		}
	}
	fmt.Println(m)
	return l
}

func main() {

	l := List{}

	l.Insert_Back("15")
	l.Insert_Back("15")
	l.Insert_Back("10")
	l.Insert_Back("10")
	l.Insert_Back("22")
	l.Insert_Back("10")

	fmt.Printf("head: %v\n", l.head)
	fmt.Printf("tail: %v\n", l.tail)

	fmt.Println("No Duplicate list:")
	new := l.NoDuplicates()
	new.Show()
}
