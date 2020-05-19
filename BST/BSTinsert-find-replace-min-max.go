package main

import (
	"errors"
	"fmt"
)

type Tree struct {
	Value  int
	Left   *Tree
	Right  *Tree
	Parent *Tree
}

func (t *Tree) Insert(v int) error {

	switch {

	case v == t.Value:
		return nil
	case v < t.Value:
		if t.Left == nil {
			t.Left = &Tree{v, nil, nil, t}
			return nil
		}
		return t.Left.Insert(v)
	case v > t.Value:
		if t.Right == nil {
			t.Right = &Tree{v, nil, nil, t}
			return nil
		}
		return t.Right.Insert(v)
	}
	return nil
}

func (t *Tree) Find(v int) (*Tree, bool) {
	switch {
	case v == t.Value:
		return t, true
	case v < t.Value:
		if t.Left == nil {
			return nil, false
		}

		return t.Left.Find(v)
	case v > t.Value:
		if t.Right == nil {
			return nil, false
		}

		return t.Right.Find(v)
	}
	return nil, false
}

func (t *Tree) Maximum() int {
	if t.Right == nil {
		return t.Value
	}
	return t.Right.Maximum()
}

func (t *Tree) Minimum() int {
	if t.Left == nil {
		return t.Value
	}
	return t.Left.Minimum()
}

func (t *Tree) Delete(v int) error {
	tree, present := t.Find(v)

	if !present {
		return errors.New("Element not present in the tree")
	}
	return tree.Replace(v)

}

func (t *Tree) Replace(v int) error {
	if t.Right == nil && t.Left == nil {
		t.Parent = nil
		return nil
	}

	switch {
	case t.Right == nil:
		t.Parent.Left = t.Left
	case t.Left == nil:
		t.Parent.Right = t.Right
	default:
		min := t.Right.Minimum()
		t.Value = min
		t.Right.Delete(min)
	}
	return nil
}

func BinaryTree(numbers []int) *Tree {
	root := numbers[0]
	tree := Tree{root, nil, nil, nil}
	for _, value := range numbers[1:] {
		tree.Insert(value)
	}
	return &tree
}

func (t *Tree) PrintInorder() {
	if t == nil {
		return
	}

	t.Left.PrintInorder()
	fmt.Printf("t: %v\n", t)
	t.Right.PrintInorder()
}

func main() {

	nums := []int{3, 4, 8, 9, 10, 0, 5}

	t := BinaryTree(nums)
	t.Delete(4)

	t.PrintInorder()
}
