package bst

import "fmt"

// Tree представляет бинарное дерево поиска
type Tree struct {
	Root *Node // Корень дерева
}

// Insert вставляет новое значение в дерево
func (t *Tree) Insert(value int) {
	t.Root = insert(t.Root, value) // Вызываем вспомогательную функцию для вставки
}

// Вспомогательная функция для вставки узла
func insert(node *Node, value int) *Node {
	if node == nil {
		// Если текущий узел пустой, создаем новый узел
		return &Node{Value: value}
	}
	if value < node.Value {
		// Если значение меньше, рекурсивно вставляем в левое поддерево
		node.Left = insert(node.Left, value)
	} else {
		// Если значение больше или равно, вставляем в правое поддерево
		node.Right = insert(node.Right, value)
	}
	return node // Возвращаем текущий узел
}

// Delete удаляет узел с заданным значением
func (t *Tree) Delete(value int) {
	t.Root = deleteNode(t.Root, value) // Обновляем корень после удаления
}

// Вспомогательная функция для удаления узла
func deleteNode(node *Node, value int) *Node {
	if node == nil {
		// Если узел пустой, возвращаем nil
		return nil
	}

	if value < node.Value {
		// Если значение меньше, ищем в левом поддереве
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		// Если значение больше, ищем в правом поддереве
		node.Right = deleteNode(node.Right, value)
	} else {
		// Узел найден
		if node.Left == nil {
			// Если у узла нет левого потомка, возвращаем правого
			return node.Right
		} else if node.Right == nil {
			// Если у узла нет правого потомка, возвращаем левого
			return node.Left
		}

		// Узел с двумя потомками: находим минимальный узел в правом поддереве
		minNode := minValueNode(node.Right)
		node.Value = minNode.Value         // Заменяем значение удаляемого узла на минимальное
		node.Right = deleteNode(node.Right, minNode.Value) // Удаляем минимальный узел
	}
	return node // Возвращаем текущий узел
}

// minValueNode находит узел с минимальным значением
func minValueNode(node *Node) *Node {
	current := node
	// Ищем самый левый узел в правом поддереве
	for current.Left != nil {
		current = current.Left
	}
	return current // Возвращаем узел с минимальным значением
}

// InOrder выводит значения дерева в порядке возрастания
func (t *Tree) InOrder() {
	inOrder(t.Root) // Запускаем вспомогательную функцию для обхода дерева
	fmt.Println()   // Печатаем перевод строки
}

// Вспомогательная функция для обхода дерева в порядке возрастания
func inOrder(node *Node) {
	if node != nil {
		inOrder(node.Left)          // Сначала обходим левое поддерево
		fmt.Print(node.Value, " ")  // Затем печатаем значение узла
		inOrder(node.Right)         // В конце обходим правое поддерево
	}
}

// IsValidBST проверяет, является ли дерево корректным бинарным деревом поиска
func (t *Tree) IsValidBST() bool {
	return isValidBST(t.Root, nil, nil) // Запускаем проверку с корнем и без ограничений
}

// Вспомогательная функция для проверки корректности дерева
func isValidBST(node *Node, min *int, max *int) bool {
	if node == nil {
		return true // Пустое поддерево корректно
	}

	// Проверяем, что значение узла находится в заданных пределах
	if (min != nil && node.Value <= *min) || (max != nil && node.Value >= *max) {
		return false // Если значение узла выходит за границы, дерево некорректно
	}

	// Рекурсивно проверяем левое и правое поддеревья с обновленными границами
	return isValidBST(node.Left, min, &node.Value) && isValidBST(node.Right, &node.Value, max)
}
