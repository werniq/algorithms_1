package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Limit        = 591257512
	ElementLimit = 20
)

func compareAlgorithms() {
	var testCase []int

	for i := 0; i < ElementLimit; i++ {
		testCase = append(testCase, rand.Intn(Limit))
	}

	quickSortStartTime := time.Now().UnixMilli()

	arr1 := quickSort(testCase, 0, len(testCase)-1)
	quickSortTimeDuration := time.Now().UnixMilli() - quickSortStartTime

	fmt.Println("Duration of quick sort algorithm: ", quickSortTimeDuration)
	fmt.Println(arr1)

	insertionSortStartTime := time.Now().UnixMilli()
	arr2 := insertionSort(testCase)

	insertionSortDurationTime := time.Now().UnixMilli() - insertionSortStartTime

	fmt.Println("Duration of insertion sort algorithm: ", insertionSortDurationTime)
	fmt.Println(arr2)

	heapSortStartTime := time.Now().UnixMilli()
	arr3 := heapSort(testCase)

	heapSortDurationTime := time.Now().UnixMilli() - heapSortStartTime
	fmt.Println("Duration of heap sort algorithm: ", heapSortDurationTime)
	fmt.Println(arr3)

}

func main() {
	compareAlgorithms()
}
