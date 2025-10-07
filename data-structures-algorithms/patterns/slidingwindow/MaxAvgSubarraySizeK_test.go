package slidingwindow

import (
	"math"
	"testing"
)

func TestMaxAvgSubarraySizeK(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		expect float64
	}{
		{
			name:   "basic positive",
			nums:   []int{1, 12, -5, -6, 50, 3},
			k:      4,
			expect: 12.75,
		},
		{
			name:   "all negatives",
			nums:   []int{-1, -2, -3, -4},
			k:      2,
			expect: -1.5,
		},
		{
			name:   "single element",
			nums:   []int{5},
			k:      1,
			expect: 5.0,
		},
		{
			name:   "window size equals array",
			nums:   []int{2, 4, 6, 8},
			k:      4,
			expect: 5.0,
		},
		{
			name:   "window size larger than array",
			nums:   []int{1, 2},
			k:      3,
			expect: math.Inf(-1),
		},
		{
			name:   "empty array",
			nums:   []int{},
			k:      1,
			expect: math.Inf(-1),
		},
	}

	for _, tt := range tests {
		got := MaxAvgSubarraySizeK(tt.nums, tt.k)
		if math.Abs(got-tt.expect) > 1e-6 {
			t.Errorf("%s: MaxAvgSubarraySizeK(%v, %d) = %v; want %v", tt.name, tt.nums, tt.k, got, tt.expect)
		}
	}
}
