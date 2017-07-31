package main

import "fmt"
import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {

	fmt.Println(RandSeq(10))

}
