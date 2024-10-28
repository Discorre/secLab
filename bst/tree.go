package bst

import "fmt"

// Node представляет узел двоичного дерева
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Tree представляет двоичное дерево поиска
type Tree struct {
	root *Node
}

// NewNode создает новый узел
func NewNode(value int) *Node {
	return &Node{
		data:  value,
		left:  nil,
		right: nil,
	}
}

// Insert добавляет элемент в дерево
func (t *Tree) Insert(value int) {
	t.root = t.insert(t.root, value)
}

func (t *Tree) insert(node *Node, value int) *Node {
	if node == nil {
		return NewNode(value)
	}
	if value < node.data {
		node.left = t.insert(node.left, value)
	} else {
		node.right = t.insert(node.right, value)
	}
	return node
}

// Delete удаляет элемент из дерева
func (t *Tree) Delete(value int) {
	t.root = t.deleteNode(t.root, value)
}

func (t *Tree) deleteNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.data {
		node.left = t.deleteNode(node.left, value)
	} else if value > node.data {
		node.right = t.deleteNode(node.right, value)
	} else {
		// Найден узел для удаления
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		// Найдем минимальный узел в правом поддереве
		minNode := t.findMin(node.right)
		node.data = minNode.data // Скопируем значение минимального узла
		node.right = t.deleteNode(node.right, minNode.data) // Удаляем минимальный узел
	}
	return node
}

// findMin находит узел с минимальным значением
func (t *Tree) findMin(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

// InOrder выполняет in-order обход дерева и выводит значения
func (t *Tree) InOrder() {
	t.inOrder(t.root)
	fmt.Println()
}

func (t *Tree) inOrder(node *Node) {
	if node != nil {
		t.inOrder(node.left)
		fmt.Print(node.data, " ")
		t.inOrder(node.right)
	}
}

// IsValidBST проверяет, является ли дерево корректным бинарным деревом поиска
func (t *Tree) IsValidBST() bool {
	return t.isValidBST(t.root, nil, nil)
}

func (t *Tree) isValidBST(node *Node, min, max *int) bool {
	if node == nil {
		return true
	}
	if (min != nil && node.data <= *min) || (max != nil && node.data >= *max) {
		return false
	}
	return t.isValidBST(node.left, min, &node.data) && t.isValidBST(node.right, &node.data, max)
}
