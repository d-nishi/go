package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	//var text string

	for scanner.Text() != "q" {
		fmt.Println("Enter your text:")
		scanner.Scan()
		if scanner.Text() != "q" {
			fmt.Println("Text entered was:", scanner.Text())
		}
	}
}
