package trees

// LC 572
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	Q := []*TreeNode{root}

	for len(Q) > 0 {
		c := Q[0]
		Q = Q[1:]
		if c == nil {
			continue
		}
		Q = append(Q, c.Left)
		Q = append(Q, c.Right)
		if checkEqual(c, subRoot) {
			return true
		}
	}
	return false
}

func checkEqual(root, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	return root != nil && subRoot != nil &&
		root.Val == subRoot.Val &&
		checkEqual(root.Left, subRoot.Left) &&
		checkEqual(root.Right, subRoot.Right)
}
