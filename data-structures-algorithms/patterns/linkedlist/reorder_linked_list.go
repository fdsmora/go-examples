package linkedlist

import "fmt"

func reorderList(head *ListNode) {
	n := getLength(head)

	half := n / 2
	/* 	if n%2 == 1 {
		half += 1
	} */

	c := head

	for i := 0; i < half; i++ {
		c = c.Next
	}

	tmp := c.Next
	c.Next = nil
	c = tmp

	r := reverseList(c)

	t := r
	for t != nil {
		fmt.Printf("%d ->", t.Val)
		t = t.Next
	}
	fmt.Println()

	l := head

	dummy := &ListNode{}
	c = dummy

	for l != nil && r != nil {
		c.Next = l
		l = l.Next
		c = c.Next
		c.Next = r
		r = r.Next
		c = c.Next
	}

	for l != nil {
		c.Next = l
		c = c.Next
		l = l.Next
	}
	for r != nil {
		c.Next = r
		c = c.Next
		r = r.Next
	}
	head = dummy.Next
}

func getLength(head *ListNode) int {
	c := head
	count := 0
	for c != nil {
		count++
		c = c.Next
	}
	return count
}
