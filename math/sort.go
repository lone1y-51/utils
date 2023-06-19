package math

import (
	"github.com/lone1y-51/utils/define"
)

func MergeSort[T define.Sorted](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	return mergeSort(MergeSort(arr[:middle]), MergeSort(arr[middle:]))
}

func mergeSort[T define.Sorted](left []T, right []T) []T {
	leftIndex := 0
	rightIndex := 0
	var result []T
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

func QuickSort[T define.Sorted](arr []T) []T {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort[T define.Sorted](arr []T, left, right int) []T {
	if left < right {
		partitionIndex := partition(arr, left, right)
		quickSort(arr, left, partitionIndex-1)
		quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition[T define.Sorted](arr []T, left, right int) int {
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
