package math

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxInt64(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MaxIntArray(arr []int) (int, error) {
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

func MaxInt64Array(arr []int64) (int64, error) {
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

func ContainsInt(arr []int, target int) bool {
	for _, ele := range arr {
		if ele == target {
			return true
		}
	}
	return false
}

func ContainsInt64(arr []int64, target int64) bool {
	for _, ele := range arr {
		if ele == target {
			return true
		}
	}
	return false
}

func ContainsString(arr []string, target string) bool {
	for _, ele := range arr {
		if ele == target {
			return true
		}
	}
	return false
}

func ArrayIntEq(arrA []int, arrB []int) bool {
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

func ArrayInt64Eq(arrA []int64, arrB []int64) bool {
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

func ArrayStringEq(arrA []string, arrB []string) bool {
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
