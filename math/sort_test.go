package math

import "testing"

func Test_MergeSort(t *testing.T) {
	testArr := []int{2, 5, 1, 4, 10, 8, 3}
	mergeResult := MergeSort(testArr)
	result := []int{1, 2, 3, 4, 5, 8, 10}
	if !ArrayEq(result, mergeResult) {
		t.Errorf("result %v != mergeResult %v", result, mergeResult)
	}
}

func Test_QuickSort(t *testing.T) {
	testArr := []int{2, 5, 1, 4, 10, 8, 3}
	quickResult := QuickSort(testArr)
	result := []int{1, 2, 3, 4, 5, 8, 10}
	if !ArrayEq(result, quickResult) {
		t.Errorf("result %v != mergeResult %v", result, quickResult)
	}
}
