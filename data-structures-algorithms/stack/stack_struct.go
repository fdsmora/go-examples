package stack

import (
	"errors"
)

var (
	errStackIsEmpty = errors.New("stack is empty")
)

type StackStruct struct {
	list []rune
	top  int
}

func NewStack() *StackStruct {
	return &StackStruct{list: make([]rune, 0), top: -1}
}

func (s *StackStruct) Push(value rune) {
	s.top++
	if s.top < len(s.list) {
		// There is capacity, overwrite
		s.list[s.top] = value
	} else {
		// Extend backing array
		s.list = append(s.list, value)
	}
}

func (s *StackStruct) Pop() (rune, error) {
	if s.top < 0 {
		return 0, errStackIsEmpty
	}
	value := s.list[s.top]
	s.top--
	return value, nil
}

func (s *StackStruct) Peek() (rune, error) {
	if s.top < 0 {
		return 0, errStackIsEmpty
	}
	return s.list[s.top], nil
}

func (s *StackStruct) Size() int {
	return s.top + 1
}

func (s *StackStruct) IsEmpty() bool {
	return s.top < 0
}
