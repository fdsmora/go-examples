package chapter2

import (
	"reflect"
	"testing"
)

func buildList(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	head := &Node{Value: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &Node{Value: v}
		curr = curr.Next
	}
	return head
}

func listToSlice(head *Node) []int {
	var res []int = []int{}
	for head != nil {
		res = append(res, head.Value)
		head = head.Next
	}
	return res
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 1, 1}, []int{1}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 2, 3, 3, 4, 4}, []int{2, 3, 4}},
	}

	for _, tt := range tests {
		list := buildList(tt.input)
		result := RemoveDuplicates(list)
		got := listToSlice(result)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("RemoveDuplicates(%v) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}

func TestRemoveDuplicatesNoHash(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 1, 1}, []int{1}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 2, 3, 3, 4, 4}, []int{2, 3, 4}},
	}

	for _, tt := range tests {
		list := buildList(tt.input)
		result := RemoveDuplicatesNoHash(list)
		got := listToSlice(result)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("RemoveDuplicatesNoHash(%v) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
