package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(printVertically("TO BE OR NOT TO BE"))
	fmt.Println(printVertically("HOW ARE YOU"))
}

func printVertically(s string) []string {

	s_arr := strings.Split(s, " ")
	s_a := make([][]string, len(s_arr))
	s_out := []string{}
	max_wl := 0

	//find the max word length in the string array
	for i := 0; i < len(s_arr); i++ {
		if max_wl < len(s_arr[i]) {
			max_wl = len(s_arr[i])
		}
	}

	//convert string array to string slice of slice of slice
	for i := 0; i < len(s_arr); i++ {
		s_a[i] = strings.Split(strings.Join(strings.Split(s_arr[i], ""), ""), "")
		if len(s_a[i]) < max_wl {
			fmt.Printf("s_a[%v] has less chars: %v\n", i, s_a[i])
			diff := max_wl - len(s_a[i])
			for diff > 0 {
				s_a[i] = append(s_a[i], " ")
				diff--
			}
		}
	}
	fmt.Printf("s_a: %v\n", s_a)

	//extract the jth element of each s_a[i] to form output array s_out
	for j := 0; j < max_wl; j++ {
		temp := []string{}
		for i := 0; i < len(s_arr); i++ {
			temp = append(temp, s_a[i][j])
			if i == (len(s_arr) - 1) {
				s_out = append(s_out, strings.TrimRight(strings.Join(temp, ""), " "))
			}
		}
	}
	fmt.Printf("s_out: %v\n", s_out)

	return s_out
}
