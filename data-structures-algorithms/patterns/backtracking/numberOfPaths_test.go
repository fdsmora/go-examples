package backtracking

import "testing"

// Problem: Count number of paths from (0,0) to (n-1, n-1) in an n×n grid
// Rules (typically):
// - Can only move East (right) or North (up)
// - Must stay on or below the diagonal (i >= j)
// - Start at (0,0), end at (n-1, n-1)

func TestNumOfPathsToDest(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "1x1 grid",
			n:    1,
			want: 1, // Only one cell, already at destination
		},
		{
			name: "2x2 grid",
			n:    2,
			want: 1, // Only one valid path: EE->N
		},
		{
			name: "3x3 grid",
			n:    3,
			want: 2, // Multiple valid paths
		},
		{
			name: "4x4 grid",
			n:    4,
			want: 5,
		},
		{
			name: "5x5 grid",
			n:    5,
			want: 14,
		},
		{
			name: "6x6 grid",
			n:    6,
			want: 42,
		},
		{
			name: "17x17 grid",
			n:    17,
			want: 35357670,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumOfPathsToDest(tt.n)
			if got != tt.want {
				t.Errorf("NumOfPathsToDest(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
