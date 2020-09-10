package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func matrixChainDp(m []int) int {
	matrixLength := len(m)
	dp := make([][]int, matrixLength)

	for i := 0; i < matrixLength; i++ {
		dp[i] = make([]int, matrixLength)
		dp[i][i] = 0
	}

	for j := 2; j < matrixLength; j++ {
		for x := 1; x < matrixLength-j+1; x++ {
			y := x + j - 1
			dp[x][y] = 1 << 31
			for k := x; k < y; k++ {
				prod := dp[x][k] + dp[k+1][y] + m[x-1]*m[k]*m[y]
				dp[x][y] = min(prod, dp[x][y])
			}
		}
	}
	return dp[1][matrixLength-1]
}
func main() {
	matrix := []int{3,2,2,3}
	fmt.Println(matrixChainDp(matrix))
}