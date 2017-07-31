package main

import "fmt"
import "os"

var strarr [][]string
var r int
var c int

//function to make row & column elements = "0" in a matrix
func make_zero(r int, c int) {

	for r1 := 0; r1 < len(strarr); r1++ {
		strarr[r1][c] = "0"
	}
	for c1 := 0; c1 < len(strarr[0]); c1++ {
		strarr[r][c1] = "0"
	}
}

func main() {

	//read a MxN matrix in a slice of slices

	strarr = [][]string{
		[]string{"2", "3", "3", "5"},
		[]string{"2", "3", "4", "5"},
		[]string{"3", "0", "4", "5"},
		[]string{"2", "3", "2", "5"},
		[]string{"1", "1", "2", "4"},
	}

	//find if an elem in matrix = "0". If yes, call function to set row & column elements to "0" and exit.

	for r = 0; r < len(strarr); r++ {
		for c = 0; c < len(strarr[0]); c++ {
			if strarr[r][c] == "0" {
				make_zero(r, c)
				fmt.Println("Elements in", "row:", r, "column:", c, "set to 0")
				//print the matrix
				for r = 0; r < len(strarr); r++ {
					fmt.Println(strarr[r])
				}
				os.Exit(0)
			}
		}
	}

	//print the matrix
	for r = 0; r < len(strarr); r++ {
		fmt.Println(strarr[r])
	}

}
