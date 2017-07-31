package main

import "fmt"

func main() {

	sum := 0
	sumsq := 0
	i := 100
	sumsq = i * (i + 1) * (2*i + 1) / 6
	sum = i * (i + 1) / 2

	fmt.Println(((sum * sum) - sumsq))
}
