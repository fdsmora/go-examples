package queue

type Queue[T any] []T

func (q *Queue[T]) Dequeue() T {
	if len(*q) == 0 {
		var zero T
		return zero
	}
	c := (*q)[0]
	*q = (*q)[1:]
	return c
}

func (q *Queue[T]) Enqueue(c T) {
	*q = append((*q), c)
}
