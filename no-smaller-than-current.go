package main

import (
   "fmt"
)

func main() {

    nums := []int{8,1,2,3,2,4}
    //sort the list with O(nlog(n))
    nums_sorted := MergeSort(nums)

    fmt.Printf("Sorted list:%d\n", nums_sorted)
    fmt.Println()

    m := make(map[int]int)
    for i:=0; i<len(nums_sorted); i++ {
        _, ok := m[nums_sorted[i]]
        if ok == false {
            m[nums_sorted[i]] = i
        } else {
            m[nums_sorted[i]] = m[nums_sorted[i]]
        }
    }

    nums_out := []int{}
    for i:=0; i<len(nums); i++ {
        nums_out = append(nums_out, m[nums[i]])
    }

    fmt.Printf("Numbers in the current array: %d \n", nums)
    fmt.Println()
    fmt.Printf("Numbers smaller than current no: %d \n", nums_out)

}

func MergeSort(s []int) []int {
    if len(s) <= 1 {
        return s
    }
    n := len(s)/2
    l := MergeSort(s[:n])
    r := MergeSort(s[n:])

    return Merge(l, r)
}

func Merge(l, r []int) []int {
    ret := []int {}
    for len(l) > 0 || len(r) > 0 {
        if len(l) == 0 {
            return append(ret, r...)
        }
        if len(r) == 0 {
            return append(ret, l...)
        }
        if l[0] < r[0] {
            ret = append(ret, l[0])
            l = l[1:]
        } else {
            ret = append(ret, r[0])
            r = r[1:]
        }
    }
    return ret
}
