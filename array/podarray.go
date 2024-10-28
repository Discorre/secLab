package array

import (
	"fmt"
	//"myarray" // Импортируем пакет с MyArray
)

// Subarrays выводит все различные подмассивы массива
func (arr *MyArray[T]) Subarrays() [][]T {
	subarrays := [][]T{}

	// Общее количество подмножеств для массива длины `arr.length` = 2^arr.length
	numSubarrays := 1 << arr.length // 2^length

	for i := 0; i < numSubarrays; i++ {
		var subset []T
		for j := 0; j < arr.length; j++ {
			// Проверяем, включен ли j-й элемент в текущий подмассив
			if i&(1<<j) != 0 {
				subset = append(subset, arr.data[j])
			}
		}
		subarrays = append(subarrays, subset)
	}

	return subarrays
}

// PrintSubarrays выводит все подмассивы в удобном формате
func (arr *MyArray[T]) PrintSubarrays() {
	subarrays := arr.Subarrays()
	fmt.Println("Все различные подмассивы:")
	for _, subset := range subarrays {
		fmt.Println(subset)
	}
}
