package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) append(value int) {
	newNode := Node{value: value, next: nil}

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

func (list *LinkedList) display() {
	for list.head != nil {
		fmt.Printf("%d -> ", list.head.value)
		list.head = list.head.next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}
	list.append(2)
	list.append(1)
	list.append(10)
	list.display()
}
