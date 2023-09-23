package main

import "fmt"

type IQueue[T any] interface {
	Enqueu(value T)
	Dequeue() T
	Front() T
	Size() int
	IsEmpty() bool
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedListQueue[T any] struct {
	head *Node[T]
}

func newLinkedListQueue[T any]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{}
}

func (queue *LinkedListQueue[T]) IsEmpty() bool {
	return queue.head == nil
}

func (queue *LinkedListQueue[T]) Size() int {
	current := queue.head
	count := 0
	for current != nil {
		current = current.next
		count++
	}
	return count
}

func (queue *LinkedListQueue[T]) Enqueue(value T) {
	newNode := Node[T]{value: value, next: nil}
	if queue.Size() == 0 {
		queue.head = &newNode
		return
	}
	current := queue.head

	for current.next != nil {
		current = current.next
	}

	current.next = &newNode
}

func (queue *LinkedListQueue[T]) Dequeue() T {
	if queue.Size() == 0 {
		fmt.Println("A fila já está vazia")
		var None T
		return None
	}
	item := queue.head.value
	queue.head = queue.head.next
	return item
}

func (queue *LinkedListQueue[T]) Front() T {
	if queue.Size() == 0 {
		fmt.Println("A fila está vazia")
		var None T
		return None
	}
	item := queue.head.value
	return item
}

func main() {
	queue := newLinkedListQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	fmt.Println("Tamanho da fila: ", queue.Size())
	fmt.Println("Frente da fila: ", queue.Front())
	fmt.Println("--Esvaziando a fila--")

	for !queue.IsEmpty() {
		item := queue.Dequeue()
		fmt.Println("Saiu da fila: ", item)
	}
}
