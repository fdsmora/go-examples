package stack

import "testing"

func TestStackInts(t *testing.T) {
	var s Stack[int]
	if len(s) != 0 {
		t.Fatalf("expected empty stack, got len %d", s.Len())
	}

	// Peek on empty returns zero value
	if got := s.Peek(); got != 0 {
		t.Fatalf("Peek on empty: got %v want 0", got)
	}

	// Pop on empty
	if _, ok := s.Pop(); ok {
		t.Fatalf("Pop on empty should return ok=false")
	}

	// Push sequence
	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Len() != 3 {
		t.Fatalf("after pushes len = %d; want 3", s.Len())
	}

	if got := s.Peek(); got != 30 {
		t.Fatalf("Peek = %v; want 30", got)
	}

	// Pop order
	if v, ok := s.Pop(); !ok || v != 30 {
		t.Fatalf("Pop 1 = (%v,%v); want (30,true)", v, ok)
	}
	if v, ok := s.Pop(); !ok || v != 20 {
		t.Fatalf("Pop 2 = (%v,%v); want (20,true)", v, ok)
	}
	if v, ok := s.Pop(); !ok || v != 10 {
		t.Fatalf("Pop 3 = (%v,%v); want (10,true)", v, ok)
	}
	if _, ok := s.Pop(); ok {
		t.Fatalf("Pop on empty after pops should return ok=false")
	}
}

func TestStackStrings(t *testing.T) {
	var s Stack[string]
	vals := []string{"a", "bc", "def"}
	for _, v := range vals {
		s.Push(v)
	}

	if s.Len() != len(vals) {
		t.Fatalf("len = %d; want %d", s.Len(), len(vals))
	}

	// check LIFO
	for i := len(vals) - 1; i >= 0; i-- {
		want := vals[i]
		if got := s.Peek(); got != want {
			t.Fatalf("peek = %q; want %q", got, want)
		}
		if v, ok := s.Pop(); !ok || v != want {
			t.Fatalf("pop = (%q,%v); want (%q,true)", v, ok, want)
		}
	}
}

func TestStackCustomType(t *testing.T) {
	type pair struct{ A, B int }
	var s Stack[pair]
	p1 := pair{1, 2}
	p2 := pair{3, 4}
	s.Push(p1)
	s.Push(p2)

	if got := s.Peek(); got != p2 {
		t.Fatalf("peek = %v; want %v", got, p2)
	}
	if v, ok := s.Pop(); !ok || v != p2 {
		t.Fatalf("pop = (%v,%v); want (%v,true)", v, ok, p2)
	}
	if v, ok := s.Pop(); !ok || v != p1 {
		t.Fatalf("pop = (%v,%v); want (%v,true)", v, ok, p1)
	}
}
