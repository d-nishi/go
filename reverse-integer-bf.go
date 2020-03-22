package main

import (
	"fmt"
	"math"
	"strings"
	"strconv"
)

func main() {
	x := -2147483648
	s := strconv.Itoa(x)
	a := strings.Split(s, "")
	sign := 0
	if a[0] == "-" {
		a = a[1:]
		sign = 1
	}
	a_out := []string{}

	for i := len(a) - 1; i >= 0; i = i - 1 {
		a_out = append(a_out, a[i])
	}

	s1, _ := strconv.Atoi(strings.Join(a_out, ""))

	limit := int(math.Pow(float64(2), float64(31)))

	if s1 < (-limit) || s1 > (limit-1) {
		s1 = 0
	} else if sign == 1 {
		s1 = (s1 * -1)
	}

	fmt.Printf("reverse integer for x:%d is:%d\n", x, s1)

}
