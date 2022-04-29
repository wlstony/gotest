package main

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	container := make([]interface{}, 0)
	return &Queue{elements: container}
}

func (s *Queue) Push(node interface{}) {
	s.elements = append(s.elements, node)
}

func (s *Queue) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Queue) Pop() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	node := s.elements[0]
	s.elements = s.elements[1:]
	return node
}
