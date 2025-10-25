package chapter3

type StackNode struct {
	value int
	prev,
	prevMin *StackNode
}

type StackLinkedList struct {
	top,
	min *StackNode
}

func (s *StackLinkedList) Push(value int) {
	if s.top == nil {
		s.top = &StackNode{value, nil, nil}
		s.min = s.top
		return
	}
	newNode := &StackNode{value, nil, nil}
	newNode.prev = s.top
	s.top = newNode
	if value <= s.min.value {
		newNode.prevMin = s.min
		s.min = newNode
	}
}

func (s *StackLinkedList) Pop() (int, bool) {
	if s.top != nil {
		node := s.top
		if s.top == s.min {
			s.min = s.min.prevMin
		}
		s.top = s.top.prev
		return node.value, true
	}
	return 0, false
}

func (s *StackLinkedList) Min() (int, bool) {
	if s.min != nil {
		return s.min.value, true
	}
	return 0, false
}
