package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfUnique(t *testing.T) {
	var tests = []struct {
		input  []int
		output int
	}{
		{input: []int{1, 2, 3, 2}, output: 4},
		{input: []int{1, 1, 1, 1, 1}, output: 0},
		{input: []int{1, 2, 3, 4, 5}, output: 15},
		{input: []int{10, 6, 9, 6, 9, 6, 8, 7}, output: 25},
	}

	for i, test := range tests {
		assert.EqualValuesf(t, test.output, sumOfUnique(test.input), "test [%d] failed", i+1)
	}
}
