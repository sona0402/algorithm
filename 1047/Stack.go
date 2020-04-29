package main

// 创建Stack
type Stack struct {
	index int
	data  [16]int
}

func (s *Stack) Push(k int) {
	s.data[s.index] = k
	s.index++
}

func (s *Stack) Pop() (ret int) {
	s.index--
	return s.data[s.index]
}

func (s *Stack) Peek() (ret int) {
	tempIndex := s.index
	tempIndex--
	return s.data[tempIndex]
}

func (s *Stack) IsEmpty() bool {
	return s.index <= 0
}
