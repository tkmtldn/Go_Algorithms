package main

import (
	"Go_Algorithms/random"
	"fmt"
)

func bubbleSort(arr []int) []int {
	step := 0
Loop:
	for i := 0; i < len(arr)-1; i++ {
		move := 0
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				move++
			}
			step ++

		}
		if move == 0 {
			break Loop
		}
	}
	fmt.Printf("Finished in %v steps.\n", step)
	return arr
}

func main() {
	arr := random.RandArray(11)
	fmt.Println(arr)
	fmt.Println(bubbleSort(arr))

	arr2 := []int{0, 2, 1, 3, 4, 5, 6, 7, 8, 10, 9}
	fmt.Println(bubbleSort(arr2))
}
