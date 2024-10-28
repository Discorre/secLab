package set

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Set - структура для реализации множества
type Set struct {
	data map[int]struct{} // Хэш-таблица для хранения элементов
	size int              // Текущий размер множества
}

// NewSet создает новое множество
func NewSet() *Set {
	return &Set{
		data: make(map[int]struct{}), // Инициализация хэш-таблицы
		size: 0,
	}
}

// ADD добавляет элемент в множество
func (s *Set) ADD(value int) error {
	if _, exists := s.data[value]; exists {
		fmt.Println("Элемент уже есть в множестве")
		return nil // Элемент уже существует
	}

	s.data[value] = struct{}{} // Добавление нового элемента
	s.size++
	return nil
}

// DELETE удаляет элемент из множества
func (s *Set) DELETE(value int) {
	if _, exists := s.data[value]; exists {
		delete(s.data, value) // Удаление элемента
		s.size--              // Уменьшаем размер множества
	}
}

// CONTAINS проверяет, содержится ли элемент в множестве
func (s *Set) CONTAINS(value int) bool {
	_, exists := s.data[value]
	return exists // Элемент найден или не найден
}

// Size возвращает текущее количество элементов в множестве
func (s *Set) Size() int {
	return s.size
}

// SaveToFile сохраняет множество в файл
func (s *Set) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for value := range s.data {
		_, err := file.WriteString(fmt.Sprintf("%d\n", value))
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadFromFile загружает множество из файла
func (s *Set) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		err = s.ADD(value) // Добавляем элемент в множество
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}