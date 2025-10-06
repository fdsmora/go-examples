package chapter1

import (
	"reflect"
	"testing"
)

func TestSetToZeros(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		// No zeros
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		// One zero in middle
		{
			input: [][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 0, 3},
				{0, 0, 0},
				{7, 0, 9},
			},
		},
		// Multiple zeros
		{
			input: [][]int{
				{0, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 5, 0},
				{0, 0, 0},
			},
		},
		// All zeros
		{
			input: [][]int{
				{0, 0},
				{0, 0},
			},
			expected: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		// 1x1 matrix zero
		{
			input:    [][]int{{0}},
			expected: [][]int{{0}},
		},
		{
			input:    [][]int{{1, 0}, {3, 4}},
			expected: [][]int{{0, 0}, {3, 0}},
		},
	}

	for _, tt := range tests {
		inputCopy := make([][]int, len(tt.input))
		for i := range tt.input {
			inputCopy[i] = append([]int(nil), tt.input[i]...)
		}
		SetToZeros(inputCopy)
		if !reflect.DeepEqual(inputCopy, tt.expected) {
			t.Errorf("SetToZeros(%v) = %v; want %v", tt.input, inputCopy, tt.expected)
		}
	}
}

func TestSetToZerosBitVector(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		// No zeros
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		// One zero in middle
		{
			input: [][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 0, 3},
				{0, 0, 0},
				{7, 0, 9},
			},
		},
		// Multiple zeros
		{
			input: [][]int{
				{0, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 5, 0},
				{0, 0, 0},
			},
		},
		// All zeros
		{
			input: [][]int{
				{0, 0},
				{0, 0},
			},
			expected: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		// 1x1 matrix zero
		{
			input:    [][]int{{0}},
			expected: [][]int{{0}},
		},
		{
			input:    [][]int{{1, 0}, {3, 4}},
			expected: [][]int{{0, 0}, {3, 0}},
		},
	}

	for _, tt := range tests {
		inputCopy := make([][]int, len(tt.input))
		for i := range tt.input {
			inputCopy[i] = append([]int(nil), tt.input[i]...)
		}
		SetToZerosBitVector(inputCopy)
		if !reflect.DeepEqual(inputCopy, tt.expected) {
			t.Errorf("SetToZerosBitVector(%v) = %v; want %v", tt.input, inputCopy, tt.expected)
		}
	}
}
