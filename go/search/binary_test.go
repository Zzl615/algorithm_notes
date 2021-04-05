package search

import "testing"

func TestBinSearch(t *testing.T) {
	var nums = []int{2, 3, 5, 9, 12, 19, 20, 46, 50, 53, 66, 68, 77, 81, 89, 92, 98, 103, 107, 112}
	var result int
	var search int = 103
	index := BinSearch(nums, search)
	if index >= 0 && index < len(nums) {
		result = nums[index]
	}
	t.Log(index, search, result)
}

func TestBinSearchRecursive(t *testing.T) {
	var nums = []int{2, 3, 5, 9, 12, 19, 20, 46, 50, 53, 66, 68, 77, 81, 89, 92, 98, 103, 107, 112}
	var result int
	var search int = 1030
	index := BinSearchRecursive(nums, search, 0, len(nums)-1)
	if index >= 0 && index < len(nums) {
		result = nums[index]
	}
	t.Log(index, search, result)
}
