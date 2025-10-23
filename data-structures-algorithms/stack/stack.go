package stack

type Stack[T interface{}] []T // type Stack[T any] []T

// O(1)
func (s *Stack[T]) Peek() T {
	if len(*s) == 0 {
		var zero T
		return zero
	}
	return (*s)[len(*s)-1]
}

// O(1) amortized
func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	idx := len(*s) - 1
	val := (*s)[idx]
	*s = (*s)[:idx] // O(1) amortized
	return val, true
}

// O(1) amortized
func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Len() int {
	return len(*s)
}
