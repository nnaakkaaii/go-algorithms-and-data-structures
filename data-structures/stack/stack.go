package stack

import "errors"

type Stack struct {
	top int
	max int
	arr []interface{}
}

// NewStack initialize Stack
func NewStack(max int) *Stack {
	stack := new(Stack)
	stack.top = 0
	stack.max = max
	stack.arr = make([]interface{}, max)
	return stack
}

// IsEmpty if Stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

// IsFull if Stack is full
func (s *Stack) IsFull() bool {
	return s.top >= s.max - 1
}

// Push adds an element
func (s *Stack) Push(v interface{}) error {
	if s.IsFull() {
		return errors.New("stack overflow")
	}
	s.top++
	s.arr[s.top] = v
	return nil
}

// Pop delete an element
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack underflow")
	}
	s.top--
	return s.arr[s.top + 1], nil
}
