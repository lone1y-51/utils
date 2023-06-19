package math

import "github.com/lone1y-51/utils/define"

func Max[T define.Sorted](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func MaxArray[T define.Sorted](arr []T) (T, error) {
	if len(arr) <= 0 {
		return 0, ErrEmptyArray
	}
	max := arr[0]
	for _, ele := range arr {
		if ele > max {
			max = ele
		}
	}
	return max, nil
}

func Contains[T comparable](arr []T, target T) bool {
	for _, ele := range arr {
		if ele == target {
			return true
		}
	}
	return false
}

func ArrayEq[T comparable](arrA []T, arrB []T) bool {
	if len(arrA) != len(arrB) {
		return false
	}
	for index := range arrA {
		if arrA[index] != arrB[index] {
			return false
		}
	}
	return true
}
