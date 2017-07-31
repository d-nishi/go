package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex, f float64) float64 {
	return (f * math.Sqrt(v.X*v.X+v.Y*v.Y))
	// return y
}

func Scale(v *Vertex, f float64) float64 {
	z := v.X*f + v.Y*f
	return z
}

func main() {
	v := Vertex{3, 4}
	t := Scale(&v, 10)
	r := Abs(v, 10)
	fmt.Println(v, t, r)
}
