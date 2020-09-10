package main

import (
	"Go_Algorithms/random"
	"fmt"
	"math/rand"
)

func QuickSort(arr []int) []int {

	if len(arr) == 0 {
		return arr
	}

	choice := arr[rand.Intn(len(arr))]

	low := make([]int, 0, len(arr))
	med := make([]int, 0, len(arr))
	high := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < choice:
			low = append(low, item)
		case item == choice:
			med = append(med, item)
		case item > choice:
			high = append(high, item)
		}
	}

	low = QuickSort(low)
	high = QuickSort(high)

	low = append(low, med...)
	low = append(low, high...)

	return low
}

func main() {
	arr := random.RandArray(10)
	fmt.Println("Before: ", arr)
	fmt.Println("After: ", QuickSort(arr))
}
