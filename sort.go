package main

import "fmt"

func Map(S []string) ([]string, map[string]int) {
	m := make(map[string]int)

	for i := 0; i < cap(S); i++ {
		c := 1
		for j := (i + 1); j < cap(S); j++ {
			if S[i] == S[j] {
				c = c + 1
			}
			m[S[j]] = c
		}
	}
	return S, m

}

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

	// 1: move string to slice

	sa := []string{"i", "c", "e", "a", "a", "n", "m"}

	// 2: map unique letters in string

	fmt.Println(Map(sa))
	fmt.Println(Sort(sa))

}
