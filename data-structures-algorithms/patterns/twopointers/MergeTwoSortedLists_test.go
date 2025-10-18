package twopointers

import (
	"reflect"
	"testing"
)

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1, l2   []int
		expected []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}}, // typical
		/* 	{[]int{}, []int{}, []int{}}, // both empty
		   	{[]int{}, []int{0}, []int{0}}, // one empty
		   	{[]int{2}, []int{1}, []int{1,2}}, // single elements
		   	{[]int{1,3,5}, []int{2,4,6}, []int{1,2,3,4,5,6}}, // interleaved
		   	{[]int{1,2,3}, []int{}, []int{1,2,3}}, // left only
		   	{[]int{}, []int{1,2,3}, []int{1,2,3}}, // right only
		   	{[]int{1,1,1}, []int{1,1,1}, []int{1,1,1,1,1,1}}, // all duplicates
		   	{[]int{-3,-2,-1}, []int{0,1,2}, []int{-3,-2,-1,0,1,2}}, // negatives and positives
		   	{[]int{1,2,3}, []int{4,5,6}, []int{1,2,3,4,5,6}}, // all left < right
		   	{[]int{4,5,6}, []int{1,2,3}, []int{1,2,3,4,5,6}}, // all right < left
		   	{[]int{1}, []int{1}, []int{1,1}}, // single duplicate
		   	{[]int{1,2,2}, []int{2,2,3}, []int{1,2,2,2,2,3}}, // overlapping duplicates */
	}

	for _, tt := range tests {
		l1 := buildList(tt.l1)
		l2 := buildList(tt.l2)
		merged := mergeTwoLists(l1, l2)
		result := listToSlice(merged)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("MergeTwoLists(%v, %v) = %v; want %v", tt.l1, tt.l2, result, tt.expected)
		}
	}
}
