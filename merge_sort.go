package main

import (
	"Go_Algorithms/random"
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	middle := len(arr) / 2
	part_1 := mergeSort(arr[:middle])
	part_2 := mergeSort(arr[middle:])
	return merge(part_1, part_2)
}

func merge(p1, p2 []int) []int {
	ans := make([]int, len(p1)+len(p2))
	i := 0
	j := 0
	for i < len(p1) && j < len(p2) {
		if p1[i] <= p2[j] {
			ans[i+j] = p1[i]
			i++
		} else {
			ans[i+j] = p2[j]
			j++
		}
	}
	for i < len(p1) {
		ans[i+j] = p1[i]
		i++
	}
	for j < len(p2) {
		ans[i+j] = p2[j]
		j++
	}
	return ans
}

func main() {
	arr := random.RandArray(10)
	fmt.Println("Before: ", arr)
	fmt.Println("After: ", mergeSort(arr))
}
