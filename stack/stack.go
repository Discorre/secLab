package stack

import "fmt"

// Узел связного списка для чисел
type Node struct {
	Value int
	Next  *Node
}

// Структура стека для чисел
type Stack struct {
	top *Node
}

// Push добавляет элемент в стек
func (s *Stack) Push(val int) {
	newNode := &Node{
		Value: val,
		Next:  s.top, // Новый элемент указывает на текущий верхний элемент
	}
	s.top = newNode // Новый элемент становится верхним
}

// Pop удаляет верхний элемент и возвращает его значение
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("стек пуст")
	}
	val := s.top.Value // Берем значение верхнего элемента
	s.top = s.top.Next // Перемещаем верхний элемент на следующий
	return val, nil
}

// Top возвращает значение верхнего элемента без его удаления
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("стек пуст")
	}
	return s.top.Value, nil
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Узел для стека операторов
type OperatorNode struct {
	Operator rune
	Next     *OperatorNode
}

// Стек для операторов
type OperatorStack struct {
	top *OperatorNode
}

// Push добавляет оператор в стек
func (s *OperatorStack) Push(op rune) {
	newNode := &OperatorNode{
		Operator: op,
		Next:     s.top,
	}
	s.top = newNode
}

// Pop удаляет верхний элемент и возвращает его
func (s *OperatorStack) Pop() (rune, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("стек пуст")
	}
	op := s.top.Operator
	s.top = s.top.Next
	return op, nil
}

// Top возвращает верхний элемент без удаления
func (s *OperatorStack) Top() (rune, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("стек пуст")
	}
	return s.top.Operator, nil
}

// IsEmpty проверяет, пуст ли стек
func (s *OperatorStack) IsEmpty() bool {
	return s.top == nil
}
