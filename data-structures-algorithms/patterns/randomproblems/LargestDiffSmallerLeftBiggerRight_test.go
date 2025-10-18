package randomproblems

import (
	"testing"
)

func TestMaximumPositiveDifference(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			"fluctuate",
			[]int{5, 1, 3, 0, 4},
			4,
		},
		{
			"ups & downs",
			[]int{3, 1, 4, 2, 8},
			7,
		},
		{
			"strictly decreasing",
			[]int{5, 4, 3, 2, 1},
			0,
		},
		{
			"example input",
			[]int{2, 5, 1, 7},
			6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumPositiveDifference(tt.input)
			if got != tt.expect {
				t.Errorf("%s: maximumPositiveDifference(%v) = %d; want %d", tt.name, tt.input, got, tt.expect)
			}
		})
	}
}
