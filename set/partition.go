package set

import "fmt"

// Partition разбивает множество на непересекающиеся подмножества с заданной суммой
func (s *Set) Partition(target int) {
	// Проверка на возможность разбивки
	if target <= 0 {
		fmt.Println("Целевая сумма должна быть положительной.")
		return
	}

	used := make([]bool, s.size)
	found := false

	fmt.Printf("Успешно разбито на подмножества с суммой %d:\n", target)
	partitionHelper(s, target, 0, used, &found)

	if !found {
		fmt.Println("Невозможно разбить множество на подмножества с заданной суммой.")
	}
}

// partitionHelper рекурсивно находит подмножества с заданной суммой
func partitionHelper(s *Set, target, start int, used []bool, found *bool) {
	if target == 0 {
		// Если достигли целевой суммы, выводим подмножество
		fmt.Print("{")
		for i := 0; i < s.size; i++ {
			if used[i] {
				fmt.Print(s.data[i], " ")
			}
		}
		fmt.Println("}")
		*found = true
		return
	}

	for i := start; i < s.size; i++ {
		if !used[i] && s.data[i] <= target {
			used[i] = true
			partitionHelper(s, target-s.data[i], i+1, used, found)
			used[i] = false // Отмена выбора
		}
	}
}
