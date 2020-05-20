package main

import "fmt"
import "os"

func Map(S []string) map[string]int {
	m := make(map[string]int)
	m[S[0]] = 1
	for i := 0; i < cap(S); i++ {
		c := 1
		for j := (i + 1); j < cap(S); j++ {
			if S[i] == S[j] {
				c = c + 1
			}
			m[S[j]] = c
		}
	}
	return m
}

func main() {

	// 1: move the 2 strings to slices
	s1 := []string{"c", "i", "e", "p", "a", "n", "m"}
	s2 := []string{"a", "c", "n", "e", "m", "i", "a"}

	// 2: check the length of the 2 strings is equal
	if len(s1) != len(s2) {
		fmt.Println(s1, s2, "are not Anagrams")
		os.Exit(3)
	}

	// 3: create map for string1
	m := Map(s1)

	// 4: compare the elements of string2 with map of string1 to verify is anagram = true or false
	for i := 0; i < cap(s2); i++ {
		v, ok := m[s2[i]]
		if v == 0 {
			fmt.Println(s1, s2, "are not Anagrams")
			os.Exit(3)
		}
		if ok == false {
			fmt.Println(s1, s2, "are not Anagrams")
			os.Exit(3)
		}
		m[s2[i]] = (v - 1)
	}

	fmt.Println(s1, s2, "are Anagrams")

}
