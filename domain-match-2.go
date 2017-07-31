package main

import "fmt"
import "strings"

var cdom_name string
var cdom_cap []string
var cdom_match []string

func main() {

	pdom := [][]string{
		[]string{"com"},
		[]string{"uk"},
		[]string{"oo.uk"},
		[]string{"gov.uk"},
		[]string{"nhs.uk"},
		[]string{"org.uk"},
		[]string{"ca.gov"},
		[]string{"cbc.co.uk"},
	}
	//Convert string to slice and find capacity of the child domain string
	cdom_name = "www.bbc.co.uk"
	cdom_cap = strings.Split(cdom_name, ".")

	//Split from the end of the child domain (uk, co.uk, bbc.co.uk) and compare to parent domain elements
	for i := cap(cdom_cap); i > 0; i-- {
		cdom := strings.SplitN(cdom_name, ".", i)
		for j := 0; j < cap(pdom); j++ {
			if pdom[j][0] != cdom[cap(cdom)-1] {
				continue
			}
			cdom_match = strings.Fields(pdom[j][0])
		}
	}

	//find the difference in length/capacity of full child domain slice and matched child domain slice
	diff := cap(cdom_cap) - cap(cdom_match)
	fmt.Println(diff)

	//Create a slice from combining child domain[diff-1] string and matched child domain string
	cdom_final := []string{cdom_cap[diff-1], strings.Join(cdom_match, ".")}

	//strings.Join final slice to generate the full string
	fmt.Println(strings.Join(cdom_final, "."))

}
