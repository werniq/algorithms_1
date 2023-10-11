package main

func partition(arr []int, start, end int) ([]int, int) {
	pivot := arr[end]

	i := start - 1
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[end] = arr[end], arr[i+1]

	return arr, i + 1
}

func quickSort(arr []int, leftBound, rightBound int) []int {
	if leftBound < rightBound {
		var p int
		arr, p = partition(arr, leftBound, rightBound)
		arr = quickSort(arr, leftBound, p-1)
		arr = quickSort(arr, p+1, rightBound)
	}

	return arr
}
