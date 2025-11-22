package trees

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func inOrder(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	inOrder(t.Left)
	fmt.Println(t.Val)
	inOrder(t.Right)
	return nil
}

func preOrder(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	fmt.Println(t.Val)
	preOrder(t.Left)
	preOrder(t.Right)
	return nil
}

func posOrder(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	posOrder(t.Left)
	posOrder(t.Right)
	fmt.Println(t.Val)
	return nil
}
