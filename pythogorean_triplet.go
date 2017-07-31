package main

import (
	"fmt"
	"os"
)

var a, b, c uint

func main() {

	for c = 1000; c > 0; c-- {
		for b = 1; b < c; b++ {
			a = (1000 - c - b)
			if a < b {
				if (a*a)+(b*b) == (c * c) {
					fmt.Println("Yes\n", a, b, c, (a*b*c))
					os.Exit(3)
				}
			}
		}
	}

}
