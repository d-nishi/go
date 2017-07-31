package main

import "fmt"

func main() {

	s := []int{6, 5, 1, 2, 8, 9, 3, 5}

	fmt.Printf("Original list:%v\n", s)
	fmt.Printf("Merged list:%v\n", Mergesort(s))

}

func Mergesort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := Mergesort(s[:n])
	r := Mergesort(s[n:])
	return Merge(l, r)
}

func Merge(l, r []int) []int {
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
