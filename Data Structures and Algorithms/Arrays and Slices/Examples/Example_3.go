// Implementing a Stack with a Slice

package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false // Stack is empty
	}
	index := len(s.data) - 1
	value := s.data[index]
	s.data = s.data[:index]
	return value, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func main() {
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Is stack empty?", stack.IsEmpty()) // Output: Is stack empty? false

	value, ok := stack.Pop()
	if ok {
		fmt.Println("Popped:", value) // Output: Popped: 30
	}

	value, ok = stack.Pop()
	if ok {
		fmt.Println("Popped:", value) // Output: Popped: 20
	}

	value, ok = stack.Pop()
	if ok {
		fmt.Println("Popped:", value) // Output: Popped: 10
	}

	fmt.Println("Is stack empty?", stack.IsEmpty()) // Output: Is stack empty? true

	_, ok = stack.Pop()
	if !ok {
		fmt.Println("Stack is empty") // Output: Stack is empty
	}
}
