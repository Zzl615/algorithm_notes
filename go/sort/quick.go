package sort

func QuickSort(nums []int, start int, end int) []int {
	pivot := nums[start]
	for start < end {
		for start < end && nums[end] >= pivot {
			end--
		}
		nums[start] = nums[end]

		for start < end && nums[start] <= pivot {
			start++
		}
		nums[end] = nums[start]
	}
	nums[start] = pivot
	QuickSort(nums, start+1, end)
	QuickSort(nums, 0, start-1)
	return nums

}
