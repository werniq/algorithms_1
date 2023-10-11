package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		Input          []int
		ExpectedOutput []int
	}{
		{
			[]int{20, 9, 15, -100, 400},
			[]int{-100, 9, 15, 20, 400},
		},
		{
			Input:          []int{80, 70, 60, 50, 40, 30, 20, 10},
			ExpectedOutput: []int{10, 20, 30, 40, 50, 60, 70, 80},
		},
	}

	quickSort(testCases[0].Input, 0, len(testCases[0].Input)-1)
	quickSort(testCases[1].Input, 0, len(testCases[1].Input)-1)

	assert.Equal(t, testCases[0].Input, testCases[0].ExpectedOutput)
	assert.Equal(t, testCases[1].Input, testCases[1].ExpectedOutput)

}
