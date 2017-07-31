package main

import "fmt"

func main() {

	S := 0
	f := []int{1, 2}
	i := 1

	for i >= 1 {
		x := f[i-1] + f[i]
		if x >= 4000000 {
			break
		}
		f = append(f, x)
		i = (i + 1)
	}

	for i := 0; i < len(f); i++ {
		if (f[i] % 2) == 0 {
			S = S + f[i]
		}
	}

	fmt.Println(S)

}
