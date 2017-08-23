package main

import (
	"errors"
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	value string
	left  *Node
	right *Node
}

func (n *Node) Insert(value string) error {
	if n == nil {
		return errors.New("Nil")
	}

	switch {
	case value == n.value:
		return nil
	case value < n.value:
		if n.left == nil {
			n.left = &Node{value: value}
			return nil
		}
		n.left.Insert(value)
		return nil
	default:
		if n.right == nil {
			n.right = &Node{value: value}
			return nil
		}
		n.right.Insert(value)
		return nil
	}
}

func (t *Tree) Insert(value string) error {
	if t.root == nil {
		t.root = &Node{value: value}
		return nil
	}
	t.root.Insert(value)
	return nil
}

func (t *Tree) Walk(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	f(n)
	t.Walk(n.left, f)
	t.Walk(n.right, f)
}

//func (t *Tree) Find(value string) bool {
//	if t.root == nil {
//		return false
//	}
//	return t.root.Find(value)
//}

func (n *Node) Find(value string) bool {
	if n == nil {
		return false
	}
	switch {
	case value == n.value:
		return true
	case value < n.value:
		return n.left.Find(value)
	default:
		return n.right.Find(value)
	}
}

func main() {

	t := Tree{}
	values := []string{"3", "2", "1", "4", "5"}
	for i := 0; i < len(values); i++ {
		err := t.Insert(values[i])
		if err != nil {
			fmt.Println("error when inserting in tree", err)
		}
	}

	t.Walk(t.root, func(n *Node) { fmt.Println("Tree value\n", n.value) })
	fmt.Println("Found value:", "5", t.root.Find("5"))
}
