package main

import (
  "fmt"
  "strings"
  "os"
)

func main() {

    strs := []string{"flower","flow","flight"}

    m := make(map[int]string)
    if len(strs) == 0 {
	fmt.Printf("longest common prefix : %d\n",m)
	os.Exit(3)
    }

    a0 := strings.Split(strs[0],"")
    fmt.Println()
    fmt.Printf("first array : %s\n", a0)

    for i:=0; i< len(a0); i++ {
        elem, ok := m[i]
        if ok == false {
            m[i] = a0[i]
        } else {
            fmt.Println(elem)
            continue
        }
    }
    fmt.Printf("first lcp : %v\n", m)

    for i:=1; i< len(strs); i++ {
        a1 := strings.Split(strs[i],"")
        fmt.Println("====================")
        fmt.Printf("next array : %s\n", a1)
        m = longestCommonPrefix(a1, m)
        fmt.Printf("next lcp : %v\n",m)
    }

    arr := []string{}
    for i:=0; i<len(m); i++ {
        arr = append(arr, m[i])
    }
    fmt.Println(strings.Join(arr,""))
}

func longestCommonPrefix(a []string, m map[int]string) (map[int]string) {
    m1 := make(map[int]string)
    l := 0
    if len(a) < len(m) {
        l = len(a)
    } else {
        l = len(m)
    }

    for i:=0; i<l; i++ {
        if m[i] == a[i] {
            m1[i] = m[i]
            continue
        } else {
            break
        }
    }
    return m1
}
