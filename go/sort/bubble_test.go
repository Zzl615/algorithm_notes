package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	var nums = []int64{1, 39, 2, 9, 7, 54, 11}
	BubbleSort(nums)
	t.Log(nums)
}

func TestPositiveBubbleSort(t *testing.T) {
	var nums = []int64{1, 2, 7, 9, 11, 39, 54}
	BubbleSort(nums)
	t.Log(nums)
}

// TODO：优化逆序集合
func TestReverseBubbleSort(t *testing.T) {
	var nums = []int64{54, 39, 11, 9, 7, 2, 1}
	BubbleSort(nums)
	t.Log(nums)
}
