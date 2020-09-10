package main

import "fmt"

func binarySearch(arr []int, q int) int {
	if len(arr) == 1 {
		if arr[0] == q {
			return q
		}
		return -1
	}
	mid := len(arr) / 2
	if q==arr[mid]{
		return q
	} else if q<arr[mid]{
		return binarySearch(arr[0:mid], q)
	} else {
		return binarySearch(arr[mid+1:], q)
	}
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(arr, 0))
	fmt.Println(binarySearch(arr, 11))
	}