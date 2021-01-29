package stack

type ArrayStack struct {
	values []interface{}
	top    int
}

func NewArrayStack() Stack {
	return &ArrayStack{
		top: -1,
	}
}

func (s *ArrayStack) Push(value interface{}) int {
	s.top++
	if s.top >= len(s.values) {
		s.values = append(s.values, value)
	} else {
		s.values[s.top] = value
	}
	return 0
}

func (s *ArrayStack) Pop() interface{} {
	if s.top < 0 {
		return nil
	}

	value := s.values[s.top]
	s.top--
	return value
}

func (s *ArrayStack) IsEmpty() bool {
	return s.top < 0
}

func (s *ArrayStack) Top() interface{} {
	if s.top < 0 {
		return nil
	}
	return s.values[s.top]
}

func (s *ArrayStack) Flush() {
	s.top = -1
}
