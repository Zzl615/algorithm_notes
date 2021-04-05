package sort

import "testing"

func TestSelectSort(t *testing.T) {
	var nums = []int{1, 39, 2, 9, 7, 54, 11, 100, 29, 20}
	t.Log(nums)
	SelectSort(nums)
	t.Log(nums)
}
