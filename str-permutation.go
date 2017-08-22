package main

import "fmt"
import "strings"

var str1 string
var strarr1 []string
var str2 string
var strarr2 []string

var m map[string]int
var count int

func is_permutation(strarr2 []string) bool {
	count = 0
	for i := 0; i < len(strarr2); i++ {
		v, ok := m[strarr2[i]]
		fmt.Println(strarr2[i], v)
		if ok == false {
			count = 1
			break
		}
		continue
	}

	if count == 1 {
		return false
	}
	return true
}

func main() {

	str1 = "abaham"
	str2 = "brhma"

	//convert string to char slices
	strarr1 = strings.Split(str1, "")
	strarr2 = strings.Split(str2, "")

	//build a hash table for string1
	m = make(map[string]int)
	m[strarr1[0]] = 1
	for i := 1; i < len(strarr1); i++ {
		v, ok := m[strarr1[i]]
		if ok == false {
			m[strarr1[i]] = 1
			continue
		}
		m[strarr1[i]] = (v + 1)
	}

	fmt.Println(is_permutation(strarr2))

}
