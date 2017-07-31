package main

import "fmt"
import "strings"

var str1 []string
var str2 []string
var x int
var result bool

func isSubstring(strnew1 []string, x int, str2 []string) bool {

	if (strings.Join(str1, "")) != (strings.Join(str2, "")) {
		result = false
	}
	if (strings.Join(str1, "")) == (strings.Join(str2, "")) {
		result = true
	}

	return result
}

func main() {

	str1 = []string{"w", "a", "t", "e", "r", "b", "o", "t", "t", "l", "e"}
	str2 = []string{"e", "r", "b", "o", "t", "t", "l", "e", "w", "a", "t"}

	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[0] {
			x = i
			//re-organize str1 split from location x to the end 1st & str1 split from 0 to x 2nd
			strnew1 := []string{(strings.Join(str1[x:], "")), (strings.Join(str1[0:x], ""))}
			//create slice with new concatenated string joined together
			str1 = []string{(strings.Join(strnew1, ""))}
			//join characters of str2 slice to make len(str1) = len(str2)
			str2 = []string{(strings.Join(str2, ""))}
			//call function to check if the 2 new slices are equal. if yes, str2 is a rotation of the original str1
			fmt.Println(isSubstring(str1, x, str2))
		}
	}

}
