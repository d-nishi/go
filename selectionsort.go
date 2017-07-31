package main

import (
	"fmt"
	"strings"
)

var C int

func main() {

	s1 := "blackish"
	sa := strings.Split(s1, "")

	fmt.Println("Original Slice", sa)
	fmt.Println("Sorted Slice", selectionsort(sa))

}

func selectionsort(sa []string) []string {

	s1 := []string{}
	for len(sa) > 0 {
		//find the min in the slice and move it to a new slice
		min := "z"
		for i := 0; i < len(sa); i++ {
			if sa[i] < min {
				min = sa[i]
				C = i
			}
		}
		s1 = append(s1, min)

		//create new slice without min value
		sa = append(sa[:C], sa[(C+1):]...)
	}

	return s1

}
