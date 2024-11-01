package set

import (
	"fmt"
	"strconv"
	"os"
)

// Тип Set на основе uint32, где каждый бит обозначает наличие числа от 0 до 31
type Set uint32

// Добавление элемента в Set
func (s *Set) Add(element int) {
	*s |= (1 << element) // Сдвигаем бит влево на `element` позиций и делаем побитовое ИЛИ с текущим Set
}

// Проверка, содержится ли элемент в Set
func (s *Set) Contains(element int) bool {
	return *s&(1<<element) != 0 // Сдвигаем бит влево и проверяем с помощью побитового И, установлен ли соответствующий бит
}

// Удаление элемента из Set
func (s *Set) Remove(element int) {
	*s &^= (1 << element) // Сдвигаем бит влево и делаем побитовое И НЕ, чтобы сбросить бит
}

// Подсчет суммы элементов множества с помощью Set
func (s *Set) Sum() int {
	sum := 0
	for i := 0; i < 32; i++ { // Перебираем все 32 возможных бита
		if s.Contains(i) {    // Если бит установлен, то число i присутствует в Set
			sum += i
		}
	}
	return sum // Возвращаем сумму всех элементов
}

// Подсчет количества элементов в Set
func (s *Set) Count() int {
	count := 0
	for i := 0; i < 32; i++ { // Перебираем все возможные элементы
		if s.Contains(i) { // Если элемент присутствует в множестве
			count++ // Увеличиваем счетчик
		}
	}
	return count // Возвращаем количество элементов
}

// Запись множества в файл
func (s *Set) WriteToFile(filename string) error {
	file, err := os.Create(filename) // Создаем файл с указанным именем
	if err != nil {
		return err
	}
	defer file.Close() // Закрываем файл после завершения

	for i := 0; i < 32; i++ { // Перебираем все возможные элементы
		if s.Contains(i) { // Если элемент присутствует в множестве
			_, err := file.WriteString(strconv.Itoa(i) + "\n") // Записываем элемент в файл
			if err != nil {
				return err
			}
		}
	}
	return nil // Успешное завершение
}

// Чтение множества из файла
func (s *Set) ReadFromFile(filename string) error {
	file, err := os.Open(filename) // Открываем файл для чтения
	if err != nil {
		return err
	}
	defer file.Close() // Закрываем файл после завершения

	var element int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &element) // Считываем каждый элемент из файла
		if err != nil {
			break // Выходим из цикла при ошибке (например, конец файла)
		}
		s.Add(element) // Добавляем элемент в множество
	}
	return nil // Успешное завершение
}

// Функция для разбиения множества на подмножества
func Partition(set Set, target int) (Set, error) {
	var result Set  // Хранение подмножеств, каждое из которых представлено типом Set
	used := Set(0)  // Set для отслеживания использованных элементов
	var backtrack func(Set, int, int) bool

	backtrack = func(currentSet Set, startElement, currentSum int) bool {
		if currentSum == target {
			// Если сумма текущего подмножества равна целевой, добавляем его в результат
			result |= currentSet  // Добавляем подмножество в результат
			// Добавляем элементы из currentSet в used, чтобы избежать их повторного использования
			used |= currentSet
			return true
		}

		// Перебор элементов для формирования подмножества
		for i := startElement; i < 32; i++ {
			if !set.Contains(i) || used.Contains(i) || currentSum+i > target {
				continue
			}

			currentSet.Add(i)
			if backtrack(currentSet, i+1, currentSum+i) {
				return true
			}
			currentSet.Remove(i)
		}
		return false
	}

	// Повторно вызываем backtrack, чтобы найти все подмножества, пока это возможно
	for {
		currentSet := Set(0)
		if !backtrack(currentSet, 1, 0) {  // Стартуем с 1, так как множество натуральных чисел
			break
		}
	}

	// Проверка, что сумма всех подмножеств равна общей сумме
	if used.Sum() != set.Sum() {
		return 0, fmt.Errorf("невозможно разбить на подмножества с равной суммой")
	}
	return result, nil
}