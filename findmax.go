package main

import "fmt"

func main() {

	a := []int{10, 20, 30, 50, 100, 200, 75, 25}
	max := findMax(0, len(a), a)
	fmt.Printf("max: %v\n", max)
}

func findMax(si int, ei int, a []int) int {

	for i := si; i < ei; i++ {

		m := si + ((ei - si) / 2)

		if m == 0 {
		}

		if m == 1 {
			max := a[0]
			for j, _ := range a {
				if max < a[j] {
					max = a[j]
				}
			}
			return max
		}

		if a[m] < a[m+1] {
			if a[m] < a[m-1] {
				fmt.Printf("elem: %v is lesser than the elem after: %v and elem before: %v\n", a[m], a[m+1], a[m-1])
				return 0
			}
			if a[m] > a[m-1] {
				return findMax(m, ei, a)
			}
		}
		if a[m] > a[m+1] {
			if a[m] < a[m-1] {
				return findMax(si, m, a)
			}
			if a[m] > a[m-1] {
				return a[m]
			}
		}

	}
	return 0
}
