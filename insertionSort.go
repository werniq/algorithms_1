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

func insertionSort(arr []int, n int) {
	for i := 1; i <= len(arr)-1; i++ {
		key := arr[i]
		j := 0

		// we are moving all elements bigger than key
		// one position forward
		for j = i; j > 0 && arr[j-1] > key; j = j {
			arr[j-1] = arr[j]
			j = j - 1
		}

		arr[j-1] = key
	}
}
