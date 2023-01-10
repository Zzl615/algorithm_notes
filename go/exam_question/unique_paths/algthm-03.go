/**
 * @Author: Noaghzil
 * @Date:   2023-01-08 21:49:37
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-01-08 21:49:37
 */

package main

import (
	"fmt"
	"math/big"
)

func uniquePaths(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
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
