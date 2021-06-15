package math

import "testing"

func arrayEqInt(arrA, arrB []int) bool {
	if len(arrA) != len(arrB) {
		return false
	}
	for i := range arrA {
		if arrA[i] != arrB[i] {
			return false
		}
	}
	return true
}

func arrayEqInt64(arrA, arrB []int64) bool {
	if len(arrA) != len(arrB) {
		return false
	}
	for i := range arrA {
		if arrA[i] != arrB[i] {
			return false
		}
	}
	return true
}

func Test_MergeSortInt(t *testing.T) {
	testArr := []int{2, 5, 1, 4, 10, 8, 3}
	mergeResult := MergeSortInt(testArr)
	result := []int{1, 2, 3, 4, 5, 8, 10}
	if !arrayEqInt(result, mergeResult) {
		t.Errorf("result %v != mergeResult %v", result, mergeResult)
	}
}

func Test_MergeSortInt64(t *testing.T) {
	testArr := []int64{2, 5, 1, 4, 10, 8, 3}
	mergeResult := MergeSortInt64(testArr)
	result := []int64{1, 2, 3, 4, 5, 8, 10}
	if !arrayEqInt64(result, mergeResult) {
		t.Errorf("result %v != mergeResult %v", result, mergeResult)
	}
}
