package sortnsearch

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{2, 2, 2}, []int{2, 2, 2}},
		{[]int{10, -1, 2, 5, 0}, []int{-1, 0, 2, 5, 10}},
		{[]int{1, 3, 2, 3, 1}, []int{1, 1, 2, 3, 3}},
	}

	for _, tt := range tests {
		inputCopy := append([]int(nil), tt.input...)
		quickSort(inputCopy)
		if !reflect.DeepEqual(inputCopy, tt.expected) {
			t.Errorf("quickSort(%v) = %v; want %v", tt.input, inputCopy, tt.expected)
		}
	}
}
