package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	var nums = []int64{1, 39, 2, 9, 7, 54, 11}
	BubbleSort(nums)
	t.Log(nums)
}
