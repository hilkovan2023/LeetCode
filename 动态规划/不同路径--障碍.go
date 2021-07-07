package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	fmt.Println(m, n)
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}

	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}

	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	for i, _ := range dp {
		for j, _ := range dp[i] {
			fmt.Printf("%.2v,", dp[i][j])
		}
		fmt.Println()
	}
	return dp[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}

	res := uniquePathsWithObstacles(obstacleGrid)

	fmt.Println(res)
}
