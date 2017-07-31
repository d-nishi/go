package main

import "fmt"

import "strings"

var s [][]string
var loop int
var middle int

func main() {

	//define matrix
	s := [][]string{
		[]string{"1", "2", "4", "8"},
		[]string{"8", "16", "32", "64"},
		[]string{"64", "128", "256", "512"},
		[]string{"512", "1024", "2048", "4096"},
	}
	//count matrix loop no and middle value
	loop := (len(s)) / 2

	//rotate loop matrix values for each loop

	if loop > 0 {

		edg0 := s[0][0]
		edg1 := s[len(s)-1][len(s)-1]
		edg2 := s[0][len(s)-1]
		edg3 := s[len(s)-1][0]

		r := (len(s) - 1)
		for c := (len(s) - 2); c >= 0; c-- {
			if c > 0 {
				s[0][c+1] = s[0][c]
			}
			if c == 0 {
				s[0][c+1] = edg0
			}
			fmt.Println(s[0][c+1])
		}
		for c := 0; c < (len(s) - 1); c++ {
			if c < (len(s) - 2) {
				s[r][c] = s[r][c+1]
			}
			if c == (len(s) - 2) {
				s[r][c] = edg1
			}
			fmt.Println(s[r][c])
		}

		c := (len(s) - 1)
		for r := (len(s) - 2); r >= 0; r-- {
			if r > 0 {
				s[r+1][c] = s[r][c]
			}
			if r == 0 {
				s[r+1][c] = edg2
			}
			fmt.Println(s[r+1][c])
		}
		for r := 0; r < (len(s) - 1); r++ {
			if r < (len(s) - 2) {
				s[r][0] = s[r+1][0]
			}
			if r == (len(s) - 2) {
				s[r][0] = edg3
			}
			fmt.Println(s[r][0])
		}
		loop = (loop - 1)
	}

	for i := 0; i < (len(s)); i++ {
		fmt.Println(strings.Join(s[i], ","))
	}

}
