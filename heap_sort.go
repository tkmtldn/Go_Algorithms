package main

import (
	"Go_Algorithms/random"
	"fmt"
)

type MaxHeap struct {
	slice []int
	size  int
}

func (h MaxHeap) getSize() int {
	return h.size
}

func BuildMaxHeap(slice []int) MaxHeap {
	h := MaxHeap{slice: slice, size: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return h
}

func (h MaxHeap) MaxHeapify(i int) {
	l, r := 2*i+1, 2*i+2
	max := i

	if l < h.getSize() && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < h.getSize() && h.slice[r] > h.slice[max] {
		max = r
	}
	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.MaxHeapify(max)
	}
}

func heapSort(slice []int) []int{
	h := BuildMaxHeap(slice)
	for i := len(h.slice)-1; i>0; i--{
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.size --
		h.MaxHeapify(0)
	}
	return h.slice
}

func main(){
	arr := random.RandArray(10)
	fmt.Println("Before: ", arr)
	h := BuildMaxHeap(arr)
	fmt.Println("Initial heap: ", h.slice)
	fmt.Println("After: ", heapSort(arr))
}
