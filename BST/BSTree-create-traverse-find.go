package main

import (
	"errors"
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (t *Tree) Insert(v int) error {
	if t.root == nil {
		t.root = &Node{
			value: v,
		}
		return nil
	}
	return t.root.Insert(v)
}

func (n *Node) Insert(v int) error {
	if n == nil {
		return errors.New("Nil tree\n")
	}

	switch {
	case v == n.value:
		return nil
	case v < n.value:
		if n.left == nil {
			n.left = &Node{
				value: v,
			}
			return nil
		}
		return n.left.Insert(v)
	case v > n.value:
		if n.right == nil {
			n.right = &Node{
				value: v,
			}
			return nil
		}
		return n.right.Insert(v)
	}
	return nil
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.left, f)
	f(n)
	t.Traverse(n.right, f)
}

func main() {

	t := Tree{}
	fmt.Println(t.Insert(1), t.Insert(20), t.Insert(9), t.Insert(2))

	t.Traverse(t.root, func(n *Node) { fmt.Println(n.value) })
}
