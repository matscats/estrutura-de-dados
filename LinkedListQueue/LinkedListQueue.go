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
	rear  *Node[T]
	front *Node[T]
}

func newLinkedListQueue[T any]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{rear: nil, front: nil}
}

func (queue *LinkedListQueue[T]) IsEmpty() bool {
	return queue.front == nil
}

func (queue *LinkedListQueue[T]) Size() int {
	current := queue.front
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
		queue.rear = &newNode
		queue.front = &newNode
		return
	}
	queue.rear.next = &newNode
	queue.rear = &newNode
}

func (queue *LinkedListQueue[T]) Dequeue() T {
	if queue.Size() == 0 {
		fmt.Println("A fila já está vazia")
		var None T
		return None
	}
	item := queue.front.value
	queue.front = queue.front.next
	return item
}

func (queue *LinkedListQueue[T]) Front() T {
	if queue.Size() == 0 {
		fmt.Println("A fila está vazia")
		var None T
		return None
	}
	item := queue.front.value
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
