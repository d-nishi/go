package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	str := "teblack"
	stra := strings.Split(str, "")
	m := make(map[string]int)

	for _, v := range stra {
		elem, ok := m[v]
		if ok == false {
			m[v] = 1
			continue
		}
		m[v] = (elem + 1)
		fmt.Println("Not Unique")
		os.Exit(3)
	}
fmt.Println(m,"Unique")

}
