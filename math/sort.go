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

func QuickSortInt(arr []int) []int {
	return quickSortInt(arr, 0, len(arr)-1)
}

func quickSortInt(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partitionInt(arr, left, right)
		quickSortInt(arr, left, partitionIndex-1)
		quickSortInt(arr, partitionIndex+1, right)
	}
	return arr
}

func partitionInt(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}
