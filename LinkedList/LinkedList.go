package main

import "fmt"

type List[T any] interface {
	AddLast(value T)
	AddPos(value T, pos int)
	Update(value T, pos int)
	RemoveLast()
	Remove(pos int)
	Get(pos int) T
	InvertList()
	Size() int
	Display()
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{head: nil}
}

func (list *LinkedList[T]) Size() int {
	size := 0
	current := list.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func (list *LinkedList[T]) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func (list *LinkedList[T]) AddLast(value T) {
	newNode := Node[T]{value: value, next: nil}

	if list.head == nil {
		list.head = &newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = &newNode
}

func (list *LinkedList[T]) RemoveLast() {

	if list.head == nil {
		fmt.Println("A lista já está vazia")
		return
	}

	if list.head.next == nil {
		list.head = nil
		return
	}

	current := list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

func (list *LinkedList[T]) AddPos(value T, pos int) {
	current := list.head
	newNode := Node[T]{value: value, next: nil}

	if pos < 0 {
		fmt.Println("Posições negativas não são permitidas")
		return
	}

	for i := 0; i < pos-1; i++ {
		current = current.next
		if current == nil {
			fmt.Println("A posição está além do tamanho da lista.")
			return
		}
	}

	if current == nil {
		fmt.Println("A posição está além do tamanho da lista.")
		return
	}

	newNode.next = current.next
	current.next = &newNode
}

func (list *LinkedList[T]) Remove(pos int) {
	if pos < 0 {
		fmt.Println("Não é possível remover em uma posição menor do que zero")
		return
	}

	if pos == 0 {
		list.head = list.head.next
		return
	}

	current := list.head

	for i := 0; i < pos-1; i++ {
		if current == nil {
			fmt.Println("A posição está fora da lista")
			return
		}
		current = current.next
	}

	if current == nil {
		fmt.Println("A posição está fora da lista")
		return
	}

	removed_node := current.next
	current.next = removed_node.next
	removed_node.next = nil
}

func (list *LinkedList[T]) Get(pos int) T {
	if pos < 0 {
		fmt.Println("Não é possível obter uma posição menor do que zero")
	}
	current := list.head
	for i := 0; i < pos; i++ {
		if current == nil {
			fmt.Println("A posição está fora da lista")
			var None T
			return None
		}
		current = current.next
	}
	if current == nil {
		fmt.Println("A posição está fora da lista")
		var None T
		return None
	}
	return current.value
}

func (list *LinkedList[T]) Update(value T, pos int) {
	if pos < 0 {
		fmt.Println("Não é possível atualizar em uma posição menor do que zero")
		return
	}

	current := list.head
	for i := 0; i < pos; i++ {
		if current == nil {
			fmt.Println("A posição está fora da lista")
			return
		}
		current = current.next
	}

	if current == nil {
		fmt.Println("A posição está fora da lista")
		return
	}

	current.value = value
}

func (list *LinkedList[T]) InvertList() {
	current := list.head
	next := Node[T]{}
	previus := Node[T]{}

	for current != nil {
		next = current.next
	}
	list.head = current
}

func main() {
	list := NewLinkedList[int]()
	list.AddLast(1)
	list.AddLast(1)
	list.AddLast(10)
	list.AddPos(3, 1)
	list.AddLast(10)
	list.AddLast(15)
	list.Display()
	// fmt.Printf("Tamanho da lista: %d\n", list.Size())
	// fmt.Println("Removendo na posição 0")
	// list.Remove(0)
	// list.Display()
	// fmt.Printf("Tamanho da lista: %d\n", list.Size())
	list.InvertList()
	list.Display()
}
