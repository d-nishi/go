package main

import (
	"fmt"
)

func main() {

	k_sorted_arrs := [][]int{
		[]int{1, 2, 5},
		[]int{3, 6, 9, 12, 31},
	}

	fmt.Println(MergeKSortedArrs(k_sorted_arrs))

}

func MergeKSortedArrs(k_sorted_arrs [][]int) []int {

	k0 := &k_sorted_arrs[0]
	k1 := &k_sorted_arrs[1]
	temp_elems := []int{}
	out := []int{}
	i, j := 0, 0

	l := len(*k0) + len(*k1)

	for l > 0 {
		temp_elems = append(temp_elems, (*k0)[i])
		temp_elems = append(temp_elems, (*k1)[j])

		if (*k0)[i] < (*k1)[j] {
			out = append(out, (*k0)[i])
			i = i + 1
		} else if (*k0)[i] > (*k1)[j] {
			out = append(out, (*k1)[j])
			j = j + 1
		} else {
			out = append(out, (*k0)[i], (*k1)[j])
			i = i + 1
			j = j + 1
		}

		if i == (len(*k0)) {
			out = append(out, (*k1)[j:]...)
			break
		} else if j == (len(*k1)) {
			out = append(out, (*k0)[i:]...)
			break
		}
		l = (l - 1)
	}

	return out
}
