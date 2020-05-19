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
	val  string
	next *Node
	prev *Node
}

func main() {
	s := "seattle"
	sa := strings.Split(s, "")
	l := &List{}
	n := &Node{}

	for i, _ := range sa {
		n = l.createList(&Node{
			val:  sa[i],
			next: nil,
			prev: nil,
		})
	}

	fmt.Println(n)

	l.showList()
}

func (l *List) createList(newnode *Node) *Node {
	if l.head == nil {
		l.head = newnode
		l.tail = newnode
		return l.head
	}

	currnode := l.head
	for currnode.next != nil {
		currnode = currnode.next
	}
	currnode.next = newnode
	newnode.prev = currnode
	l.tail = newnode

	return l.head
}

func (l *List) showList() {
	currnode := l.head
	for currnode != nil {
		fmt.Printf("currnode: %v\n", currnode)
		currnode = currnode.next
	}
}
