package main

type Stack struct {
	elements []interface{}
}

func NewStack() *Stack {
	container := make([]interface{}, 0)
	return &Stack{elements: container}
}

func (s *Stack) Push(node interface{}) {
	s.elements = append(s.elements, node)
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Pop() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	last := len(s.elements) - 1
	node := s.elements[last]
	s.elements = s.elements[:last]
	return node
}
