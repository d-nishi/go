package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 0}
	index := []int{0, 1, 2, 3, 0}

	out := []int{}

	for i := 0; i < len(nums); i++ {
		out = append(out, 0)
		copy(out[index[i]+1:], out[index[i]:])
		out[index[i]] = nums[i]
	}

	fmt.Printf("out: %v\n", out)
}
