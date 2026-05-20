package trees

import "fmt"

var printf = fmt.Printf

// LeetCode problem 98. Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/
// Medium
// Time complexity: O(n) where n is the number of nodes in the tree, since we need to visit each node once.
// Space complexity: O(h) where h is the height of the tree, due to the recursive call stack. In the worst case of a completely unbalanced tree, this could be O(n).
func isValidBST(node *TreeNode) bool {
	isValid, _, _ := visit(node)
	return isValid
}

func visit(node *TreeNode) (bool, int, int) {
	if node.Left == nil && node.Right == nil {
		// debug
		//	printf("leaf node, val: %d\n", node.Val)
		return true, node.Val, node.Val
	}

	if node.Right == nil {
		isValid, smallest, biggest := visit(node.Left)
		//printf("node val: %d, left is valid? %b, greatest:%d\n", node.Val, isValid, greatest)
		if !isValid || biggest >= node.Val {
			return false, 0, 0
		}
		return true, smallest, node.Val
	}

	if node.Left == nil {
		isValid, smallest, biggest := visit(node.Right)
		//	printf("node val: %d, right is valid? %b, smallest:%d\n", node.Val, isValid, smallest)
		if !isValid || smallest <= node.Val {
			return false, 0, 0
		}
		return true, node.Val, biggest
	}

	isValidL, smallestL, biggestL := visit(node.Left)
	isValidR, smallestR, biggestR := visit(node.Right)

	if isValidL && isValidR && biggestL < node.Val && smallestR > node.Val {
		//printf("val: %d, biggest: %d, smallest: %d\n", node.Val, biggest, smallest)
		return true, smallestL, biggestR
	}
	return false, 0, 0
}
