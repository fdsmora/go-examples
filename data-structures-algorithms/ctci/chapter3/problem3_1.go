package chapter3

// CTCI problem 3.1
type ThreeStacks struct {
	totalStacks               int
	tops, limits, lowerLimits [3]int
	list                      []int
}

func New(n int) (*ThreeStacks, bool) {
	if n < 3 {
		return nil, false
	}
	stacks := &ThreeStacks{}
	stacks.totalStacks = 3
	size := n / stacks.totalStacks
	stacks.tops = [3]int{-1, size - 1, 2*size - 1}
	stacks.limits = [3]int{size, 2 * size, n}
	stacks.lowerLimits = [3]int{0, size, 2 * size}
	stacks.list = make([]int, n)
	return stacks, true
}

func (s *ThreeStacks) Push(stackIndex, value int) bool {
	if stackIndex < 0 || stackIndex >= s.totalStacks {
		return false
	}
	top := s.tops[stackIndex]
	limit := s.limits[stackIndex]
	if top+1 >= limit {
		return false
	}
	top++
	s.list[top] = value
	s.tops[stackIndex] = top
	return true
}

func (s *ThreeStacks) Pop(stackIndex int) (int, bool) {
	if stackIndex < 0 || stackIndex >= s.totalStacks {
		return 0, false
	}
	top := s.tops[stackIndex]
	if top < s.lowerLimits[stackIndex] {
		return 0, false
	}
	value := s.list[top]
	top--
	s.tops[stackIndex] = top
	return value, true
}

func (s *ThreeStacks) Peek(stackIndex int) (int, bool) {
	if stackIndex < 0 || stackIndex >= s.totalStacks {
		return 0, false
	}
	top := s.tops[stackIndex]
	if top < s.lowerLimits[stackIndex] {
		return 0, false
	}
	return s.list[top], true
}
