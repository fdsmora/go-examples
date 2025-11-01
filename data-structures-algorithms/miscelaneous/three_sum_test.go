package miscelaneous

import (
	"reflect"
	"sort"
	"testing"
)

func normalizeTriples(in [][]int) [][]int {
	out := make([][]int, 0, len(in))
	for _, t := range in {
		s := append([]int(nil), t...)
		sort.Ints(s)
		out = append(out, s)
	}
	// sort the list of triples lexicographically
	sort.Slice(out, func(i, j int) bool {
		for x := 0; x < 3; x++ {
			if out[i][x] != out[j][x] {
				return out[i][x] < out[j][x]
			}
		}
		return false
	})
	return out
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"classic", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"all zeros", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"no solution", []int{1, 2, -2, -1}, [][]int{}},
		{"duplicates", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{"empty", []int{}, [][]int{}},
		{"small", []int{0, 1}, [][]int{}},
		{"zeros many", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"negatives and positives", []int{-3, -1, 0, 1, 2}, [][]int{{-3, 1, 2}, {-1, 0, 1}}},
		{"mixed duplicates", []int{-1, -1, 0, 1, 1}, [][]int{{-1, 0, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			ngot := normalizeTriples(got)
			nwant := normalizeTriples(tt.want)
			if !reflect.DeepEqual(ngot, nwant) {
				t.Errorf("threeSum(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}
