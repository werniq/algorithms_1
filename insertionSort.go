package main

/*
	i ← 1
	while i < length(A)
		j ← i
		while j > 0 and A[j-1] > A[j]
			swap A[j] and A[j-1]
			j ← j - 1
		end while
		i ← i + 1
	end while
*/

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Shift elements of arr[0..i-1] that are greater than key
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

	return arr
}
