package chapter2

import (
	"reflect"
	"testing"
)

func TestRemoveMiddleNode(t *testing.T) {
	tests := []struct {
		input    []int
		removeAt int // index of node to remove (not head or tail)
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}}, // remove 3
		{[]int{1, 2, 3, 4}, 1, []int{1, 3, 4}},       // remove 2
		{[]int{1, 2, 3, 4}, 2, []int{1, 2, 4}},       // remove 3
	}

	for _, tt := range tests {
		list := buildList(tt.input)
		// Find node to remove (not head or tail)
		target := list
		for i := 0; i < tt.removeAt; i++ {
			target = target.Next
		}
		RemoveMiddleNode(target)
		got := listToSlice(list)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("RemoveMiddleNode(%v, removeAt=%d) = %v; want %v", tt.input, tt.removeAt, got, tt.expected)
		}
	}
}
