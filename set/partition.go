package set

import (
	"fmt"
)

// Partition разбивает множество на непересекающиеся подмножества с заданной суммой
func (s *Set) Partition(target int) {
	// Проверка на возможность разбивки
	if target <= 0 {
		fmt.Println("Целевая сумма должна быть положительной.")
		return
	}

	// Преобразуем ключи хэшмапа в срез для обработки
	elements := make([]int, 0, len(s.data))
	for key := range s.data {
		elements = append(elements, key)
	}

	used := make([]bool, len(elements))
	found := false

	fmt.Printf("Успешно разбито на подмножества с суммой %d:\n", target)
	partitionHelper(elements, target, 0, used, &found)

	if !found {
		fmt.Println("Невозможно разбить множество на подмножества с заданной суммой.")
	}
}

// partitionHelper рекурсивно находит подмножества с заданной суммой
func partitionHelper(elements []int, target, start int, used []bool, found *bool) {
	if target == 0 {
		// Если достигли целевой суммы, выводим подмножество
		fmt.Print("{")
		for i, val := range elements {
			if used[i] {
				fmt.Print(val, " ")
			}
		}
		fmt.Println("}")
		*found = true
		return
	}

	for i := start; i < len(elements); i++ {
		if !used[i] && elements[i] <= target {
			used[i] = true
			partitionHelper(elements, target-elements[i], i+1, used, found)
			used[i] = false // Отмена выбора
		}
	}
}
