package problem4_1

import (
	"testing"
)

func TestFindRouteBFS(t *testing.T) {
	// Test 1: Simple path
	a, b, c := &NodeB{Value: 1}, &NodeB{Value: 2}, &NodeB{Value: 3}
	a.Neighbours = []*NodeB{b}
	b.Neighbours = []*NodeB{c}
	if !FindRouteBFS(a, c) {
		t.Errorf("Test 1 failed: expected true, got false")
	}

	// Test 2: Disconnected NodeBs
	d := &NodeB{Value: 4}
	if FindRouteBFS(a, d) {
		t.Errorf("Test 2 failed: expected false, got true")
	}

	// Test 3: Self-loop
	f := &NodeB{Value: 6}
	f.Neighbours = []*NodeB{f}
	if !FindRouteBFS(f, f) {
		t.Errorf("Test 3 failed: expected true, got false")
	}

	// Test 4: Cyclic graph with a path
	g, h, i, j := &NodeB{Value: 7}, &NodeB{Value: 8}, &NodeB{Value: 9}, &NodeB{Value: 10}
	g.Neighbours = []*NodeB{h}
	h.Neighbours = []*NodeB{i}
	i.Neighbours = []*NodeB{j, h} // Cycle
	if !FindRouteBFS(g, j) {
		t.Errorf("Test 4 failed: expected true, got false")
	}

	// Test 5: Cyclic graph without a path
	k, l := &NodeB{Value: 11}, &NodeB{Value: 12}
	l.Neighbours = []*NodeB{k}
	if FindRouteBFS(l, g) {
		t.Errorf("Test 5 failed: expected false, got true")
	}

	// Test 6: Large graph
	m, n, o, p, q, r := &NodeB{Value: 13}, &NodeB{Value: 14}, &NodeB{Value: 15}, &NodeB{Value: 16}, &NodeB{Value: 17}, &NodeB{Value: 18}
	m.Neighbours = []*NodeB{n}
	n.Neighbours = []*NodeB{o}
	o.Neighbours = []*NodeB{p}
	p.Neighbours = []*NodeB{q}
	q.Neighbours = []*NodeB{r}
	if !FindRouteBFS(m, r) {
		t.Errorf("Test 6 failed: expected true, got false")
	}
}
