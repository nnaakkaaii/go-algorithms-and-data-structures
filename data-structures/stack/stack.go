package stack

import "errors"

type Stack struct {
	top int
	max int
	arr []int
}

// NewStack initialize Stack
func NewStack(max int) *Stack {
	stack := new(Stack)
	stack.top = 0
	stack.max = max
	stack.arr = make([]int, max)
	return stack
}

// isEmpty if Stack is empty
func (s *Stack) isEmpty() bool {
	return s.top == 0
}

// isFull if Stack is full
func (s *Stack) isFull() bool {
	return s.top >= s.max - 1
}

// Push adds an element
func (s *Stack) Push(v int) error {
	if s.isFull() {
		return errors.New("stack overflow")
	}
	s.top++
	s.arr[s.top] = v
	return nil
}

// Pop delete an element
func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack underflow")
	}
	s.top--
	return s.arr[s.top + 1], nil
}
