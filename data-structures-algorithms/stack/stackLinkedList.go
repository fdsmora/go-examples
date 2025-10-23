package stack

import "errors"

type StackLinkedList[T any] struct {
	top *stackNode[T]
}

type stackNode[T any] struct {
	value T
	next  *stackNode[T]
}

func NewStackLinkedList[T any]() *StackLinkedList[T] {
	return &StackLinkedList[T]{}
}

func (s *StackLinkedList[T]) Pop() (T, error) {
	if s.top == nil {
		var zero T
		return zero, errors.New("stack is empty")
	}
	value := s.top.value
	s.top = s.top.next
	return value, nil
}

func (s *StackLinkedList[T]) Put(value T) {
	newNode := &stackNode[T]{value, nil}
	newNode.next = s.top
	s.top = newNode
}

func (s *StackLinkedList[T]) Peek() (T, error) {
	if s.top == nil {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.top.value, nil
}

func (s *StackLinkedList[T]) IsEmpty() bool {
	return s.top == nil
}
