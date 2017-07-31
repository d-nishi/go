package main

import "fmt"
import "strings"

var str1 string
var str2 string
var strarr1 []string
var strarr2 []string

var m map[string]int
var count int

func main() {

	str1 = "bales"
	str2 = "bill"

	//create a hash table for string 1

	strarr1 = strings.Split(str1, "")
	strarr2 = strings.Split(str2, "")

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

	fmt.Println(m)

	//compare each character of string 2 with hash table and count differences
	count = 0
	for i := 0; i < len(strarr2); i++ {
		v, ok := m[strarr2[i]]
		if ok == false || v == 0 {
			count = count + 1
			continue
		}
		m[strarr1[i]] = (v - 1)
	}

	fmt.Println(m)

	//if difference in char count > 1 then false else true

	if count > 1 {
		fmt.Println("false")
	}
	if count <= 1 {
		fmt.Println("true")
	}

}
