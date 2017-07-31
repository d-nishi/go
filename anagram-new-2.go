package main

import "fmt"
import "os"

func Sort(S []string) []string {
	for i := 0; i < cap(S); i++ {
		for j := (i + 1); j < cap(S); j++ {
			if S[i] >= S[j] {
				tmp := S[i]
				S[i] = S[j]
				S[j] = tmp
			}
		}
	}
	return S
}

func main() {

	// 1: move the 2 strings to slices
	s1 := []string{"c", "i", "e", "p", "a", "n", "m"}
	s2 := []string{"c", "i", "n", "e", "m", "p", "a"}

	// 2: check the length of the 2 strings is equal
	if len(s1) != len(s2) {
		fmt.Println(s1, s2, "are not Anagrams")
		os.Exit(3)
	}

	// 3: sort both strings
	Sort(s1)
	Sort(s2)

	// 4: compare the same element of both arrays to verify is anagram = true or false
	for i := 0; i < cap(s1); i++ {
		if s1[i] != s2[i] {
			fmt.Println(s1, s2, "are not Anagrams")
			os.Exit(3)
		}
	}

	fmt.Println(s1, s2, "are Anagrams")

}
