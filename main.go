package main

import (
	"fmt"
	"secLab/stack"
	"secLab/set"
	"secLab/array"
)

func main() {
	expression := "3 + 5 * (2 - 8)"
	result, err := stack.Evaluate(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
		
	fmt.Println("\n")

	//--------------------------------------------------------------------------------------------------------

	// Создаем новое множество
	mySet := set.NewSet()

	// Добавляем элементы
	err1 := mySet.SETADD(10)
	if err1 != nil {
		fmt.Println("Ошибка при добавлении элемента:", err1)
	}
	err1 = mySet.SETADD(20)
	if err1 != nil {
		fmt.Println("Ошибка при добавлении элемента:", err1)
	}
	err1 = mySet.SETADD(10) // Попытка добавить дубликат
	if err1 != nil {
		fmt.Println("Ошибка при добавлении элемента:", err1)
	}

	// Проверяем наличие элементов
	fmt.Println("Содержит 10:", mySet.SET_AT(10)) // Ожидается: true
	fmt.Println("Содержит 20:", mySet.SET_AT(20)) // Ожидается: true
	fmt.Println("Содержит 30:", mySet.SET_AT(30)) // Ожидается: false

	// Удаляем элемент
	mySet.SETDEL(10)
	fmt.Println("Содержит 10 после удаления:", mySet.SET_AT(10)) // Ожидается: false

	// Печатаем текущее количество элементов в множестве
	fmt.Println("Текущее количество элементов в множестве:", mySet.Size())
	
	fmt.Println("\n")

	//---------------------------------------------------------------------------------------------------------

	// Создаем новое множество
	mySet2 := set.NewSet()

	// Добавляем элементы в множество
	mySet2.SETADD(4)
	mySet2.SETADD(10)
	mySet2.SETADD(5)
	mySet2.SETADD(1)
	mySet2.SETADD(3)
	mySet2.SETADD(7)

	// Попытка разбить множество на подмножества с равной суммой
	mySet2.Partition(0)

	// Пример с другим множеством
	fmt.Println("\nДругой пример:")
	mySet3 := set.NewSet()
	mySet3.SETADD(10)
	mySet3.SETADD(3)
	mySet3.SETADD(7)
	mySet3.SETADD(4)
	mySet3.SETADD(5)
	mySet3.SETADD(1)

	// Попытка разбить другое множество на подмножества с равной суммой
	mySet3.Partition(10)

	fmt.Println("\n")

	//---------------------------------------------------------------------------------------------------------

	// Исходный массив
	S := []string{"x", "y", "z"}

	// Генерируем и выводим подмножества
	array.GenerateSubsets(S)
}
