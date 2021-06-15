package math

func MergeSortInt(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	return mergeSortInt(MergeSortInt(arr[:middle]), MergeSortInt(arr[middle:]))
}

func mergeSortInt(left []int, right []int) []int {
	leftIndex := 0
	rightIndex := 0
	var result []int
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}
	if leftIndex < len(left) {
		result = append(result, left[leftIndex:]...)
	}
	if rightIndex < len(right) {
		result = append(result, right[rightIndex:]...)
	}
	return result
}

func MergeSortInt64(arr []int64) []int64 {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	return mergeSortInt64(MergeSortInt64(arr[:middle]), MergeSortInt64(arr[middle:]))
}

func mergeSortInt64(left []int64, right []int64) []int64 {
	leftIndex := 0
	rightIndex := 0
	var result []int64
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}
	if leftIndex < len(left) {
		result = append(result, left[leftIndex:]...)
	}
	if rightIndex < len(right) {
		result = append(result, right[rightIndex:]...)
	}
	return result
}
