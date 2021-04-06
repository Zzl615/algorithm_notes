package sort

import "testing"

func TestHillSort(t *testing.T) {
	nums := []int{1, 39, 2, 9, 7, 54, 11, 100, 29, 20}
	t.Log(nums)
	HillSort(nums)
	t.Log(nums)
}
