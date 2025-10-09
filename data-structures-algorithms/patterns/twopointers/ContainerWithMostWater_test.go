package twopointers

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{1,8,6,2,5,4,8,3,7}, 49}, // classic case
		{[]int{1,1}, 1},                 // minimal input
		{[]int{4,3,2,1,4}, 16},          // equal ends
		{[]int{1,2,1}, 2},               // small peak
		{[]int{2,3,4,5,18,17,6}, 17},    // tall peak
		{[]int{1,2,3,4,5}, 6},           // increasing
		{[]int{5,4,3,2,1}, 6},           // decreasing
		{[]int{0,0,0,0}, 0},             // all zero
		{[]int{1}, 0},                   // single element
		{[]int{}, 0},                    // empty
	}

	for _, tt := range tests {
		result := MaxArea(tt.height)
		if result != tt.expected {
			t.Errorf("MaxArea(%v) = %d; want %d", tt.height, result, tt.expected)
		}
	}
}
