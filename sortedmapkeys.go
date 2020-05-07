package main

import (
	"fmt"
)

func main() {

	a := []int{1, 3, 2, 2, 4, 5, 6, 6, 6, 2, 3}
	m := make(map[int]int)
	mflip := make(map[int][]int)

	for i := 0; i < len(a); i++ {
		m[a[i]] = m[a[i]] + 1
	}
	fmt.Printf("m: %v\n", m)

	for i, v := range m {
		mflip[v] = append(mflip[v], i)
	}
	fmt.Printf("mflip: %v\n", mflip)

	keys := []int{}
	for i, _ := range mflip {
		keys = append(keys, i)
	}

	keys = MergeSort(keys)

	fmt.Printf("keys: %v\n", keys)

	m_new := make(map[int][]int)
	for i, v := range keys {
		m_new[v] = mflip[keys[i]]
		fmt.Printf("m_new: %v\n", m_new)
	}
}

func MergeSort(k []int) []int {
	if len(k) <= 1 {
		return k
	}
	size := len(k) / 2
	l := MergeSort(k[:size])
	r := MergeSort(k[size:])
	return Merge(l, r)
}

func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			ret = append(ret, r...)
			return ret
		}
		if len(r) == 0 {
			ret = append(ret, l...)
			return ret
		}
		if l[0] > r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}
