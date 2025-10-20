package chapter2

type Node struct {
	Value int
	Next  *Node
}

// CTCI problem 2.1
// Runtime: O(n)
// Space: O(n) because of the hashtable
func RemoveDuplicates(n *Node) *Node {
	if n == nil || n.Next == nil {
		return n
	}
	var prev *Node
	curr := n
	dup := make(map[int]bool)
	for curr != nil {
		if dup[curr.Value] {
			prev.Next = curr.Next
		} else {
			dup[curr.Value] = true
			prev = curr
		}
		curr = curr.Next
	}
	return n
}

// works but more complicated
/* func RemoveDuplicatesNoHash(list *Node) *Node {
	var curr, runner *Node = list, nil
	for curr != nil {
		runner = curr
		for runner != nil && runner.Next != nil {
			if runner.Next.Value == curr.Value {
				runner2 := runner.Next
				for runner2 != nil && runner2.Value == curr.Value {
					runner2 = runner2.Next
				}
				runner.Next = runner2
			}
			runner = runner.Next
		}
		curr = curr.Next
	}
	return list
} */

// Runtime: O(n(n-1)/2) because runner first iterates over n elements, then n-1, n-2...1 -> O(n^2)
// Space: O(1)
func RemoveDuplicatesNoHash(list *Node) *Node {
	var curr, runner *Node = list, nil
	for curr != nil {
		runner = curr
		for runner.Next != nil {
			if runner.Next.Value == curr.Value {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		curr = curr.Next
	}
	return list
}
