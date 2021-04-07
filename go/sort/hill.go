package sort

import "fmt"

func HillSort(nums []int) []int {
	var length int = len(nums)
	var gap int = 1
	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := nums[i]
			for j := i - gap; j >= 0; j = j - gap {
				if nums[i] < nums[j] {
					nums[j+gap] = nums[j]
				} else {
					nums[j+gap] = temp
					break
				}
			}
		}
		fmt.Println(gap, nums)
		gap = gap / 3
	}
	return nums
}
