/**
 * @Author: Noaghzil
 * @Date:   2023-01-15 18:18:23
 * @Last Modified by:   Noaghzil
 * 题目：
 * 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 判断你是否能够到达最后一个下标
 * @Last Modified time: 2023-01-15 21:29:34
 */

package main

import "fmt"

var dp map[int]bool

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	for key, value := range nums {
		if value <= 0 {
			return false
		}
		for i := 1; i <= value; i++ {
			// fmt.Println("canJump i:", key, i, value)
			next_key := key + i
			if value, ok := dp[next_key]; ok {
				// fmt.Println("canJump value")
				return value
			} else if next_key == len(nums)-1 {
				// fmt.Println("canJump true")
				return true
			} else {
				// fmt.Println("canJump next_key:", next_key, len(nums)-1, nums[next_key:len(nums)+0])
				if next_key < len(nums)-1 {
					result := canJump(nums[next_key : len(nums)+0])
					dp[key] = result
					// fmt.Println("canJump result")
					// return result
				} else {
					// fmt.Println("canJump false")
					return false
				}
			}
		}
	}
	return false
}

func main() {
	testList := [][]int{
		// {2, 0, 0},
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
		{1},
		{0},
		{0, 2, 3},
	}
	for _, item := range testList {
		dp = make(map[int]bool)
		fmt.Println("canJump nums:", item)
		fmt.Println("canJump result:", canJump(item))
	}

}
