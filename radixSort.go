package main

import (
	"fmt"
	"strconv"
)

func sortRadix(arr []int) []int {
	m := max(arr)
	step := 1
	for i := 1; m/i > 0; i *= 10 {
		step++
	}

	radixSort(arr, step)

	return arr
}

func radixSort(arr []int, step int) {
	var strArr []string

	// Convert integers to fixed-length strings for comparison
	for _, v := range arr {
		str := fmt.Sprintf("%0*d", step, v)
		strArr = append(strArr, str)
	}

	for i := step - 1; i >= 0; i-- {
		findBiggestNumberByPosition(strArr, i)
	}

	// Convert the sorted strings back to integers
	for i, str := range strArr {
		num, _ := strconv.Atoi(str)
		arr[i] = num
	}
}

func findBiggestNumberByPosition(arr []string, pos int) {
	for i := 0; i < len(arr)-1; i++ {
		if len(arr[i]) <= pos || len(arr[i+1]) <= pos {
			// Handle strings with different lengths
			if len(arr[i]) < len(arr[i+1]) {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
			continue
		}

		if arr[i][pos] > arr[i+1][pos] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
}
