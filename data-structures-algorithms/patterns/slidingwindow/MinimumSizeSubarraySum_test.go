package slidingwindow

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		target   int
		nums     []int
		expected int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},       // [4,3]
		{15, []int{1, 2, 3, 4, 5}, 5},         // whole array
		{11, []int{1, 2, 3, 4, 5}, 3},         // [3,4,5]
		{4, []int{1, 4, 4}, 1},                // single element
		{100, []int{1, 2, 3, 4, 5}, 0},        // not possible
		{3, []int{1, 1, 1, 1, 1, 1, 1, 1}, 3}, // [1,1,1]
		{6, []int{10, 2, 3}, 1},               // first element
		{1, []int{}, 0},                       // empty array
		{1, []int{1}, 1},                      // single element
	}

	for _, tt := range tests {
		result := MinSubArrayLen(tt.target, tt.nums)
		if result != tt.expected {
			t.Errorf("MinSubArrayLen(%d, %v) = %d; want %d", tt.target, tt.nums, result, tt.expected)
		}
	}
}
