package main

import (
	"fmt"
)

type List[T any] interface {
	AddLast(value T)
	AddPos(value T, pos int)
	Update(value T, pos int)
	RemoveLast()
	Remove(pos int)
	Get(pos int) T
	Size() int
}

type ArrayList[T any] struct {
	values   []T
	inserted int
}

func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{values: make([]T, 10), inserted: 0}
}

func (list *ArrayList[T]) doubleArray() {
	curSize := len(list.values)
	doubledValues := make([]T, 2*curSize)

	for i := 0; i < curSize; i++ {
		doubledValues[i] = list.values[i]
	}
	list.values = doubledValues
}

func (list *ArrayList[T]) AddLast(value T) {

	if list.inserted >= len(list.values) {
		list.doubleArray()
	}
	list.values[list.inserted] = value
	list.inserted++
}

func (list *ArrayList[T]) AddPos(value T, pos int) {
	if pos < 0 || pos > list.inserted {
		fmt.Println("Posição fora dos limites do ArrayList")
	}
	if list.inserted >= len(list.values) {
		list.doubleArray()
	}
	for i := list.inserted; i > pos; i-- {
		list.values[i] = list.values[i-1]
	}
	list.values[pos] = value
	list.inserted++
}

func (list *ArrayList[T]) Size() int {
	return list.inserted
}

func (list *ArrayList[T]) Update(value T, pos int) {
	if pos < 0 || pos > list.inserted {
		fmt.Println("Posição fora dos limites do ArrayList")
		return
	}
	list.values[pos] = value
}

func (list *ArrayList[T]) RemoveLast() {
	if list.inserted == 0 {
		fmt.Print("A lista já está vazia")
		return
	}
	list.inserted--
}

func (list *ArrayList[T]) Remove(pos int) {
	if list.inserted == 0 {
		fmt.Print("A lista já está vazia")
		return
	}
	if pos < 0 || pos > list.inserted {
		fmt.Println("Posição fora dos limites do ArrayList")
		return
	}
	for i := pos; i < list.inserted; i++ {
		list.values[i] = list.values[i+1]
	}
	list.inserted--
}

func (list *ArrayList[T]) Get(pos int) T {
	if pos < 0 || pos > list.inserted {
		fmt.Println("Posição fora dos limites do ArrayList")
	}
	return list.values[pos]
}

func main() {
	lista := NewArrayList[int]()
	lista.AddLast(1)
	lista.AddLast(2)
	lista.AddLast(3)
	lista.AddLast(4)
	lista.AddLast(5)
	lista.AddLast(6)
	lista.AddLast(7)
	lista.AddLast(8)
	lista.AddLast(9)
	lista.AddLast(10)
	lista.AddLast(11)
	lista.AddLast(12)

	for i := 0; i < lista.Size(); i++ {
		fmt.Println(lista.Get(i))
	}

	fmt.Println("----------------------")

	lista.AddPos(4, 1)

	for i := 0; i < lista.Size(); i++ {
		fmt.Println(lista.Get(i))
	}

	fmt.Println("----------------------")

	lista.Remove(2)

	for i := 0; i < lista.Size(); i++ {
		fmt.Println(lista.Get(i))
	}

	fmt.Println("----------------------")

	lista.RemoveLast()

	for i := 0; i < lista.Size(); i++ {
		fmt.Println(lista.Get(i))
	}
}
