package main

import "fmt"

func main() {

	s := []int{6, 5, 1, 3, 4, 7, 10, 9, 2}
	fmt.Printf("Original list:%v\n", (s))
	fmt.Printf("Sorted list:%v\n", Insersort(s))

}

func Insersort(s []int) []int {

	for i := 1; i < len(s); i++ {
		pivot := s[i]
		for j := (len(s[:i]) - 1); j >= 0; j-- {
			if pivot < s[j] {
				tmp := s[j]
				s[j] = pivot
				s[j+1] = tmp
			}
		}
	}
	return (s)

}
