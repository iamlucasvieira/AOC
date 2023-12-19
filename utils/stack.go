package utils

// Stack is a stack of T
type Stack[T any] []T

// Push pushes a value onto the stack
func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// Pop pops a value off the stack
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(*s) == 0 {
		return zero, false
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, true
}

// Peek returns the top value of the stack without removing it
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(*s) == 0 {
		return zero, false
	}

	v := (*s)[len(*s)-1]
	return v, true
}
