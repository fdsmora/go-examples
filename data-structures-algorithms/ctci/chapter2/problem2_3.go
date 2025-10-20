package chapter2

// CTCI problem 2.3
func RemoveMiddleNode(target *Node) {
	target.Value = target.Next.Value
	target.Next = target.Next.Next
}
