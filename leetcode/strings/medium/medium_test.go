package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPartitions(t *testing.T) {
	var tests = []struct {
		input  string
		output int
	}{
		{input: "32", output: 3},
		{input: "82734", output: 8},
		{input: "27346209830709182346", output: 9},
	}

	for i, test := range tests {
		assert.EqualValuesf(t, test.output, minPartitions(test.input), "test [%d] failed", i+1)
	}
}

func TestMinOperations(t *testing.T) {
	var tests = []struct {
		input  string
		output []int
	}{
		{input: "110", output: []int{1, 1, 3}},
		{input: "001011", output: []int{11, 8, 5, 4, 3, 4}},
	}

	for i, test := range tests {
		assert.EqualValuesf(t, test.output, minOperations(test.input), "test [%d] failed", i+1)
	}
}

func TestExecuteInstructions(t *testing.T) {
	var tests = []struct {
		size     int
		path     string
		startPos []int
		output   []int
	}{
		{size: 3, startPos: []int{0, 1}, path: "RRDDLU", output: []int{1, 5, 4, 3, 1, 0}},
		{size: 2, startPos: []int{1, 1}, path: "LURD", output: []int{4, 1, 0, 0}},
		{size: 1, startPos: []int{0, 0}, path: "LRUD", output: []int{0, 0, 0, 0}},
	}

	for i, test := range tests {
		res := executeInstructions(test.size, test.startPos, test.path)
		assert.EqualValuesf(t, test.output, res, "test [%d] failed", i+1)
	}
}

func TestPartitionLabels(t *testing.T) {
	var tests = []struct {
		input  string
		output []int
	}{
		{input: "ababcbacadefegdehijhklij", output: []int{9, 7, 8}},
		{input: "eccbbbbdec", output: []int{10}},
	}

	for i, test := range tests {
		assert.EqualValuesf(t, test.output, partitionLabels(test.input), "test [%d] failed", i+1)
	}
}
