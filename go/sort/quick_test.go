package sort

import "testing"


func TestQuickSort(t *testing.T) {
	list := []int{5, 9, 1, 9, 5, 3, 7, 6, 1}
	QuickSort(list, 0, len(list)-1)
	t.Log(list)
}
