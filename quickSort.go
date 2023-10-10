package main

func partition(arr []int, start, end int) ([]int, int) {
	pivot := arr[end]

	i := start
	for j, v := range arr {
		if v < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	return arr, i
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
