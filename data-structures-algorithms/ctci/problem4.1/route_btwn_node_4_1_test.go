package problem4_1

import (
	"testing"
)

func TestFindRoute(t *testing.T) {
	// Test 1: Simple path
	a, b, c := &Node{Value: 1}, &Node{Value: 2}, &Node{Value: 3}
	a.Neighbours = []*Node{b}
	b.Neighbours = []*Node{c}
	if !FindRoute(a, c) {
		t.Errorf("Test 1 failed: expected true, got false")
	}

	// Test 2: Disconnected nodes
	d := &Node{Value: 4}
	if FindRoute(a, d) {
		t.Errorf("Test 2 failed: expected false, got true")
	}

	// Test 3: Self-loop
	f := &Node{Value: 6}
	f.Neighbours = []*Node{f}
	if !FindRoute(f, f) {
		t.Errorf("Test 3 failed: expected true, got false")
	}

	// Test 4: Cyclic graph with a path
	g, h, i, j := &Node{Value: 7}, &Node{Value: 8}, &Node{Value: 9}, &Node{Value: 10}
	g.Neighbours = []*Node{h}
	h.Neighbours = []*Node{i}
	i.Neighbours = []*Node{j, h} // Cycle
	if !FindRoute(g, j) {
		t.Errorf("Test 4 failed: expected true, got false")
	}

	// Test 5: Cyclic graph without a path
	k, l := &Node{Value: 11}, &Node{Value: 12}
	l.Neighbours = []*Node{k}
	if FindRoute(l, g) {
		t.Errorf("Test 5 failed: expected false, got true")
	}

	// Test 6: Large graph
	m, n, o, p, q, r := &Node{Value: 13}, &Node{Value: 14}, &Node{Value: 15}, &Node{Value: 16}, &Node{Value: 17}, &Node{Value: 18}
	m.Neighbours = []*Node{n}
	n.Neighbours = []*Node{o}
	o.Neighbours = []*Node{p}
	p.Neighbours = []*Node{q}
	q.Neighbours = []*Node{r}
	if !FindRoute(m, r) {
		t.Errorf("Test 6 failed: expected true, got false")
	}
}
