package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	dummyDict := []string{"from", "time", "item", "form", "off", "test"}
	Anagram(dummyDict)
}

func Anagram(dummyDict []string) {

	list := make(map[string][]string)

	for _, word := range dummyDict {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		//for _, w := range words {
		//	fmt.Print(w, " ")
		//}
		fmt.Print(words)
		fmt.Println()
	}
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
