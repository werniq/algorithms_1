package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Limit              = 591257512
	ElementLimit       = 10000000
	ResultsArrayLength = 10
)

func calculateAverageFromArray(arr []int64) int {
	sum := 0

	count := 0
	for _, v := range arr {
		sum += int(v)
		count++
	}

	return sum / count
}

func calculateAverageOfQuickSortAlgorithm() int {
	results := []int64{}

	for i := 0; i < ResultsArrayLength; i++ {
		arr := generateRandomIntArray()

		startTime := time.Now().UnixMilli()

		arr = quickSort(arr, 0, len(arr)-1)

		duration := time.Now().UnixMilli() - startTime

		results = append(results, duration)
	}

	return calculateAverageFromArray(results)
}

func calculateAverageOfHeapSortAlgorithm() int {
	results := []int64{}

	for i := 0; i < ResultsArrayLength; i++ {
		arr := generateRandomIntArray()

		startTime := time.Now().UnixMilli()

		arr = heapSort(arr)

		duration := time.Now().UnixMilli() - startTime

		results = append(results, duration)
	}

	return calculateAverageFromArray(results)
}

func calculateAverageOfInsertionSortAlgorithm() int {
	results := []int64{}

	for i := 0; i < ResultsArrayLength; i++ {
		arr := generateRandomIntArray()

		startTime := time.Now().UnixMilli()

		arr = insertionSort(arr)

		duration := time.Now().UnixMilli() - startTime

		results = append(results, duration)
	}

	return calculateAverageFromArray(results)
}

func calculateAverageOfRadixSortAlgorithm() int {
	results := []int64{}

	for i := 0; i < ResultsArrayLength; i++ {
		arr := generateRandomIntArray()

		startTime := time.Now().UnixMilli()

		arr = sortRadix(arr)

		duration := time.Now().UnixMilli() - startTime

		results = append(results, duration)
	}

	return calculateAverageFromArray(results)
}

func generateRandomIntArray() []int {
	arr := []int{}
	for i := 0; i < ElementLimit; i++ {
		arr = append(arr, rand.Intn(Limit))
	}

	return arr
}

func compareAlgorithms() {
	var testCase []int

	for i := 0; i < ElementLimit; i++ {
		testCase = append(testCase, rand.Intn(Limit))
	}

	avQuickSortTime := calculateAverageOfQuickSortAlgorithm()
	avRadixSortTime := calculateAverageOfRadixSortAlgorithm()
	avHeapSortTime := calculateAverageOfHeapSortAlgorithm()
	avInstSortTime := calculateAverageOfInsertionSortAlgorithm()

	fmt.Printf(`
			Average amount of time, was needed to sort 100 arrays of %d elements (in milliseconds):
			Quick sort: %d
			Insertion sort: %d
			Radix sort: %d
			Heap sort: %d
		
		`,
		ElementLimit,
		avQuickSortTime,
		avInstSortTime,
		avRadixSortTime,
		avHeapSortTime,
	)
}

func main() {
	compareAlgorithms()
}
