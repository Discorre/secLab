package hashmap

import "errors"

// Node представляет узел в хеш-таблице
type Node struct {
	Key   rune
	Value rune
	Next  *Node
}

// MyHashMap представляет хеш-таблицу
type MyHashMap struct {
	buckets [128]*Node // Массив из 128 "ведер" (достаточно для хранения символов ASCII)
}

// NewHashMap создает новую хеш-таблицу
func NewHashMap() *MyHashMap {
	return &MyHashMap{}
}

// Put добавляет соответствие ключа и значения в хеш-таблицу
func (h *MyHashMap) Put(key rune, value rune) error {
	index := key % 128 // Вычисляем индекс в массиве
	node := h.buckets[index]

	// Если в "ведре" нет узлов, создаем новый
	if node == nil {
		h.buckets[index] = &Node{Key: key, Value: value}
		return nil
	}

	// Ищем существующий узел или добавляем новый
	for node != nil {
		if node.Key == key {
			// Если ключ уже существует, обновляем значение
			node.Value = value
			return nil
		}
		if node.Next == nil {
			// Если дошли до конца списка, добавляем новый узел
			node.Next = &Node{Key: key, Value: value}
			return nil
		}
		node = node.Next
	}
	return nil
}

// Get получает значение по ключу из хеш-таблицы
func (h *MyHashMap) Get(key rune) (rune, error) {
	index := key % 128 // Вычисляем индекс в массиве
	node := h.buckets[index]

	for node != nil {
		if node.Key == key {
			return node.Value, nil // Если нашли, возвращаем значение
		}
		node = node.Next
	}
	return 0, errors.New("ключ не найден") // Если не нашли, возвращаем ошибку
}
