package trees

import (
	"fmt"
	"math"
)

/*
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if (p != nil && q == nil) || (p==nil && q!=nil) {
        return false
    }
    return (p==nil && q == nil) ||
        (isSameTree(p.Left, q.Left) && p.Val == q.Val && isSameTree(p.Right, q.Right))
}*/

// LeetCode 100
/*
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}
	return (p == nil && q == nil) ||
		(isSameTree(p.Left, q.Left) && p.Val == q.Val && isSameTree(p.Right, q.Right))
}*/

func isSameTree(p, q *TreeNode) bool {
	chP := make(chan int)
	chQ := make(chan int)
	doneP, doneQ := make(chan bool), make(chan bool)
	go func() {
		isSameTreeX(p, chP)
		doneP <- true
	}()
	go func() {
		isSameTreeX(q, chQ)
		doneQ <- true
	}()

	lP := []int{}
	lQ := []int{}

	var dP, dQ bool
	done := make(chan bool)
	go func() {
		for {
			select {
			case vP := <-chP:
				lP = append(lP, vP)
			case vQ := <-chQ:
				lQ = append(lQ, vQ)
			case dP = <-doneP:
				continue
			case dQ = <-doneQ:
				continue
			default:
				if dP && dQ {
					done <- true
					return
				}
			}
		}
	}()

	<-done

	if len(lP) != len(lQ) {
		return false
	}
	fmt.Printf("LP: %v\n", lP)
	fmt.Printf("LQ: %v\n", lP)

	for i, v := range lP {
		if v != lQ[i] {
			return false
		}
	}
	return true
}

func isSameTreeX(node *TreeNode, ch chan int) {
	if node == nil {
		ch <- math.MinInt
		return
	}
	ch <- node.Val
	isSameTreeX(node.Left, ch)
	isSameTreeX(node.Right, ch)
}
