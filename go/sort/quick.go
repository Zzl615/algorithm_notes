package sort

func QuickSort(nums []int, start int, end int) []int {
	// 递归结束条件
	if start > end {
		return nums
	}
	// 分治的中枢
	pivot := nums[start]
	// 保留该轮的起止值
	low := start
	high := end
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
	QuickSort(nums, low, start-1)
	QuickSort(nums, start+1, high)
	return nums
}
