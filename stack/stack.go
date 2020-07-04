package stack

// Stack data-structure
type Stack []interface{}

// Push a value onto the top of the stack
func (s *Stack) Push(data interface{}) {
	*s = append(*s, data)
}

// Top return a top value of stack
func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return (*s)[len(*s)-1]
}

// Pop the top of item on the stack
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	val := s.Top()

	(*s) = (*s)[:(len(*s) - 1)]

	return val
}

// IsEmpty checks the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
