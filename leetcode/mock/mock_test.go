package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalSum(t *testing.T) {
	var tests = []struct {
		inputMat [][]int
		output   int
	}{
		{
			inputMat: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			output: 25,
		},
		{
			inputMat: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			output: 8,
		},
		{
			inputMat: [][]int{{5}},
			output:   5,
		},
	}

	for i, test := range tests {
		assert.Equalf(t, test.output, diagonalSum(test.inputMat), "test[%d] failed", i+1)
	}
}

func TestDieSimulator(t *testing.T) {
	var tests = []struct {
		input   int
		rollMax []int
		output  int
	}{
		{
			input:   2,
			rollMax: []int{1, 1, 2, 2, 2, 3},
			output:  34,
		},
		{
			input:   2,
			rollMax: []int{1, 1, 1, 1, 1, 1},
			output:  30,
		},
		{
			input:   3,
			rollMax: []int{1, 1, 1, 2, 2, 3},
			output:  181,
		},
	}
	for i, test := range tests {
		assert.Equalf(t, test.output, dieSimulator(test.input, test.rollMax), "test[%d] failed", i+1)
	}
}

func TestCanFormArray(t *testing.T) {
	var tests = []struct {
		arr    []int
		pieces [][]int
		output bool
	}{
		{
			arr:    []int{15, 88},
			pieces: [][]int{{88}, {15}},
			output: true,
		},
		{
			arr:    []int{49, 18, 16},
			pieces: [][]int{{16, 18, 49}},
			output: false,
		},
		{
			arr:    []int{91, 4, 64, 78},
			pieces: [][]int{{78}, {4, 64}, {91}},
			output: true,
		},
	}
	for i, test := range tests {
		assert.Equalf(t, test.output, canFormArray(test.arr, test.pieces), "test[%d] failed", i+1)
	}
}

func TestCanConstruct(t *testing.T) {
	var tests = []struct {
		s      string
		k      int
		output bool
	}{
		{
			s:      "annabelle",
			k:      2,
			output: true,
		},
		{
			s:      "leetcode",
			k:      3,
			output: false,
		},
		{
			s:      "true",
			k:      4,
			output: true,
		},
	}
	for i, test := range tests {
		assert.Equalf(t, test.output, canConstruct(test.s, test.k), "test[%d] failed", i+1)
	}
}
