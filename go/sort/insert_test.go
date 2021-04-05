package sort

import "testing"

func TestInsertSort(t *testing.T) {
	var nums = []int{1, 39, 2, 9, 7, 54, 11, 100, 29, 20}
	t.Log(nums)
	InsertSort(nums)
	t.Log(nums)
}
