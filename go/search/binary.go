package search

import "fmt"

// nums: 有序数列，从小到大
func BinSearch(nums []int, obj int) int {
	start := 0
	end := len(nums)
	var mid int
	// 循环条件：需要等于，两边收缩到只剩两数[16,17]，只能检查其中一个
	for start <= end {
		mid = (start + end) / 2
		fmt.Println(start, end)
		if nums[mid] == obj {
			return mid
		} else if nums[mid] < obj {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1

}

func BinSearchRecursive(nums []int, obj int, start int, end int) int {
	// 终止条件
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	fmt.Println(start, end, mid, nums[mid])
	if obj == nums[mid] {
		return mid
	} else if obj > nums[mid] {
		return BinSearchRecursive(nums, obj, mid+1, end)
	} else {
		return BinSearchRecursive(nums, obj, start, mid-1)
	}
}
