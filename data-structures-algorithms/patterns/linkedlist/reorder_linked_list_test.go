package linkedlist

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		/* 		{"empty", []int{}, []int{}},
		   		{"single", []int{1}, []int{1}},
		   		{"two", []int{1, 2}, []int{1, 2}},
		   		{"three", []int{1, 2, 3}, []int{1, 3, 2}},
		   		{"four", []int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		   		{"five", []int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
		   		{"six", []int{1, 2, 3, 4, 5, 6}, []int{1, 6, 2, 5, 3, 4}}, */
		{"another", []int{2, 4, 6, 8, 10}, []int{2, 10, 4, 8, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			reorderList(head)
			got := listToSlice(head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("reorderList(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}
