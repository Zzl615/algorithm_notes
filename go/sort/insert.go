package sort

import "fmt"

func InsertSort(nums []int) []int {
	var length int = len(nums)
	var temp int
	for i := 0; i < length; i++ {
		temp = nums[i]
		for j := i - 1; j >= 0; j-- {
			if temp < nums[j] {
				nums[j+1] = nums[j]
			} else {
				nums[j+1] = temp
				break
			}
		}
		fmt.Println(i, nums)
	}
	return nums
}
