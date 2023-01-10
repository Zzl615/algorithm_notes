/**
 * @Author: Noaghzil
 * @Date:   2023-01-08 21:46:45
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-01-08 21:47:21
 */
package main

import "fmt"

func uniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	networkMap := [][]int{
		{2, 2},
		{3, 2},
		{3, 3},
		{3, 7},
	}
	for _, one := range networkMap {
		fmt.Println("one of map:", one)
		number := uniquePaths(one[0], one[1])
		fmt.Println("number is", number)
	}

}
