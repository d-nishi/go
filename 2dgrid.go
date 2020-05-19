package main

import (
	"fmt"
)

func main() {

	stock_price := []int{10, 5, 6, 9, 100, 50, 156, 65}

	profit := make([][]int, len(stock_price))

	for i := 0; i < len(stock_price); i++ {
		profit[i] = make([]int, len(stock_price))
		for j := 1; j < len(stock_price); j++ {
			profit[i][j] = stock_price[j] - stock_price[i]
		}
	}

	fmt.Println(profit)

}
