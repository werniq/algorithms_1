package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Limit        = 591257512
	ElementLimit = 1000000
)

func compareAlgorithms() {
	var testCase []int

	for i := 0; i < ElementLimit; i++ {
		testCase = append(testCase, rand.Intn(Limit))
	}

	arr := testCase

	quickSortStartTime := time.Now().UnixMilli()

	arr1 := quickSort(arr, 0, len(testCase)-1)
	quickSortTimeDuration := time.Now().UnixMilli() - quickSortStartTime

	arr = testCase

	insertionSortStartTime := time.Now().UnixMilli()

	arr2 := insertionSort(arr)

	insertionSortDurationTime := time.Now().UnixMilli() - insertionSortStartTime

	arr = testCase

	heapSortStartTime := time.Now().UnixMilli()
	arr3 := heapSort(arr)

	heapSortDurationTime := time.Now().UnixMilli() - heapSortStartTime

	arr = testCase

	radixSortTime := time.Now().UnixMilli()
	arr4 := sortRadix(arr)

	fmt.Println(arr1 == nil)
	fmt.Println(arr2 == nil)
	fmt.Println(arr3 == nil)
	fmt.Println(arr4 == nil)

	radixSortDurationTime := time.Now().UnixMilli() - radixSortTime

	fmt.Println("Duration of insertion sort algorithm: ", insertionSortDurationTime)
	fmt.Println("Duration of quick sort algorithm: ", quickSortTimeDuration)
	fmt.Println("Duration of heap sort algorithm: ", heapSortDurationTime)
	fmt.Println("Duration of radix sort algorithm: ", radixSortDurationTime)
}

func main() {
	compareAlgorithms()
}
