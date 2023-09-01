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
	Display()
}

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type DoublyLinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{head: nil, tail: nil, length: 0}
}

func (list *DoublyLinkedList[T]) AddLast(value T) {
	newNode := Node[T]{value: value, prev: nil, next: nil}
	if list.length == 0 {
		list.head = &newNode
		list.tail = &newNode
		list.length++
		return
	}
	newNode.prev = list.tail
	newNode.next = nil
	list.tail.next = &newNode
	list.tail = &newNode
	list.length++
}

func (list *DoublyLinkedList[T]) RemoveLast() {
	if list.length == 0 {
		fmt.Println("Não há elemento para remover")
		return
	}
	if list.length == 1 {
		list.head = nil
		list.tail = nil
		list.length--
		return
	}
	prevNode := list.tail.prev
	prevNode.next = nil
	list.tail = prevNode
	list.length--
}

func (list *DoublyLinkedList[T]) AddPos(value T, pos int) {
	if pos < 0 || pos > list.length-1 {
		fmt.Println("Posição inválida")
		return
	}
	newNode := Node[T]{value: value, prev: nil, next: nil}
	current := list.head
	if pos == 0 {
		newNode.next = list.head
		newNode.prev = nil
		list.head = &newNode
		list.length++
		return
	}
	if pos == list.length-1 {
		list.AddLast(value)
		return
	}
	for i := 0; i < pos-1; i++ {
		current = current.next
	}
	newNode.next = current.next
	newNode.prev = current
	current.next = &newNode
	list.length++
}

func (list *DoublyLinkedList[T]) Remove(pos int) {
	if pos < 0 || pos > list.length-1 {
		fmt.Println("Posição inválida")
		return
	}
	if pos == 0 {
		list.head = list.head.next
		list.length--
		return
	}
	if pos == list.length-1 {
		prevNode := list.tail.prev
		prevNode.next = nil
		list.tail = prevNode
		list.length--
		return
	}
	current := list.head
	for i := 0; i < pos-1; i++ {
		current = current.next
	}
	removedNode := current.next
	removedNodeNext := removedNode.next
	current.next = removedNode.next
	removedNodeNext.prev = current
	removedNode.next = nil
	removedNode.prev = nil
	list.length--
}

func (list *DoublyLinkedList[T]) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%v <-> ", current.value)
		current = current.next
	}
	fmt.Printf("nil\n")
}

func (list *DoublyLinkedList[T]) Update(value T, pos int) {
	current := list.head
	if pos < 0 || pos > list.length-1 {
		fmt.Println("Posição inválida")
		return
	}
	for i := 0; i < pos; i++ {
		current = current.next
	}
	current.value = value
}

func (list *DoublyLinkedList[T]) Get(pos int) T {
	if pos < 0 || pos > list.length-1 {
		fmt.Println("Posição inválida")
		var None T
		return None
	}
	current := list.head
	for i := 0; i < pos; i++ {
		current = current.next
	}
	return current.value
}

func (list *DoublyLinkedList[T]) Size() int {
	return list.length
}

func main() {
	list := NewDoublyLinkedList[int]()
	list.AddLast(0)
	list.AddLast(4)
	list.AddLast(10)
	list.Display()
	list.AddPos(5, 0)
	list.Display()
	list.Update(7, 3)
	list.Display()
	list.Remove(3)
	list.Display()
	list.AddPos(86, 2)
	list.Display()
}
