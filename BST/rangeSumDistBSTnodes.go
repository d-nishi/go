package main

import "fmt"

type Tree struct {
	Val    int
	Parent *Tree
	Left   *Tree
	Right  *Tree
}

func main() {

	nums := []int{10, 5, 15, 3, 7, 0, 18}
	t := BST(nums)

	t.PrintInorder()

	L := nums[5]
	R := nums[0]
	sum := (rangeSumBST(t, L, R))
	dist := (rangeDistBST(t, L, R))

	fmt.Printf("sum between node with val: %v and %v is %v\n", L, R, sum)

	fmt.Printf("distance between node with val: %v and %v is %v\n", L, R, dist-1)

}

func rangeDistBST(root *Tree, L int, R int) int {

	if root == nil {
		return 0
	}

	dist := 0

	if root.Val >= L && root.Val <= R {
		dist = dist + 1
		fmt.Print(dist)
		fmt.Println()
	}

	if root.Val > L {
		dist = dist + rangeDistBST(root.Left, L, R)
	}

	if root.Val < R {
		dist = dist + rangeDistBST(root.Right, L, R)
	}

	return dist

}

func rangeSumBST(root *Tree, L int, R int) int {

	if root == nil {
		return 0
	}

	sum := 0

	if root.Val >= L && root.Val <= R {
		sum = sum + root.Val
		fmt.Printf("sum: %v\n", sum)
	}

	if root.Val > L {
		sum = sum + rangeSumBST(root.Left, L, R)
	}

	if root.Val < R {
		sum = sum + rangeSumBST(root.Right, L, R)
	}

	return sum
}

func BST(nums []int) *Tree {

	root := nums[0]
	t := &Tree{root, nil, nil, nil}

	for i := 1; i < len(nums); i++ {
		t.Insert(nums[i])
	}

	return t
}

func (t *Tree) PrintInorder() {
	if t == nil {
		return
	}

	t.Left.PrintInorder()
	fmt.Printf("node: %v\n", t)
	t.Right.PrintInorder()

}

func (t *Tree) Insert(val int) error {

	switch {
	case t.Val == val:
		fmt.Print("node present")
		fmt.Println()
		return nil
	case t.Val < val:
		if t.Right == nil {
			t.Right = &Tree{
				Val:    val,
				Parent: t,
				Left:   nil,
				Right:  nil,
			}
			return nil
		}
		return t.Right.Insert(val)
	case t.Val > val:
		if t.Left == nil {
			t.Left = &Tree{
				Val:    val,
				Parent: t,
				Left:   nil,
				Right:  nil,
			}
			return nil
		}
		return t.Left.Insert(val)
	}

	return nil
}
