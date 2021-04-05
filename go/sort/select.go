package sort

import "fmt"

func SelectSort(nums []int) []int {
	var length int = len(nums)
	var min int
	for i := 0; i < length; i++ {
		min = i
		for j := i + 1; j < length; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
		fmt.Println(i, min, nums)
	}
	return nums
}
