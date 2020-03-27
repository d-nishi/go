package main

import (
  "fmt"
)

func main() {

	a :=[]int{10,6,90,2,30,80}

	for i:=1; i<len(a)-1; i=i+1 {
	    j := i
	    for j>0 && a[j] < a[j-1] {
		a[j], a[j-1] = a[j-1], a[j]
		j--
	    }
	    fmt.Println(i)
	    fmt.Println(a)
	}

	fmt.Println(a)

}
