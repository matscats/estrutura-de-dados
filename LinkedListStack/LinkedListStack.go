package main

import "fmt"

type IStack[T any] interface {
	Push(value T)
	Pop() T
	Peek() T
	IsEmpty() bool
	Size() int
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedListStack[T any] struct {
	head *Node[T]
}

func NewLinkedListStack[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{head: nil}
}

func (stack *LinkedListStack[T]) IsEmpty() bool {
	return stack.head == nil
}

func (stack *LinkedListStack[T]) Size() int {
	count := 0
	current := stack.head

	for current != nil {
		current = current.next
		count++
	}

	return count
}

func (stack *LinkedListStack[T]) Push(value T) {
	newNode := Node[T]{value: value, next: nil}
	if stack.IsEmpty() {
		stack.head = &newNode
		return
	}
	current := stack.head

	for current.next != nil {
		current = current.next
	}
	current.next = &newNode
}

func (stack *LinkedListStack[T]) Pop() T {
	if stack.IsEmpty() {
		fmt.Println("A Stack já está vazia")
		var None T
		return None
	}

	if stack.Size() == 1 {
		item := stack.head.value
		stack.head = nil
		return item
	}

	current := stack.head

	for current.next.next != nil {
		current = current.next
	}
	item := current.next.value
	current.next = nil
	return item
}

func (stack *LinkedListStack[T]) Peek() T {
	if stack.IsEmpty() {
		fmt.Println("A Stack está vazia")
		var None T
		return None
	}

	current := stack.head

	for current.next != nil {
		current = current.next
	}

	item := current.value
	return item
}

func main() {
	stack := NewLinkedListStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println("Tamanho da pilha: ", stack.Size())
	fmt.Println("Topo da pilha: ", stack.Peek())

	for !stack.IsEmpty() {
		item := stack.Pop()
		fmt.Println("Item desempilhado:", item)
	}
}
