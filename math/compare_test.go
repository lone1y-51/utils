package math

import (
	"testing"
)

func Test_Max(t *testing.T) {
	a, b := 2, 3
	result := Max(a, b)
	if result != b {
		t.Errorf("max %d result %d", b, result)
	}
	c, d := int64(2), int64(3)
	result2 := Max(c, d)
	if result2 != d {
		t.Errorf("max %d result %d", b, result)
	}
}

func Test_MaxArray(t *testing.T) {
	arr := []int{3, 6, 7, 1, 10, 45, 80, -2, 0}
	result, err := MaxArray(arr)
	if err != nil || result != 80 {
		t.Errorf("err %v, result %d, max %d", err, result, 80)
	}
	arr2 := []int64{3, 6, 7, 1, 10, 45, 80, -2, 0}
	result2, err := MaxArray(arr2)
	if err != nil || result2 != 80 {
		t.Errorf("err %v, result %d, max %d", err, result2, 80)
	}
}

func Test_Contains(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	result := Contains(arr, 4)
	if result != true {
		t.Errorf("contains err result %t, contains %t", result, true)
	}
	arr2 := []int64{1, 2, 3, 4, 5, 6, 7}
	result2 := Contains(arr2, 4)
	if result2 != true {
		t.Errorf("contains err result %t, contains %t", result2, true)
	}
	arr3 := []string{"a", "b", "c"}
	result3 := Contains(arr3, "c")
	if result3 != true {
		t.Errorf("contains err result %t, contains %t", result3, true)
	}
}

func Test_ArrayEq(t *testing.T) {
	arrA := []int{1, 2, 3, 4, 5}
	arrB := []int{1, 2, 3, 4, 5}
	result := ArrayEq(arrA, arrB)
	if result != true {
		t.Errorf("arrayEq err result %t, eq %t", result, true)
	}
	arr2A := []int64{1, 2, 3, 4, 5}
	arr2B := []int64{1, 2, 3, 4, 5}
	result2 := ArrayEq(arr2A, arr2B)
	if result2 != true {
		t.Errorf("arrayEq err result %t, eq %t", result2, true)
	}
	arr3A := []string{"1", "2", "3"}
	arr3B := []string{"1", "2", "3"}
	result3 := ArrayEq(arr3A, arr3B)
	if result3 != true {
		t.Errorf("arrayEq err result %t, eq %t", result3, true)
	}
}
