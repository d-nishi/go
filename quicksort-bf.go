package main

import (
  "fmt"
)

func main(){
	arr := []int{10,30,80,50,20,70}
	(sort(arr, 0, len(arr)-1))
}

func sort(a []int, start int, end int) {

	if (end - start) < 0 {
		return
	}

	pivot := a[end]
	splitIndex := start

	for i := start; i < end; i++ {
	    if a[i] < pivot {
	       a[splitIndex], a[i] = a[i], a[splitIndex]
	       splitIndex++
	    }
	}

	a[end] = a[splitIndex]
	a[splitIndex] = pivot

	sort(a, start, splitIndex-1)
	sort(a, splitIndex+1, end)

	fmt.Println(a)

}
