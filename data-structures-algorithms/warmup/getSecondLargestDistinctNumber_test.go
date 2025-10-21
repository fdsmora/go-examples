package patterns

import "testing"

func TestGetSecondLargestDistinctNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"ascending", []int{1, 2, 3, 4, 5}, 4},
		{"all equal", []int{5, 5, 5, 5}, 5}, // no distinct second -> returns largest
		{"two elements", []int{10, 9}, 9},
		{"negative numbers", []int{-3, -1, -2, -4}, -2},
		{"largest repeated", []int{2, 5, 5, 3}, 3},
		{"single element", []int{7}, 7},
		{"duplicates with second", []int{9, 9, 8, 8}, 8},
		{"descending", []int{10, 9, 8, 7}, 9},
		{"mixed", []int{1, 3, 2, 3}, 2},
		{"large values", []int{1000000000, 999999999, 1000000000}, 999999999},
	}

	for _, tt := range tests {
		got := getSecondLargestDistinctNumber(tt.nums)
		if got != tt.want {
			t.Errorf("%s: got %d want %d for input %v", tt.name, got, tt.want, tt.nums)
		}
	}
}
