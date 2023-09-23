package main

import "fmt"

type IStack[T any] interface {
	Push(value T)
	Pop() T
	Peek() T
	IsEmpty() bool
	Size() int
}

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.data) == 0
}

func (stack *Stack[T]) Push(value T) {
	stack.data = append(stack.data, value)
}

func (stack *Stack[T]) Pop() T {
	if stack.IsEmpty() {
		fmt.Println("A Stack já está vazia")
		var None T
		return None
	}
	index := len(stack.data) - 1
	item := stack.data[index]
	stack.data = stack.data[:index]
	return item
}

func (stack *Stack[T]) Peek() T {
	if stack.IsEmpty() {
		fmt.Println("A stack está vazia")
		var None T
		return None
	}
	return stack.data[len(stack.data)-1]
}

func (stack *Stack[T]) Size() int {
	return len(stack.data)
}

func main() {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println("Tamanho da pilha: ", stack.Size())
	fmt.Println("Topo da pilha: ", stack.Peek())

	for !stack.IsEmpty() {
		item := stack.Pop()
		fmt.Println("Item desempilhado: ", item)
	}
}
