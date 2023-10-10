package main

// 1. Sort in descending order (make a max binary heap)
// 2. Swap the largest (root) element with the last one  <-]
// 3. Remove last element								   |
// 4. Repeat ----------------------------------------------]

func buildMaxHeap(arr []int) {
	l := len(arr)
	for i := l/2 - 1; i >= 0; i-- {
		heapify(arr, l, i)
	}
}

func swap(arr []int, pos1, pos2 int) {
	arr[pos1], arr[pos2] = arr[pos2], arr[pos1]
}

func heapify(arr []int, l, i int) {
	largest := i
	start := 2*i + 1
	end := 2*i + 2

	if start < l && arr[start] > arr[largest] {
		largest = start
	}

	if end < l && arr[end] > arr[largest] {
		largest = end
	}

	if largest != i {
		swap(arr, i, largest)
		heapify(arr, l, largest)
	}
}

func heapSort(arr []int) []int {
	l := len(arr)

	buildMaxHeap(arr)

	for i := l - 1; i >= 0; i-- {
		swap(arr, 0, i)
		heapify(arr, i, 0)
	}

	return arr
}
