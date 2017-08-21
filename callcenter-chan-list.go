package main

import "fmt"

type List struct {
	head *Emp
	tail *Emp
}

type Emp struct {
	Level string
	State string
	Id    string
	prev  *Emp
	next  *Emp
}

var N, X, Y, Z, C int

func (el *List) CreateQueue(l string, s string, i string) {
	emp := &Emp{
		Level: l,
		State: s,
		Id:    i,
		prev:  el.tail,
	}

	if el.tail != nil {
		el.tail.next = emp
	}
	el.tail = emp

	if el.head == nil {
		el.head = emp
	}
}

func (el *List) ShowQueue() {
	n := el.head
	for n != nil {
		fmt.Println(n.Level, n.State, n.Id)
		n = n.next
	}
	fmt.Println()
}

func (el *List) DispatchCall(N int, s string, c chan int) {

	//in the list: find employee with Level == s; turn their state to busy;count calls dispatched from total N
	for n := el.head; n != nil; n = n.next {
		if n.Level == s && n.State == "free" {
			n.State = "busy"
			C = C + 1
		}
	}
	N = N - C
	c <- N
}

func main() {
	el := List{}
	el.CreateQueue("r", "free", "1")
	el.CreateQueue("m", "free", "2")
	el.CreateQueue("d", "free", "3")
	el.CreateQueue("r", "free", "4")
	el.CreateQueue("r", "free", "5")
	el.CreateQueue("m", "free", "6")

	fmt.Println("Existing employees:\n")
	el.ShowQueue()

	c := make(chan int)
	N := 14

	go el.DispatchCall(N, "r", c)
	X = <-c
	fmt.Println("Employees after", (N - X), "calls were dispatched:\n")
	el.ShowQueue()

	go el.DispatchCall(N, "m", c)
	Y = <-c
	fmt.Println("Employees after", (N - Y), "calls were dispatched:\n")
	el.ShowQueue()

	go el.DispatchCall(N, "d", c)
	Z = <-c
	fmt.Println("Employees after", (N - Z), "calls were dispatched:\n")
	el.ShowQueue()
}
