package chapter1

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
		valid    bool
	}{
		// 3x3 matrix
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
			valid: true,
		},
		// 4x4 matrix
		{
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expected: [][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
			valid: true,
		},
		// 1x1 matrix
		{
			input:    [][]int{{1}},
			expected: [][]int{{1}},
			valid:    true,
		},
		// empty matrix
		{
			input:    [][]int{},
			expected: [][]int{},
			valid:    false,
		},
		// non-square matrix
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		inputCopy := make([][]int, len(tt.input))
		for i := range tt.input {
			inputCopy[i] = append([]int(nil), tt.input[i]...)
		}
		ok := RotateMatrix(inputCopy)
		if ok != tt.valid {
			t.Errorf("RotateMatrix(%v) validity = %v; want %v", tt.input, ok, tt.valid)
		}
		if ok && !reflect.DeepEqual(inputCopy, tt.expected) {
			t.Errorf("RotateMatrix(%v) = %v; want %v", tt.input, inputCopy, tt.expected)
		}
	}
}
