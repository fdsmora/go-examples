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

func (seq Sequence) Count() int {
	return len(seq)
}

func New(value int) *Tree {
	seq := NewSequence(value)
	root := &Tree{nil, 0, nil}

	for seq.Count() > 0 {
		allocate(root, seq)
	}
	return root
}

func NewSequence(v int) Sequence {
	count := 10
	values := make(Sequence, count)
	for i := 0; len(values) > i; i++ {
		values[i] = (i + 1) * v
	}

	return values
}

func allocate(tree *Tree, seq Sequence) {
	if tree == nil || seq.Count() == 0 {
		return
	}

	if tree.Value == 0 {
		tree.Value = ExtractRandomValue(&seq)
	}

	if hasChild := tree.Left != nil || tree.Right != nil; !hasChild {
		var r *rand.Rand
		r = GetRandom()

		numberOfChildren := r.Intn(3)
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
	i := r.Intn(slice.Count())
	var selected int
	selected, slice[i] = slice[i], slice[slice.Count()-1]
	slice = slice[:slice.Count()-1]
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
