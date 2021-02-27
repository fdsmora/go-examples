package tree

import (
	"math/rand"
	"testing"
	"time"
)

func TestCreateOneRandomChild(t *testing.T) {
	node := &Tree{nil, 0, nil}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	CreateOneRandomChild(node, r)

	if hasChild := node.Left != nil || node.Right != nil; !hasChild {
		t.Fail()
	}

	if node.Left != nil {
		t.Log("Node left is ", node.Left)
	} else {
		t.Log("Node right is ", node.Right)
	}
}
