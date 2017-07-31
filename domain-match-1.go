package main

import "fmt"
import "strings"

var cdom_name string
var cdom_word []string
var cdom_match string
var count int

func compare(a string) string {
	pdom := [][]string{
		[]string{"com"},
		[]string{"uk"},
		[]string{"co.uk"},
		[]string{"gov.uk"},
		[]string{"nhs.uk"},
		[]string{"org.uk"},
		[]string{"ca.gov"},
		[]string{"ccc.co.uk"},
	}

	for i := 0; i < cap(pdom); i++ {
		if pdom[i][0] != a {
			continue
		}
		cdom_match = pdom[i][0]
	}

	return cdom_match

}

func main() {
	//split the child domain string into words separated by ".". Find size of the child domain
	cdom_name = "www.bbc.co.uk"
	cdom_word = strings.SplitAfter(cdom_name, ".")

	c := cap(cdom_word)

	//combine child domain words in combinations starting from the end of the domain (uk, co.uk, bbc.co.uk) to compare with parent domain names
	for i := (cap(cdom_word) - 1); i >= 0; i-- {
		cdom_chk := strings.Join(cdom_word[i:c], "")

		cdom_match = compare(cdom_chk)
	}

	//find cap of longest matched child domain name. find diff between total cap of child domain and cap of matched domain
	cdom_match_word := strings.SplitAfter(cdom_match, ".")
	diff := (cap(cdom_word) - cap(cdom_match_word))

	//print plus one word as original child domain[diff], matched domain
	fmt.Println(cdom_word[diff-1], cdom_match)

}
