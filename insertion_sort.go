package main

import (
	"Go_Algorithms/random"
	"fmt"
)

func main() {
	arr := random.RandArray(10)
	fmt.Println("Before: ", arr)

	for x := 1; x < len(arr); x++ {
		y := x
		for y>0 && arr[y-1]>arr[y] {
			arr[y], arr[y-1] = arr[y-1],arr[y]
			y = y-1
		}
	}

	fmt.Println("After: ", arr)
}