package main

import "fmt"
import "strings"

var str string
var strarr []string
var last_elem int
var count int

func main() {

	str = "Mr John Smith     "

	strarr = strings.Split(str, " ")

	for i := (len(strarr) - 1); i >= 0; i-- {
		if strarr[i] == "" {
			count = count + 1
			continue
		}
		break
	}

	strarr = (strarr[0:(len(strarr) - count)])

	fmt.Println(strings.Join(strarr, "%20"))

}
