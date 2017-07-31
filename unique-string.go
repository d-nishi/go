package main

import "fmt"
import "strings"

var str string
var char string
var m map[string]int

func main() {

	//read the string and split into characters

	str = "bnhfgtmrsuvwxyzaeoipq"
	char := strings.Split(str, "")

	//build a hash table for the char mapped to count
	m := make(map[string]int)
	m[char[0]] = 1

	for i := 1; i < len(char); i++ {
		v, ok := m[char[i]]
		if ok == false {
			m[char[i]] = 1
			continue
		}
		if v >= 1 {
			fmt.Println("Not Unique")
			break
		}
	}

	fmt.Println("Unique", m)

}
