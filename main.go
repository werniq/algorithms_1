package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Limit              = 591257512
	ElementLimit       = 1000
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

type srt func([]int) []int

func calc(fn srt, s int) int {
	results := []int64{}

	for i := 0; i < ResultsArrayLength; i++ {
		arr := generateRandomIntArray(s)

		startTime := time.Now().UnixNano()

		arr = fn(arr)

		duration := time.Now().UnixNano() - startTime

		results = append(results, duration)
	}

	return calculateAverageFromArray(results)
}

func generateRandomIntArray(s int) []int {
	arr := []int{}
	for i := 0; i < s; i++ {
		arr = append(arr, rand.Intn(Limit))
	}

	return arr
}

func compareAlgorithms() {
	fmt.Printf(`S Q I R H
`)

	for p := 10; p <= 10000; p *= 10 {
		for d := 1; d < 10; d++ {
			s := d * p

			avQuickSortTime := calc(func(arr []int) []int { return quickSort(arr, 0, s-1) }, s)
			avInstSortTime := calc(insertionSort, s)
			avRadixSortTime := calc(sortRadix, s)
			avHeapSortTime := calc(heapSort, s)

			fmt.Printf(`%d %d %d %d %d
`,
				s,
				avQuickSortTime,
				avInstSortTime,
				avRadixSortTime,
				avHeapSortTime,
			)
		}
	}

}

func main() {
	compareAlgorithms()
}
