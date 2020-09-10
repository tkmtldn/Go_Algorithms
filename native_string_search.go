package main

import "fmt"

func nativeStrinSearch(text string, pattern string) []int {
	var pos []int
	for i := 0; i < len(text); i++ {
		var match bool = true
		for j := 0; j < len(pattern); j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			pos = append(pos, i)
		}
	}
	return pos
}

func main() {
	text := "My sweet home, only home sweet home"
	pattern := "sweet"
	var pos []int
	pos = nativeStrinSearch(text, pattern)
	if len(pos) ==0{
		fmt.Println("Pattern not found")
	} else {
		fmt.Println("Pattern found in current postion(positions):", pos)
	}
}