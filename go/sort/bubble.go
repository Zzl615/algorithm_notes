package sort

import "fmt"

func BubbleSort(nums []int64) []int64 {	
	for i := 0; i < len(nums)-1; i++ {
		for j := i; j < len(nums)-1; j++ {
			fmt.Println(i, j)
			if j != len(nums)-1 && nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	return nums
}
