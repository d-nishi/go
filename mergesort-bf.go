package main

import (
   "fmt"
)

func main() {
	a := []int{10,6,2,80,90,30,8}

	fmt.Printf("Merged list:%d\n", MergeSort(a))

}

func MergeSort(s []int) []int{
	if len(s) <= 1 {
   	       return s
	}
	n := len(s)/2
	l := MergeSort(s[:n])
	r := MergeSort(s[n:])
	return Merge(l,r)
}

func Merge(l,r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
		   return append(ret, r...)
		}
		if len(r) == 0 {
		   return append(ret, l...)
		}
		if l[0] < r[0] {
		   ret = append(ret, l[0])
		   l = l[1:]
		} else {
		   ret = append(ret, r[0])
		   r = r[1:]
		}
	}
	return ret

}
