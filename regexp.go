package main

import (
	"fmt"
	"regexp"
)

func main() {

	re := regexp.MustCompile(`[    ]`)

	value := "Mr. John Smith   "

	result := re.Split(value, -1)

	for i := range result {
		fmt.Println(result[i])
	}

}
