package main

import (
	"fmt"
	"math/rand"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func RodCutting(price []int, length int) int {
	memoArray := make([]int, length+1)
	memoArray[0] = 0
	for j := 1; j <= length; j++ {
		q := -1
		for i := 1; i <= j; i++ {
			q = max(q, price[i]+memoArray[j-i])
		}
		memoArray[j] = q
	}
	return memoArray[length]
}

func main() {
	length := 10
	price := make([]int, length+1)
	for i := 0; i < length+1; i++ {
		price[i] = rand.Intn(20)
	}
	fmt.Println(price)
	fmt.Println(RodCutting(price, length))
}
