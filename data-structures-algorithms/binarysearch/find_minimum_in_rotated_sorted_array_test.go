package binarysearch

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"rotated", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"not rotated", []int{1, 2, 3, 4, 5}, 1},
		{"single element", []int{1}, 1},
		{"two elements rotated", []int{2, 1}, 1},
		{"two elements not rotated", []int{1, 2}, 1},
		{"min at end", []int{2, 3, 4, 5, 1}, 1},
		{"small rotation", []int{3, 4, 5, 1, 2}, 1},
		{"negatives rotated", []int{-1, 0, 1, -3, -2}, -3},
		{"another rotated", []int{10, 11, 12, 5, 6, 7, 8, 9}, 5},
		{"another", []int{3, 4, 5, 6, 1, 2}, 1},
		{"another", []int{4, 5, 0, 1, 2, 3}, 0},
	}

	for _, tt := range tests {
		if got := findMin(tt.nums); got != tt.want {
			t.Errorf("%s: findMin(%v) = %d; want %d", tt.name, tt.nums, got, tt.want)
		}
	}
}
