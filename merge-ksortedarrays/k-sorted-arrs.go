package main

import "fmt"

func main() {
	k_sorted_arrs := [][]int{
		[]int{1, 2, 5},
		[]int{3, 7},
		[]int{8, 9, 14, 30},
	}

	fmt.Printf("OutputArr: %v\n", MergeKSortedArrs(k_sorted_arrs))
}

func MergeKSortedArrs(k_sorted_arrs [][]int) []int {

	kindex := []int{}
	l := 0

	for i := 0; i < len(k_sorted_arrs); i++ {
		kindex = append(kindex, 0)
		l = l + len(k_sorted_arrs[i])*(i+1)
	}

	i := 0
	j := 0
	min := k_sorted_arrs[i][kindex[i]]
	out := []int{}
	mout := make(map[int]int)
	arrsread := 0

	for l > 0 {

		if kindex[i] != -1 {
			if min > k_sorted_arrs[i][kindex[i]] {
				min = k_sorted_arrs[i][kindex[i]]
				j = i
			}

			if i == len(k_sorted_arrs)-1 {
				kindex[j] = kindex[j] + 1
				out = append(out, min)
				mout[min] = mout[min] + 1
				fmt.Println(out)

				for k, _ := range kindex {
					if kindex[k] == len(k_sorted_arrs[k]) && mout[k_sorted_arrs[k][kindex[k]-1]] >= 1 {
						kindex[k] = -1
						arrsread = arrsread + 1
					}
				}

				if arrsread != len(kindex)-1 {
					i = arrsread
					j = arrsread
					if kindex[i] == -1 {
						for k, _ := range kindex {
							if kindex[k] != -1 {
								i = k
								break
							}
						}
						j = i
						fmt.Printf("test continue i:%d\n", i)
						min = k_sorted_arrs[i][kindex[i]]
						continue
					}
				} else {
					for k, _ := range kindex {
						if kindex[k] != -1 {
							i = k
							break
						}
					}
					out = append(out, k_sorted_arrs[i][kindex[i]:]...)
					break
				}

				min = k_sorted_arrs[i][kindex[i]]
			}

			i = i + 1
			l = l - 1

		} else {
			i = i + 1
		}

	}

	fmt.Printf("OutputArr: %d\n", out)
	return out

}
