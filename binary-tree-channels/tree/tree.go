package tree

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

type Sequence []int

func New(value int) *Tree {
	seq := NewSequence(value)
	root := &Tree{nil, 0, nil}

	for len(*seq) > 0 {
		allocate(root, seq)
	}
	return root
}

func NewSequence(v int) *Sequence {
	count := 10
	values := make(Sequence, count)
	for i := 0; len(values) > i; i++ {
		values[i] = (i + 1) * v
	}

	return &values
}

func allocate(tree *Tree, seq *Sequence) {
	if tree == nil || len(*seq) == 0 {
		return
	}

	if tree.Value == 0 {
		tree.Value = ExtractRandomValue(seq)
	}

	if hasChild := tree.Left != nil || tree.Right != nil; !hasChild {
		var r *rand.Rand
		r = GetRandom()

		numberOfChildren := Min(r.Intn(3), len(*seq))
		switch numberOfChildren {
		case 1:
			CreateOneRandomChild(tree, r)
		case 2:
			tree.Left = &Tree{nil, 0, nil}
			tree.Right = &Tree{nil, 0, nil}
		}
	}

	allocate(tree.Left, seq)
	allocate(tree.Right, seq)
}

func GetRandom() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s)
}

func CreateOneRandomChild(tree *Tree, r *rand.Rand) {
	var child **Tree
	children := []**Tree{&tree.Left, &tree.Right}
	child = children[r.Intn(2)]
	*child = &Tree{nil, 0, nil}
}

func ExtractRandomValue(seq *Sequence) int {
	r := GetRandom()
	slice := *seq
	i := r.Intn(len(slice))
	var selected int
	selected, slice[i] = slice[i], slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	*seq = slice
	//	*seq = Sequence(slice[:slice.Count()-1])
	return selected
}

func (t *Tree) Print() {
	PrintInOrder(t)
}

func PrintInOrder(t *Tree) {
	if t == nil {
		return
	}
	PrintInOrder(t.Left)
	fmt.Printf("%d ", t.Value)
	PrintInOrder(t.Right)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
