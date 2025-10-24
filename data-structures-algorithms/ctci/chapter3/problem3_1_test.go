package chapter3

import "testing"

func TestThreeStacksBasic(t *testing.T) {
	// create array of size 9 -> each stack size 3
	s, ok := New(9)
	if !ok || s == nil {
		t.Fatalf("New(9) failed")
	}

	// initially peeks/pops should fail (empty)
	if _, ok := s.Peek(0); ok {
		t.Fatalf("Peek should fail on empty stack 0")
	}
	if _, ok := s.Pop(1); ok {
		t.Fatalf("Pop should fail on empty stack 1")
	}

	// push up to capacity for each stack
	for si := 0; si < 3; si++ {
		for i := 0; i < 3; i++ {
			if !s.Push(si, si*10+i) {
				t.Fatalf("Push failed for stack %d at item %d", si, i)
			}
		}
		// next push should fail (full)
		if s.Push(si, 999) {
			t.Fatalf("Push succeeded when stack %d should be full", si)
		}
	}

	// verify Peek returns last pushed values
	for si := 0; si < 3; si++ {
		want := si*10 + 2
		if v, ok := s.Peek(si); !ok || v != want {
			t.Fatalf("Peek stack %d = (%d,%v); want (%d,true)", si, v, ok, want)
		}
	}

	// pop all and verify order (LIFO)
	for si := 0; si < 3; si++ {
		for i := 2; i >= 0; i-- {
			want := si*10 + i
			if v, ok := s.Pop(si); !ok || v != want {
				t.Fatalf("Pop stack %d = (%d,%v); want (%d,true)", si, v, ok, want)
			}
		}
		// now empty
		if _, ok := s.Pop(si); ok {
			t.Fatalf("Pop should fail on empty stack %d after popping all", si)
		}
	}
}

func TestThreeStacksInvalidIndices(t *testing.T) {
	s, ok := New(6) // each stack size 2
	if !ok {
		t.Fatalf("New(6) failed")
	}

	// negative index
	if s.Push(-1, 1) {
		t.Fatalf("Push with negative index should fail")
	}
	if _, ok := s.Pop(3); ok {
		t.Fatalf("Pop with out-of-range index should fail")
	}
	if _, ok := s.Peek(5); ok {
		t.Fatalf("Peek with out-of-range index should fail")
	}
}

func TestThreeStacksEdgeSizes(t *testing.T) {
	// n not divisible by 3 -> integer division used
	s, ok := New(10)
	if !ok {
		t.Fatalf("New(10) failed")
	}
	// with n=10 size = 3 each, limits are at 3,6,10 and lowerLimits 0,3,6
	// push to second stack's capacity (3 elements)
	if !s.Push(1, 11) || !s.Push(1, 12) || !s.Push(1, 13) {
		t.Fatalf("Push into stack 1 failed")
	}
	// third stack has capacity 4 (limit up to n)
	for i := 0; i < 4; i++ {
		if !s.Push(2, 20+i) {
			t.Fatalf("Push into stack 2 failed at %d", i)
		}
	}
	// further push into stack 2 should fail
	if s.Push(2, 99) {
		t.Fatalf("Push succeeded into full stack 2")
	}
}
