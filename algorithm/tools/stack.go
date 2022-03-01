package tools

type Stack struct {
	len int
	value []interface{}
}

func NewStack() *Stack {
	return &Stack{
		len:   0,
		value: make([]interface{}, 0),
	}
}
func (s * Stack) Pop() interface{}{
	if s.len == 0 {
		return nil
	}
	lastIndex := s.len-1
	ret := s.value[lastIndex]
	s.value = s.value[0:lastIndex]
	s.len--
	return ret
}
func (s *Stack) Push(v interface{})  {
	s.value = append(s.value, v)
	s.len++
}
