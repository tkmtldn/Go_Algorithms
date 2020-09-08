package main

import (
	"Go_Algorithms/random"
	"fmt"
)

func main() {
	arr := random.RandArray(10)
	fmt.Println("Before: ", arr)
	fmt.Println("After: ", mergeSort(arr))
}
