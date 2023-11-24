package main

import "fmt"

type Tree interface {
	Add(value int)
	Search(value int) bool
	Min() int
	Max() int
	PrintPre()
	PrintIn()
	PrintPos()
	PrintLevels()
	Height() int
	Remove(value int) *BSTNode
}

type BSTNode struct {
	value int
	left  *BSTNode
	right *BSTNode
}

func (node *BSTNode) Add(value int) {
	if node == nil {
		*node = BSTNode{value: value, left: nil, right: nil}
		return
	}
	if value < node.value {
		if node.left == nil {
			node.left = &BSTNode{value: value, left: nil, right: nil}
		} else {
			node.left.Add(value)
		}
	} else if value > node.value {
		if node.right == nil {
			node.right = &BSTNode{value: value, left: nil, right: nil}
		} else {
			node.right.Add(value)
		}
	}
}

func (node *BSTNode) Search(value int) bool {
	if node == nil {
		return false
	}
	if value == node.value {
		return true
	} else if value < node.value {
		return node.left.Search(value)
	} else {
		return node.right.Search(value)
	}
}

func (node *BSTNode) Min() int {
	if node.left == nil {
		return node.value
	}
	return node.left.Min()
}

func (node *BSTNode) Max() int {
	if node.right == nil {
		return node.value
	}
	return node.right.Max()
}

func (node *BSTNode) PrintPre() {
	if node != nil {
		fmt.Print(node.value, " ")
		node.left.PrintPre()
		node.right.PrintPre()
	}
}

func (node *BSTNode) PrintIn() {
	if node != nil {
		node.left.PrintPre()
		fmt.Print(node.value, " ")
		node.right.PrintPre()
	}
}

func (node *BSTNode) PrintPos() {
	if node != nil {
		node.left.PrintPre()
		node.right.PrintPre()
		fmt.Print(node.value, " ")
	}
}

func (node *BSTNode) PrintLevels() {
	if node == nil {
		return
	}

	queue := []*BSTNode{node}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Print(current.value, " ")

		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}

func (node *BSTNode) Height() int {
	if node == nil {
		return 0
	}

	leftHeight := node.left.Height()
	rightHeight := node.right.Height()

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (node *BSTNode) Remove(val int) *BSTNode {
	if val < node.value && node.left != nil {
		node.left = node.left.Remove(val)
	} else if val > node.value && node.right != nil {
		node.right = node.right.Remove(val)
	} else {
		//encontramos o nó a ser removido
		//detectar se é caso 1, 2 ou 3
		if node.left == nil && node.right == nil { //caso1: remover nó folha
			//preciso retornar o endereço que vai ser armazenado no ponteiro do pai
			//que antes apontava para esse nó que queremos remover
			return nil
		} else if node.left != nil && node.right == nil { //caso2: remover nó com 1filho à esq
			//preciso retornar o endereço que vai ser armazenado no ponteiro
			//que antes apontava para esse nó que queremos remover
			return node.left
		} else if node.left == nil && node.right != nil { //caso2: remover nó com 1filho à dir
			//preciso retornar o endereço que vai ser armazenado no ponteiro
			//que antes apontava para esse nó que queremos remover
			return node.right
		} else { //caso3: remover nó com dois filhos
			//copiar para o nó a ser removido o valor do maior filho da sub-árvore esquerda ⇒ MAX
			node.value = node.left.Max()
			//remover o maior filho da sub-árvore esquerda
			node.left = node.left.Remove(node.value)

			//preciso retornar o endereço que vai ser armazenado no ponteiro
			//que antes apontava para esse nó que queremos remover
			return node
		}
	}
	return node
}
