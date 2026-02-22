package data_structure

type Queue[T any] struct {
	items []T
}

// Front Q
// front pop, access
func (q *Queue[T]) Front() T {
	if len(q.items) == 0 {
		panic("Q underflow")
	}
	return q.items[0]
}

// Push back push
func (q *Queue[T]) Push(value T) {
	q.items = append(q.items, value)
}

// Size size()
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// IsEmpty isEmpty()
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Pop pop()
func (q *Queue[T]) Pop() T {

	if len(q.items) == 0 {
		panic("Q Underflow")
	}
	first := q.items[0]
	q.items = append([]T{}, q.items[1:]...)
	return first
}
