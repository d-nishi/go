package main

import "fmt"
import "math"

func Sqrt(x float64) float64 {
	S := 1.0
	for c := 1.0; c < (x / 2); c++ {
		S0 := S
		S = (S*S + x) / (2 * S)
		if S0 == S {
			return (S0)
		}
	}
	return S
}

func main() {
	fmt.Println(Sqrt(18.0))
	fmt.Println(math.Sqrt(18.0))
}
