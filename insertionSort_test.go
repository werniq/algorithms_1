package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
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

	insertionSort(testCases[0].Input)
	insertionSort(testCases[1].Input)

	assert.Equal(t, testCases[0].Input, testCases[0].ExpectedOutput)
	assert.Equal(t, testCases[1].Input, testCases[1].ExpectedOutput)
}
