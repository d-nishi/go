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
        //push list elements to an array
	a1 := []int{}
        a2 := []int{}

        n1 := l1.head
        for n1 != nil {
	   fmt.Println(n1.val)
           a1 = append(a1, n1.val)
           n1 = n1.next
	}

	n2 := l2.head
	for n2 != nil {
	   fmt.Println(n2.val)
           a2 = append(a2, n2.val)
           n2 = n2.next
	}

	fmt.Printf("print array1 length%d\n", a1)
	fmt.Printf("print array1 length%d\n", a2)

	//add the array elements and get the sum
	sum1 := 0
	sum2 := 0

	for i:=(len(a1)-1); i>=0; i-- {
	    sum1 = sum1 + int(math.Pow(10,float64(i))) * a1[i]

	}
	fmt.Printf("print array1 no: %d\n",sum1)

	for i:=(len(a2)-1); i>=0; i-- {
	    sum2 = sum2 + int(math.Pow(10,float64(i))) * a2[i]

	}
	fmt.Printf("print array2 no: %d\n",sum2)

	sum := sum1 + sum2

	//move the array elements to a list	
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
