package problem4_8

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	if l != nil {
		r := lowestCommonAncestor(root.Right, p, q)
		if r != nil {
			return root
		}
		return l
	}
	return lowestCommonAncestor(root.Right, p, q)
}
