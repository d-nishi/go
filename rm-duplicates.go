package main

import (
  "fmt"
)

func main() {

   nums := []int{0,0,1,1,1,2,3,4,4}
   fmt.Println(removeDuplicates(nums))

}

func removeDuplicates(nums []int) int {
    m := make(map[int]int)

    for i:=0; i<len(nums); i++ {
        _, ok := m[nums[i]]
        if ok == false {
            m[nums[i]] = 1
        } else {
            if i != len(nums)-1 {
                nums = append(nums[:i], nums[i+1:]...)
                i = i-1
            } else {
                nums = nums[:i]
                break
            }
        }
    }

    fmt.Println(m)
    fmt.Println(nums)

    return len(nums)
}
