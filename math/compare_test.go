package math

import (
	"testing"
)

func Test_MaxInt(t *testing.T) {
	a, b := 2, 3
	result := MaxInt(a, b)
	if result != b {
		t.Errorf("max %d result %d", b, result)
	}
}

func Test_MaxInt64(t *testing.T) {
	a, b := int64(2), int64(3)
	result := MaxInt64(a, b)
	if result != b {
		t.Errorf("max %d result %d", b, result)
	}
}

func Test_MaxIntArray(t *testing.T) {
	arr := []int{3, 6, 7, 1, 10, 45, 80, -2, 0}
	result, err := MaxIntArray(arr)
	if err != nil || result != 80 {
		t.Errorf("err %v, result %d, max %d", err, result, 80)
	}
}

func Test_MaxInt64Array(t *testing.T) {
	arr := []int64{3, 6, 7, 1, 10, 45, 80, -2, 0}
	result, err := MaxInt64Array(arr)
	if err != nil || result != 80 {
		t.Errorf("err %v, result %d, max %d", err, result, 80)
	}
}

func Test_ContainsInt(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	result := ContainsInt(arr, 4)
	if result != true {
		t.Errorf("contains err result %t, contains %t", result, true)
	}
}

func Test_ContainsInt64(t *testing.T) {
	arr := []int64{1, 2, 3, 4, 5, 6, 7}
	result := ContainsInt64(arr, 4)
	if result != true {
		t.Errorf("contains err result %t, contains %t", result, true)
	}
}

func Test_ContainsString(t *testing.T) {
	arr := []string{"a", "b", "c"}
	result := ContainsString(arr, "c")
	if result != true {
		t.Errorf("contains err result %t, contains %t", result, true)
	}
}

func Test_ArrayIntEq(t *testing.T) {
	arrA := []int{1, 2, 3, 4, 5}
	arrB := []int{1, 2, 3, 4, 5}
	result := ArrayIntEq(arrA, arrB)
	if result != true {
		t.Errorf("arrayEq err result %t, eq %t", result, true)
	}
}

func Test_ArrayInt64Eq(t *testing.T) {
	arrA := []int64{1, 2, 3, 4, 5}
	arrB := []int64{1, 2, 3, 4, 5}
	result := ArrayInt64Eq(arrA, arrB)
	if result != true {
		t.Errorf("arrayEq err result %t, eq %t", result, true)
	}
}

func Test_ArrayStringEq(t *testing.T) {
	arrA := []string{"1", "2", "3"}
	arrB := []string{"1", "2", "3"}
	result := ArrayStringEq(arrA, arrB)
	if result != true {
		t.Errorf("arrayEq err result %t, eq %t", result, true)
	}
}
