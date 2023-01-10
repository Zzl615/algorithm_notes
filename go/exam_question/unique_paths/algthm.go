/**
 * @Author: Noaghzil
 * @Date:   2023-01-07 22:35:59
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-01-07 23:59:29
 */
package main

import "fmt"

var result [][]int{
	
}

func uniquePaths(m int, n int) int {
	var totalNumber int
	if m == n && m == 1 {
		totalNumber = 1
	} else {
		var rightNumber int = 0
		var downNumber int = 0
		if m > 1 {
			if n > 1 {
				rightNumber = uniquePaths(m, n-1)
			} else {
				rightNumber = 1
			}
		}
		if n > 1 {
			if m > 1 {
				downNumber = uniquePaths(m-1, n)
			} else {
				downNumber = 1
			}
		}
		totalNumber = rightNumber + downNumber
	}
	return totalNumber
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
