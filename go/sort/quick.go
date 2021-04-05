package sort

// 交换排序

func QuickSort(nums []int, start int, end int) []int {
	// 递归结束条件
	if start > end {
		return nums
	}
	// 基准: 空出start
	pivot := nums[start]
	// 保留该轮的起止值
	low := start
	high := end
	for start < end {
		for start < end && nums[end] >= pivot {
			end--
		}
		// 填补start，空出end
		nums[start] = nums[end]

		for start < end && nums[start] <= pivot {
			start++
		}
		// 填补end，空出start
		nums[end] = nums[start]
	}
	// 使用pivot，填补最后空位
	nums[start] = pivot
	QuickSort(nums, low, start-1)
	QuickSort(nums, start+1, high)
	return nums
}
