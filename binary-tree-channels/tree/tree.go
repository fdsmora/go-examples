package tree

import (
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

type Sequence struct {
	values []int
	count  int
}

func New(value int) *Tree {
	seq := makeSequence(value)
	root := &Tree{nil, 0, nil}

	for seq.count > 0 {
		allocate(root, seq)
	}
	return root
}

func makeSequence(v int) *Sequence {
	count := 10
	values := make([]int, 10)
	for i := 0; len(values) > 0; i++ {
		values[i] = (i + 1) * v
	}

	return &Sequence{values, count}
}

func allocate(tree *Tree, seq *Sequence) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	tree.Value = seq.GetValue()

	if hasChild := tree.Left != nil || tree.Right != nil; !hasChild {
		numberOfChildren := r.Intn(3)
		switch numberOfChildren {
		case 1:
			CreateOneRandomChild(tree, r)
		case 2:
			tree.Left = &Tree{nil, 0, nil}
			tree.Right = &Tree{nil, 0, nil}
		}
	}
}

func CreateOneRandomChild(tree *Tree, r *rand.Rand) {
	var child **Tree
	children := []**Tree{&tree.Left, &tree.Right}
	child = children[r.Intn(2)]
	*child = &Tree{nil, 1, nil}
}

func (seq *Sequence) GetValue() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// TODO : Include decrease count
	for {
		i := r.Intn(seq.count)
		if seq.values[i] > -1 {
			selected := seq.values[i]
			seq.values[i] = -1
			return selected
		}
	}
}
