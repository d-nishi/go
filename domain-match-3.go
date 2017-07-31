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
		[]string{"co.uk"},
		[]string{"gov.uk"},
		[]string{"nhs.abc.uk"},
		[]string{"trupanion.org.uk"},
		[]string{"ca.gov"},
		[]string{"bbc.co.uk"},
	}

	//create map based on string length for each domain entry
	m1 := make(map[string]string)
	for i := range pdom {
		x := strings.Split(pdom[i][0], ".")
		v, ok := m1[x[len(x)-1]]
		fmt.Println(v)
		if ok == true {
			len := len(x)
			y := strings.SplitN(pdom[i][0], ".", (len - 1))
			m1[y[len-2]] = pdom[i][0]
		}
		m1[x[len(x)-1]] = pdom[i][0]
	}
	fmt.Println(pdom)

	fmt.Println(m1)

	cdom_name = "www.bbc.co.uk"
	cdom_cap = strings.Split(cdom_name, ".")

}
