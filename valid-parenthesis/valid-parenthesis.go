package main

import (
	"fmt"
	"strings"
)

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	s    string
	prev *Node
	next *Node
}

func main() {

	fmt.Println(isValid("()[{{{[]}}}]"))
}

func isValid(s string) bool {

	a := strings.Split(s, "")
	l := List{}

	for i := 0; i < len(a); i = i + 1 {

		if a[i] == "(" || a[i] == "{" || a[i] == "[" {
			fmt.Printf("pushing: %v\n", a[i])
			l.PushElem(a[i])
		}

		if a[i] == ")" {
			out := l.PopElem("(")
			if out == "(" {
				fmt.Printf("popped: %v\n", out)
				continue
			} else {
				return false
			}
		}

		if a[i] == "}" {
			out := l.PopElem("{")
			if out == "{" {
				fmt.Printf("popped: %v\n", out)
				continue
			} else {
				return false
			}
		}

		if a[i] == "]" {
			out := l.PopElem("[")
			if out == "[" {
				fmt.Printf("popped: %v\n", out)
				continue
			} else {
				return false
			}

		}

	}

	count := l.ShowList()

	if count > 0 {
		return false
	}

	return true
}

func (l *List) PushElem(value string) {
	n := &Node{
		s: value,
	}

	if l.head == nil && l.tail == nil {
		l.head = n
		l.tail = n
		return
	}

	if l.head != nil {
		l.head.prev = n
		n.next = l.head
		l.head = n
		return
	}
}

func (l *List) PopElem(x string) string {

	if l.head == nil {
		fmt.Println("Nil")
		return ""
	}

	if l.head != nil {
		if l.head.s != x {
			fmt.Printf("%v does not match %v\n", x, l.head.s)
			return x
		}

		if l.head.s == x {
			fmt.Printf("%v matches %v\n", x, l.head.s)

			if l.head.next != nil {
				n := l.head.next
				n.prev = l.head.prev
				l.head = n
				return x
			}

			if l.head.next == nil {
				l.head = nil
				l.tail = nil
				return x
			}
		}
	}

	return ""
}

func (l *List) ShowList() int {

	elempushed_count := 0
	if l.tail == nil {
		fmt.Println("Nil list")
		return 0
	}

	for l.tail != nil {
		fmt.Println(l.tail.s)
		elempushed_count++
		l.tail = l.tail.prev
	}
	return elempushed_count

}
