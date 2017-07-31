package main

import (
	"fmt"
	//	"os"
	"strings"
)

func main() {

	str := "teblack"
	stra := strings.Split(str, "")
	fmt.Println(insertion_sort(stra))
}

func insertion_sort(stra []string) []string {
	for i := 0; i < len(stra); i++ {
		for j := i; j < len(stra); j++ {
			if stra[i] < stra[j] {
				tmp := stra[i]
				stra[i] = stra[j]
				stra[j] = tmp
			}
		}
	}
return stra
}
