package main

import "fmt"

func main() {

	sa := []string{"H", "o", "w", "", "a", "r", "e", "", "y", "o", "u", "", "t", "o", "d", "a", "y", "?"}
	fmt.Println(sa)

	i0 := 0

	for i := 0; i < cap(sa); i++ {
		if sa[i] == "" {
			fmt.Println(sa[i0:i])
			i0 = (i + 1)
		}
	}

}
