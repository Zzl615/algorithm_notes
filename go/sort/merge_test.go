package sort

import "testing"

func TestMergeSort(t *testing.T) {
	nums := []int{1, 39, 2, 9, 7, 54, 11, 100, 29, 20}
	t.Log(nums)
	MergeSort(nums)
	t.Log(nums)
}
