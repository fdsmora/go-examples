package tree

import (
	"math/rand"
	"reflect"
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

func TestNewSequence(t *testing.T) {
	expected := Sequence{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	if !reflect.DeepEqual(expected, NewSequence(2)) {
		t.Fail()
	}
	//	fmt.Printf("%v\n", reflect.DeepEqual(expected, NewSequence(2)))
}

func TestExtractRandomValue(t *testing.T) {
	seq := NewSequence(1)
	for seq.Count() > 0 {
		extracted := seq.ExtractRandomValue()
		t.Log(extracted)
	}
}
