package problem4_5

/**
 * Definition for a binary tree node. */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, _, isBST := isBinarySearchTree(root)
	return isBST
}

func isBinarySearchTree(root *TreeNode) (int, int, bool) {
	var greatestL, greatestR, smallestL, smallestR = root.Val, root.Val, root.Val, root.Val
	var isBST = true
	if root.Left != nil {
		greatestL, smallestL, isBST = isBinarySearchTree(root.Left)
		if !isBST || greatestL >= root.Val {
			return 0, 0, false
		}
	}
	if root.Right != nil {
		greatestR, smallestR, isBST = isBinarySearchTree(root.Right)
		if !isBST || smallestR <= root.Val {
			return 0, 0, false
		}
	}

	return greatestR, smallestL, isBST
}
