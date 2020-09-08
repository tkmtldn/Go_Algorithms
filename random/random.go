package random

import "math/rand"

func RandArray(n int)[]int  {
	ans := make([]int, n)
	for i:=0; i<n; i++{
		ans[i] = rand.Intn(n*5)
	}
	return ans
}