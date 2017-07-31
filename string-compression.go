package main

import "fmt"
import "strings"
import "strconv"

var str string
var strarr []string
var strarrnew []string
var count int

var m map[string]int

func main() {

	str = "aabbbcccdefgggghyyyyxxxxtrsuv"
	strarr = strings.Split(str, "")

	strarrnew = []string{}

	fmt.Println(len(strarr))
	//compare element 0 to element 1 in strarray. if they match, increase count and compare elem 1 with elem 2.
	count = 1
	for i := 1; i < (len(strarr)); i++ {
		if strarr[i-1] != strarr[i] {
			strarrnew = append(strarrnew, strarr[i-1])
			strarrnew = append(strarrnew, strconv.Itoa(count))
			if i == (len(strarr) - 1) {
				strarrnew = append(strarrnew, strarr[i])
				strarrnew = append(strarrnew, strconv.Itoa(count))
			}
			count = 1
			continue
		}
		count = count + 1
	}

	fmt.Println(strarr)
	fmt.Println(strarrnew)

}
