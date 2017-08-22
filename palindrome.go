package main

import "fmt"
import "strconv"
import "strings"

func palindrome_chk(a []string) string {

	r := "true"
	for i := 0; i < len(a); i++ {
		if a[i] != a[len(a)-i-1] {
			r = "false"
			break
		}
	}
	return r
}

func main() {

	var a []string
	var tst string
	palin := []int{0}

	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			s := strconv.Itoa(i * j)
			a = strings.SplitAfter(s, "")
			tst = palindrome_chk(a)
			if tst == "true" {
				if (i * j) > palin[0] {
					palin[0] = (i * j)
				}
				fmt.Println(palin)
				break
			}
		}
	}

	fmt.Println(palin, tst)

}
