package main

import "fmt"

// Stack 泛型栈实现
type Stack[T any] struct {
    elements []T
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T]{}
}

func (s *Stack[T]) Push(element T) {
    s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }
    
    index := len(s.elements) - 1
    element := s.elements[index]
    s.elements = s.elements[:index]
    return element, true
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.elements) == 0
}

func main() {
    // 整型栈
    intStack := NewStack[int]()
    intStack.Push(1)
    intStack.Push(2)
    if val, ok := intStack.Pop(); ok {
        fmt.Printf("Popped integer: %d\n", val)
    }

    // 字符串栈
    stringStack := NewStack[string]()
    stringStack.Push("hello")
    stringStack.Push("world")
    if val, ok := stringStack.Pop(); ok {
        fmt.Printf("Popped string: %s\n", val)
    }
}
