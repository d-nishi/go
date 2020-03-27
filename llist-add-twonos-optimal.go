package main

import (
  "fmt"
  "math"
)

type Node struct {
     val int
     next *Node
     prev *Node
}

type List struct {
     head *Node
     tail *Node
}

func (l *List)insertback_list(v int) {
        n := &Node {
             val: v,
             prev: l.tail,
        }

	if l.tail != nil {
	     l.tail.next = n
	}
        l.tail = n

	if l.head == nil {
	     l.head = n
	}
}

func (l *List)show_list() {
	n := l.head
	fmt.Println("print list elements")
	for n != nil {
	    fmt.Printf("%d\n", n.val)
	    n = n.next
	}

}

func add_lists_return_reverselist(l1 *List, l2 *List) (List){
        //push list elements to a number
	sum1 := 0
	sum2 := 0
        counter1 := 0
        counter2 := 0

        n1 := l1.head
        for n1 != nil {
           sum1 = (int(math.Pow(10,float64(counter1))) * n1.val) + sum1
	   n1 = n1.next
	   counter1 = counter1 + 1
	}

	n2 := l2.head
	for n2 != nil {
           sum2 = (int(math.Pow(10,float64(counter2))) * n2.val) + sum2
           n2 = n2.next
	   counter2 = counter2 + 1
	}

	fmt.Println(sum1)
	fmt.Println(sum2)

	sum := sum1 + sum2

	//move the final sum to a list	
	l := List{}
	for sum > 0 {
		endlist_elem := sum%10
		sum = sum/10
		l.insertback_list(endlist_elem)
	}

	return l
}

func main() {
	l1 := List{}
        l2 := List{}

        l1.insertback_list(9)
        l1.insertback_list(9)
	l1.insertback_list(9)

        l2.insertback_list(9)
        l2.insertback_list(9)
        l2.insertback_list(9)

	l := add_lists_return_reverselist(&l1, &l2)
	l.show_list()
}
