package main

import (
	"errors"
	"fmt"
	"math"
)

type list struct {
	head *node
}

type node struct {
	val  int
	next *node
}

func main() {

	l := &list{}
	nums := []int{1, 2, 3, 4, 5}
	k := 4
	for i := 0; i < len(nums); i++ {
		l.Insert(nums[i])
	}

	l.Show()
	l.Rotate(k)
	l.Show()

}

func (l *list) Insert(nums int) error {
	n := &node{
		val:  nums,
		next: nil,
	}

	if l.head == nil {
		l.head = n
		return nil
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
	}
	return nil

}

func (l *list) Show() {
	if l.head == nil {
		fmt.Print("empty list")
		return
	}

	curr := l.head
	for curr.next != nil {
		fmt.Printf("Val: %v\n", curr.val)
		curr = curr.next
	}
	fmt.Printf("Val: %v\n", curr.val)
	fmt.Println()
}

func (l *list) Rotate(k int) error {
	if l.head == nil {
		errors.New("nil list")
	}

	curr := l.head
	head := l.head
	ll := 1

	for curr.next != nil {
		curr = curr.next
		ll++
	}

	if curr.next == nil {
		curr.next = head
	}

	endNode := math.Abs(float64(ll - k%ll))

	for endNode > 0 {
		curr = curr.next
		fmt.Printf("new curr: %v and new curr.next: %v\n", curr.val, curr.next.val)
		endNode--
		fmt.Printf("endNode: %v\n", endNode)
	}

	fmt.Println()

	temp := curr.next
	fmt.Printf("temp: %v\n", temp)
	curr.next = nil
	fmt.Printf("curr.next: %v\n", curr.next)
	l.head = temp
	fmt.Printf("l.head: %v\n", l.head)
	return nil
}
