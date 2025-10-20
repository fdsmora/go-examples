package chapter2

// CTCI problem 2.7 & Leetcode 160.
/*
	runtime: O(A+B)
	space: O(1)
*/
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type Result struct {
	Tail   *ListNode
	Length int
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var longest, shortest *ListNode = nil, nil

	resultA := getTailAndLength(headA)
	resultB := getTailAndLength(headB)

	if resultA.Tail != resultB.Tail {
		return nil
	}

	if resultA.Length > resultB.Length {
		longest = headA
		shortest = headB
	} else {
		shortest = headA
		longest = headB
	}

	longest = getToKthNode(longest, abs(resultA.Length-resultB.Length))

	for longest != shortest {
		longest = longest.Next
		shortest = shortest.Next
	}

	return longest
}

func getTailAndLength(list *ListNode) *Result {
	length := 1
	for list.Next != nil {
		list = list.Next
		length++
	}
	return &Result{list, length}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getToKthNode(n *ListNode, k int) *ListNode {
	for k > 0 {
		n = n.Next
		k--
	}
	return n
}

// my solution
/*
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    diff := 0
	var longest, other, pA, pB *ListNode = nil, nil, headA, headB

	for pA != nil && pB != nil {
		pA = pA.Next
		pB = pB.Next
	}

	if pA != nil || pB != nil {
		if pA == nil {
			longest = headB
			other = headA
			for pB != nil {
				pB = pB.Next
				diff++
			}
		} else if pB == nil {
			longest = headA
			other = headB
			for pA != nil {
				pA = pA.Next
				diff++
			}
		}

		// move longest to the node such that the remainder of both lists are same lenght
		for diff > 0 {
			longest = longest.Next
			diff--
		}
	} else {
		// lists are same size, longest and other can be any of the lists
		longest = headA
		other = headB
	}

	for longest != nil && other != nil {
		if longest == other {
			return longest
		}
		longest = longest.Next
		other = other.Next
	}
	return nil
}

*/

/*
A:       3->5->6->7->9
                     a
               o
B: 2->1->0->4->I
                     b
               l
diff 0
longest B
other   A
*/
