package queue

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Dequeue() T {
	if len(q.items) == 0 {
		var zero T
		return zero
	}
	c := q.items[0]
	q.items = q.items[1:]
	return c
}

func (q *Queue[T]) Enqueue(c T) {
	q.items = append(q.items, c)
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}
