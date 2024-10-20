package array

import "fmt"

// GenerateSubsets генерирует все подмножества массива
func GenerateSubsets(arr []string) {
	// Получаем все подмножества
	subsets := make([][]string, 0)
	n := len(arr)
	// Общее количество подмножеств = 2^n
	totalSubsets := 1 << n // 2^n

	for i := 0; i < totalSubsets; i++ {
		subset := []string{}
		for j := 0; j < n; j++ {
			// Проверяем, включен ли элемент в текущее подмножество
			if (i & (1 << j)) != 0 {
				subset = append(subset, arr[j])
			}
		}
		subsets = append(subsets, subset)
	}

	// Вывод всех подмножеств
	fmt.Println("Все различные подмножества:")
	for _, s := range subsets {
		fmt.Println(s)
	}
}
