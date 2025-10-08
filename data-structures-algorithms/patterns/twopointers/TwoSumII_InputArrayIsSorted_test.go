package twopointers

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},             // basic
		{[]int{1, 2, 3, 4, 4, 9, 56, 90}, 8, []int{4, 5}}, // duplicate values
		{[]int{5, 25, 75}, 100, []int{2, 3}},              // end of array
		{[]int{1, 2}, 3, []int{1, 2}},                     // minimal input
		{[]int{1, 2, 3, 4, 5}, 6, []int{1, 5}},            // first and last
		{[]int{1, 2, 3, 4, 5}, 7, []int{2, 5}},            // middle
		{[]int{-3, 1, 3, 4, 6, 8, 9, 21, 22, 25, 26}, 14, []int{5, 6}},
	}

	for _, tt := range tests {
		result := TwoSum(tt.numbers, tt.target)
		if len(result) != 2 || result[0] < 1 || result[1] < 1 || result[0] > len(tt.numbers) || result[1] > len(tt.numbers) {
			t.Errorf("TwoSum(%v, %d) returned invalid indices: %v", tt.numbers, tt.target, result)
			continue
		}
		if tt.numbers[result[0]-1]+tt.numbers[result[1]-1] != tt.target {
			t.Errorf("TwoSum(%v, %d) indices %v do not sum to target", tt.numbers, tt.target, result)
		}
		if result[0] != tt.expected[0] || result[1] != tt.expected[1] {
			t.Errorf("TwoSum(%v, %d) = %v; want %v", tt.numbers, tt.target, result, tt.expected)
		}
	}
}
