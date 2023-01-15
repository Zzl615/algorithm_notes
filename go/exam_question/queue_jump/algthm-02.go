/**
 * @Author: Noaghzil
 * @Date:   2023-01-15 21:40:38
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-01-15 22:02:13
 */
package main

import "fmt"

func canJump(nums []int) bool {
	var max_key int
	length := len(nums) - 1

	for key, value := range nums {
		// fmt.Println("key:", key, "value:", value, "max_key", max_key)
		if key > max_key {
			return false
		}
		total := key + value
		if total > max_key {
			max_key = total
		}
		if length <= max_key {
			// fmt.Println("length:", length, "max_key", max_key)
			return true
		}
	}
	return false
}

func main() {
	testList := [][]int{
		// {2, 0, 0},
		// {2, 3, 1, 1, 4},
		// {3, 2, 1, 0, 4},
		// {1},
		// {0},
		// {0, 2, 3},
		{3, 0, 8, 2, 0, 0, 1},
	}
	for _, item := range testList {
		fmt.Println("canJump nums:", item)
		fmt.Println("canJump result:", canJump(item))
	}

}
