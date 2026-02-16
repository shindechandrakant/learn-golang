package data_structure

type Stack struct {
	size   int
	values []int
}

// push
// pop
// front
// isEmpty
// size
// capacity

func (s *Stack) Push(element int) {
	s.size++
	s.values = append(s.values, element)
}

func (s *Stack) Pop() {

	if s.size == 0 {
		panic("Stack Underflow")
	}
	s.values = append([]int{}, s.values[1:]...)
	s.size--
}

func (s *Stack) Front() int {
	if s.size == 0 {
		panic("Stack Underflow")
	}
	return s.values[0]
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Length() int {
	return s.size
}
