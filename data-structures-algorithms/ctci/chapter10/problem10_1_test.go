package chapter10

import (
	"reflect"
	"testing"
)

func TestSortArrays(t *testing.T) {
	tests := []struct {
		A      []int
		B      []int
		expect []int
	}{
		// Example from comments
		{[]int{3, 3, 4, 5, 7, 8, 9, 11, 13, 15, 19, 0, 0, 0, 0}, []int{2, 4, 7, 19}, []int{2, 3, 3, 4, 4, 5, 7, 7, 8, 9, 11, 13, 15, 19, 19}},
		// Both arrays empty
		//{[]int{}, []int{}, []int{}},
		// B empty
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		// A empty (all zeros)
		{[]int{0, 0, 0}, []int{1, 2, 3}, []int{1, 2, 3}},
		// Interleaved
		{[]int{1, 3, 5, 0, 0, 0}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		// All B less than A
		{[]int{5, 6, 7, 0, 0, 0}, []int{1, 2, 3}, []int{1, 2, 3, 5, 6, 7}},
		// All A less than B
		{[]int{1, 2, 3, 0, 0, 0}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		// Duplicates
		{[]int{2, 2, 2, 0, 0}, []int{2, 2}, []int{2, 2, 2, 2, 2}},
		// A and B single element
		{[]int{1, 0}, []int{2}, []int{1, 2}},
		// A single, B empty
		{[]int{1}, []int{}, []int{1}},
		// A empty, B single
		{[]int{0}, []int{2}, []int{2}},
		// All zeros in A, B empty
		{[]int{0, 0, 0}, []int{}, []int{0, 0, 0}},
		// B longer than A
		{[]int{5, 0, 0, 0}, []int{1, 2, 3}, []int{1, 2, 3, 5}},
		// A and B both negative
		{[]int{-3, -2, -1, 0, 0, 0}, []int{-6, -5, -4}, []int{-6, -5, -4, -3, -2, -1}},
		// Mixed negative and positive
		{[]int{-2, 0, 0}, []int{1, 2}, []int{-2, 1, 2}},
	}

	for _, tt := range tests {
		A := append([]int(nil), tt.A...)
		countA := len(tt.A) - len(tt.B)
		countB := len(tt.B)
		sortArrays(A, tt.B, countA, countB)
		if !reflect.DeepEqual(A, tt.expect) {
			t.Errorf("sortArrays(%v, %v, %d, %d) = %v; want %v", tt.A, tt.B, countA, countB, A, tt.expect)
		}
	}
}
