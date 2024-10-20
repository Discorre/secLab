package set

import "errors"

// Set - структура для реализации множества
type Set struct {
	data [100]int // Массив фиксированной длины для хранения элементов
	size int      // Текущий размер множества
}

// NewSet создает новое множество
func NewSet() *Set {
	return &Set{
		size: 0,
	}
}

// SETADD добавляет элемент в множество
func (s *Set) SETADD(value int) error {
	if s.size >= len(s.data) {
		return errors.New("множество переполнено") // Проверка на переполнение
	}

	// Проверка, существует ли элемент уже в множестве
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			return nil // Элемент уже существует
		}
	}

	// Добавление нового элемента
	s.data[s.size] = value
	s.size++
	return nil
}

// SETDEL удаляет элемент из множества
func (s *Set) SETDEL(value int) {
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			// Сдвигаем все элементы влево, чтобы удалить элемент
			for j := i; j < s.size-1; j++ {
				s.data[j] = s.data[j+1]
			}
			s.size-- // Уменьшаем размер множества
			return
		}
	}
}

// SET_AT проверяет, содержится ли элемент в множестве
func (s *Set) SET_AT(value int) bool {
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			return true // Элемент найден
		}
	}
	return false // Элемент не найден
}

// Size возвращает текущее количество элементов в множестве
func (s *Set) Size() int {
	return s.size
}
