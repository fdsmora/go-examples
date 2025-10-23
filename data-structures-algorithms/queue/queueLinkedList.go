package queue

import "errors"

type QueueLinkedList[T any] struct {
	head, tail *QueueNode[T]
}

type QueueNode[T any] struct {
	value T
	next  *QueueNode[T]
}

func (q *QueueLinkedList[T]) Pop() (T, error) {
	if q.head == nil {
		var zero T
		return zero, errors.New("empty list")
	}
	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return value, nil
}

func (q *QueueLinkedList[T]) Put(value T) {
	newNode := &QueueNode[T]{value, nil}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return
	}
	q.tail.next = newNode
	q.tail = newNode
}

func (q *QueueLinkedList[T]) Peek() (T, error) {
	if q.head == nil {
		var zero T
		return zero, errors.New("empty queue")
	}
	return q.head.value, nil
}

func (q *QueueLinkedList[T]) IsEmpty() bool {
	return q.head == nil
}
