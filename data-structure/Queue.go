package data_structure

type Queue struct {
	items []int
}

// Q
// front pop, access
func (q *Queue) Front() int {
	if len(q.items) == 0 {
		panic("Q underflow")
	}
	return q.items[0]
}

// back push
func (q *Queue) Push(value int) {
	q.items = append(q.items, value)
}

// size()
func (q *Queue) Size() int {
	return len(q.items)
}

// isEmpty()
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// pop()
func (q *Queue) Pop() int {

	if len(q.items) == 0 {
		panic("Q Underflow")
	}
	first := q.items[0]
	q.items = append([]int{}, q.items[1:]...)
	return first
}
