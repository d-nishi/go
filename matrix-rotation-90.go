package main

import "fmt"
import "strings"

var strarr [][]string
var str string
var tmp []string
var loop int
var edg string

func main() {

	strarr = [][]string{
		[]string{"1", "2", "4"},
		[]string{"8", "16", "32"},
		[]string{"64", "128", "256"},
	}

	for i := 0; i < len(strarr); i++ {
		fmt.Println(strarr[i])
	}
	fmt.Println()

	str = strings.Join(strarr[0], ",")
	edg = strarr[0][len(strarr)-1]
	loop = len(strarr) % 2

	for loop > 0 {
		for i := (len(strarr) - 1); i >= 0; i-- {
			strarr[0][i] = strarr[(len(strarr)-1)-i][0]
		}
		for i := 0; i < len(strarr); i++ {
			strarr[i][0] = strarr[len(strarr)-1][i]
		}
		for i := 0; i < len(strarr); i++ {
			strarr[len(strarr)-1][i] = strarr[len(strarr)-1-i][len(strarr)-1]
			if i == (len(strarr) - 1) {
				strarr[len(strarr)-1][i] = edg
			}
		}
		for i := (len(strarr) - 1); i >= 0; i-- {
			tmp = strings.Split(str, ",")
			strarr[i][len(strarr)-1] = tmp[i]
		}
		loop = loop - 1
	}

	for i := 0; i < len(strarr); i++ {
		fmt.Println(strarr[i])
	}

}
