package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "abbc"
	b := "babcabbacaabcbabcacbb"
	m1 := make(map[rune]int)
	m2 := m1
	perm := []string{}
	all_perms := []string{}

	for _, v := range s {
		m1[v] = m1[v] + 1
	}
	fmt.Println(m1)

	j := 0
	matches := 0
	for i, v := range b {
		_, ok := m2[v]

		if ok == true && matches < 4 {
			perm = append(perm, string(v))
			matches++
			m2[v] = m2[v] - 1
			continue
		}

		if ok == false && matches < 4 {
			matches = 0
			perm = []string{}
			m2 = m1
			j = j + 1
			i = j
		}

		if matches == 4 {
			all_perms = append(all_perms, strings.Join(perm, ""))
			matches = 0
			perm = []string{}
			m2 = m1
			i = i + 1
			j = i
		}
	}

	fmt.Printf("all perms: %v\n", all_perms)

}
