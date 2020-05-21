package main

import (
	"fmt"
)

type TreeNode struct {
	val    int
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
}

func main() {
	nums := []int{7, 1, 0, 5, 4, 10, 3, 6, 8, 20, 2}
	t := BST(nums)

	t.PrintInorder()

	fmt.Println(level_order(t))
}

//BFS search
func level_order(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	level_order := breadthfirstsearch(root)
	for _, level := range level_order {
		var lvl_vals []int
		for _, item := range level {
			lvl_vals = append(lvl_vals, item.val)
		}
		res = append(res, lvl_vals)
	}
	return res
}

func breadthfirstsearch(root *TreeNode) [][]*TreeNode {
	var res [][]*TreeNode
	level := []*TreeNode{root}
	for len(level) > 0 {
		res = append(res, level)
		var next_level []*TreeNode
		for _, item := range level {
			next_level = append(next_level, get_children(item)...)
		}
		level = next_level
	}
	return res
}

func get_children(n *TreeNode) []*TreeNode {
	var res []*TreeNode
	if n.left != nil {
		res = append(res, n.left)
	}
	if n.right != nil {
		res = append(res, n.right)
	}
	return res
}

//create BST
func BST(nums []int) *TreeNode {
	root := nums[0]
	t := &TreeNode{root, nil, nil, nil}
	fmt.Printf("value inserted: %v\n", root)

	for _, value := range nums[1:] {
		fmt.Printf("value inserted: %v\n", value)
		t.Insert(value)
	}
	return t
}

func (t *TreeNode) Insert(v int) error {
	switch {

	case v == t.val:
		return nil
	case v < t.val:
		if t.left == nil {
			t.left = &TreeNode{v, nil, nil, t}
			return nil
		}
		return t.left.Insert(v)
	case v > t.val:
		if t.right == nil {
			t.right = &TreeNode{v, nil, nil, t}
			return nil
		}
		return t.right.Insert(v)
	}

	return nil
}

func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}
	t.left.PrintInorder()
	fmt.Printf("val: %v\n", t.val)
	t.right.PrintInorder()
}
