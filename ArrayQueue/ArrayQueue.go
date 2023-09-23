package main

import "fmt"

type IQueue[T any] interface {
	Enqueue(value T)
	Dequeue() T
	Front() T
	IsEmpty() bool
	Size() int
}

type Queue[T any] struct {
	values []T
}

func newQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (queue *Queue[T]) IsEmpty() bool {
	return len(queue.values) == 0
}

func (queue *Queue[T]) Size() int {
	return len(queue.values)
}

func (queue *Queue[T]) Enqueue(value T) {
	queue.values = append(queue.values, value)
}

func (queue *Queue[T]) Dequeue() T {
	if queue.IsEmpty() {
		fmt.Println("Não há itens na fila")
		var None T
		return None
	}
	item := queue.values[0]
	queue.values = queue.values[1:]
	return item
}

func (queue *Queue[T]) Front() T {
	if queue.IsEmpty() {
		fmt.Println("Não há itens na fila")
		var None T
		return None
	}
	return queue.values[0]
}

func main() {
	queue := newQueue[int]()
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
