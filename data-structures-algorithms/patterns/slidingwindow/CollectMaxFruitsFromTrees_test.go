package slidingwindow

import "testing"

func TestCollectMaxFruitTrees2Types(t *testing.T) {
	tests := []struct {
		name   string
		trees  []int
		expect int
	}{
		{
			name:   "basic example",
			trees:  []int{1, 2, 1},
			expect: 3,
		},
		{
			name:   "example with switch",
			trees:  []int{0, 1, 2, 2},
			expect: 3,
		},
		{
			name:   "all same type",
			trees:  []int{3, 3, 3, 3, 3},
			expect: 5,
		},
		{
			name:   "three types, max in middle",
			trees:  []int{1, 2, 3, 2, 2},
			expect: 4,
		},
		{
			name:   "alternating types",
			trees:  []int{1, 2, 1, 2, 1, 2},
			expect: 6,
		},
		{
			name:   "empty array",
			trees:  []int{},
			expect: 0,
		},
		{
			name:   "single element",
			trees:  []int{5},
			expect: 1,
		},
		{
			name:   "two elements, two types",
			trees:  []int{1, 2},
			expect: 2,
		},
		{
			name:   "window at end",
			trees:  []int{1, 2, 3, 4, 3, 3, 4},
			expect: 5,
		},
		{
			name:   "another",
			trees:  []int{1, 2, 1, 2, 3},
			expect: 4,
		},
		{
			name:   "repeated at start",
			trees:  []int{0, 0, 1, 1},
			expect: 4,
		},
		{
			name:   "...",
			trees:  []int{1, 2, 3, 2, 2},
			expect: 4,
		},
	}

	for _, tt := range tests {
		got := totalFruit(tt.trees)
		if got != tt.expect {
			t.Errorf("%s: CollectMaxFruitTrees2Types(%v) = %d; want %d", tt.name, tt.trees, got, tt.expect)
		}
	}
}
