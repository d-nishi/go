package main

import (
	"fmt"
	"errors"
)

type Tree struct {
	root *Node
}

type Node struct {
	value string
	data  string
	left  *Node
	right *Node
}

func (t *Tree) Insert(value, data string) error {
	if t.root == nil {
		t.root = &Node{
			value: value,
			data:  data,
		}
		return nil
	}
	return t.root.Insert(value, data)
}

func (n *Node) Insert(value, data string) error {
	if n == nil {
		return errors.New("Nil tree")
	}
	switch {

	case value == n.value:
		return nil
	case value < n.value:
		if n.left == nil {
			n.left = &Node{
				value: value,
				data:  data,
			}
			return nil
		}
		return n.left.Insert(value, data)
	case value > n.value:
		if n.right == nil {
			n.right = &Node{
				value: value,
				data:  data,
			}
			return nil
		}
		return n.right.Insert(value, data)
	}
	return nil
}

func (n *Node) Find(s string) (string, bool) {
	if n == nil {
		return "", false 
	}

	switch {

	case s == n.value:
		return n.data, true
	case s < n.value:
		return n.left.Find(s)
	default:
		return n.right.Find(s)
	}
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n.right == nil {
		return n, parent
	}
	return n.right.findMax(n)
}

func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("Nil node to replace")
	}
	if n == parent.left {
		parent.left = replacement
		return nil
	}
	parent.right = replacement
	return nil
}

func (n *Node) deleteNode(s string, parent *Node) error {
	if n == nil {
		return errors.New("Nil node to delete")
	}

	switch {
	case s < n.value:
		n.left.deleteNode(s, n)
		return nil
	case s > n.value:
		n.right.deleteNode(s, n)
		return nil
	default:
		if n.left == nil && n.right == nil {
			n.replaceNode(parent, nil)
			return nil
		}
		if n.left == nil {
			n.replaceNode(parent, n.right)
			return nil
		}
		if n.right == nil {
			n.replaceNode(parent, n.left)
			return nil
		}

		replacement, replparent := n.left.findMax(n)

		n.value = replacement.value
		n.data = replacement.data

		return replacement.deleteNode(replacement.value, replparent)

	}
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	f(n)
	t.Traverse(n.left, f)
	t.Traverse(n.right, f)
}

func main() {

	t := Tree{}
	fmt.Println(t.Insert("com","1"), t.Insert("google","2"), t.Insert("www","3"), t.Insert("abc","4"))

	t.Traverse(t.root, func(n *Node) { fmt.Println(n.value) })

}
