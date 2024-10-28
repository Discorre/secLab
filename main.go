package main

import (
	"bufio"
	"fmt"
	"os"
	"secLab/array"
	"secLab/bst"
	"secLab/hashmap"
	"secLab/set"
	"secLab/stack"
	"strconv"
	"strings"
)

func main() {
	// Создаем новый сканер для чтения ввода из консоли
	reader := bufio.NewReader(os.Stdin)

	for {
		// Приглашение для ввода
		fmt.Print("Введите номер задания (1-6) или exit для выхода из проограммы: ")

		// Считываем ввод пользователя
		input, _ := reader.ReadString('\n')

		// Убираем символы новой строки
		input = strings.TrimSpace(input)

		// Используем switch для обработки ввода
		switch input {
		case "1":
			fmt.Print("Введите выражение (+, -, *): ")
			expression, _ := reader.ReadString('\n') //"3 + 5 * (2 - 8)"
			result, err := stack.Evaluate(expression)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Результат:", result)
			}
		case "2":
			exit := false

			// Создаем новое множество
			mySet := set.NewSet()
		
			// Создаем новый сканер для чтения ввода из консоли
			reader := bufio.NewReader(os.Stdin)
		
			for !exit {
				// Приглашение для ввода
				fmt.Print("Введите команду (add, contains, delete, size, save, load, back): ")
		
				// Считываем ввод пользователя
				input, _ := reader.ReadString('\n')
		
				// Убираем символы новой строки
				input = strings.TrimSpace(input)
		
				// Разбиваем ввод на команду и аргументы
				parts := strings.Split(input, " ")
				command := parts[0]
		
				switch command {
				case "add":
					if len(parts) < 2 {
						fmt.Println("Недостаточно аргументов для команды add")
						continue
					}
					element, err := strconv.Atoi(parts[1])
					if err != nil {
						fmt.Println("Ошибка при преобразовании аргумента:", err)
						continue
					}
					mySet.ADD(element)
					//fmt.Println("Элемент добавлен")
		
				case "contains":
					if len(parts) < 2 {
						fmt.Println("Недостаточно аргументов для команды contains")
						continue
					}
					element, err := strconv.Atoi(parts[1])
					if err != nil {
						fmt.Println("Ошибка при преобразовании аргумента:", err)
						continue
					}
					fmt.Println("Содержит элемент:", mySet.CONTAINS(element))
		
				case "delete":
					if len(parts) < 2 {
						fmt.Println("Недостаточно аргументов для команды delete")
						continue
					}
					element, err := strconv.Atoi(parts[1])
					if err != nil {
						fmt.Println("Ошибка при преобразовании аргумента:", err)
						continue
					}
					mySet.DELETE(element)
					fmt.Println("Элемент удален")
		
				case "size":
					fmt.Println("Текущее количество элементов в множестве:", mySet.Size())
		
				case "save":
					if len(parts) < 2 {
						fmt.Println("Недостаточно аргументов для команды save")
						continue
					}
					filename := parts[1]
					err := mySet.SaveToFile(filename)
					if err != nil {
						fmt.Println("Ошибка при сохранении в файл:", err)
					} else {
						fmt.Println("Множество сохранено в файл:", filename)
					}
		
				case "load":
					if len(parts) < 2 {
						fmt.Println("Недостаточно аргументов для команды load")
						continue
					}
					filename := parts[1]
					err := mySet.LoadFromFile(filename)
					if err != nil {
						fmt.Println("Ошибка при загрузке из файла:", err)
					} else {
						fmt.Println("Множество загружено из файла:", filename)
					}
		
				case "back":
					fmt.Println("Назад...")
					exit = true
		
				default:
					fmt.Println("Неизвестная команда")
				}
			}
		case "3":
			exit := false
		mySet := set.NewSet()
		reader := bufio.NewReader(os.Stdin)

		for !exit {
			fmt.Print("Введите команду (add, contains, delete, size, save, load, partit, back): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			parts := strings.Split(input, " ")
			command := parts[0]

			switch command {
			case "add":
				if len(parts) < 2 {
					fmt.Println("Недостаточно аргументов для команды add")
					continue
				}
				element, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Ошибка при преобразовании аргумента:", err)
					continue
				}
				err = mySet.ADD(element)
				if err != nil {
					fmt.Println("Ошибка при добавлении элемента:", err)
				} else {
					fmt.Println("Элемент добавлен")
				}
			case "contains":
				if len(parts) < 2 {
					fmt.Println("Недостаточно аргументов для команды contains")
					continue
				}
				element, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Ошибка при преобразовании аргумента:", err)
					continue
				}
				fmt.Println("Содержит элемент:", mySet.CONTAINS(element))
			case "delete":
				if len(parts) < 2 {
					fmt.Println("Недостаточно аргументов для команды delete")
					continue
				}
				element, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Ошибка при преобразовании аргумента:", err)
					continue
				}
				mySet.DELETE(element)
				fmt.Println("Элемент удален")
			case "size":
				fmt.Println("Текущее количество элементов в множестве:", mySet.Size())
			case "save":
				if len(parts) < 2 {
					fmt.Println("Недостаточно аргументов для команды save")
					continue
				}
				filename := parts[1]
				err := mySet.SaveToFile(filename)
				if err != nil {
					fmt.Println("Ошибка при сохранении в файл:", err)
				} else {
					fmt.Println("Множество сохранено в файл:", filename)
				}
			case "load":
				if len(parts) < 2 {
					fmt.Println("Недостаточно аргументов для команды load")
					continue
				}
				filename := parts[1]
				err := mySet.LoadFromFile(filename)
				if err != nil {
					fmt.Println("Ошибка при загрузке из файла:", err)
				} else {
					fmt.Println("Множество загружено из файла:", filename)
				}
			case "partit":
				if len(parts) < 2 {
					fmt.Println("Недостаточно аргументов для команды partit")
					continue
				}
				element, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Ошибка при преобразовании аргумента:", err)
					continue
				}
				mySet.Partition(element)
			case "back":
				fmt.Println("Назад...")
				exit = true
			default:
				fmt.Println("Неизвестная команда")
			}
		}
		case "4":
		// Чтение строки с элементами массива от пользователя
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Введите элементы массива через пробел:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Разбиваем строку на элементы массива
		S := strings.Split(input, " ")

		// Создаем MyArray и добавляем элементы
		arr := array.NewMyArray[string](len(S))
		for _, item := range S {
			arr.MPUSH(item)
		}

		// Генерируем и выводим подмассивы
		arr.PrintSubarrays()
		case "5":
			tree := &bst.Tree{}

			// Ввод значений для массива values
			fmt.Println("Введите значения для вставки в дерево (через пробел):")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// Разбиваем строку на отдельные значения
			valuesStr := strings.Split(input, " ")
			values := []int{}

			// Преобразуем строки в числа и добавляем их в массив values
			for _, v := range valuesStr {
				value, err := strconv.Atoi(v)
				if err == nil {
					values = append(values, value)
				} else {
					fmt.Printf("Некорректное значение: %s\n", v)
				}
			}

			// Вставка значений в дерево
			for _, v := range values {
				tree.Insert(v)
			}

			fmt.Printf("Дерево корректно: %v\n", tree.IsValidBST())

			// Ввод значения для удаления
			fmt.Println("Введите значение для удаления из дерева:")
			input1, _ := reader.ReadString('\n')
			value, err := strconv.Atoi(strings.TrimSpace(input1))
			if err == nil {
				tree.Delete(value)
			} else {
				fmt.Println("Некорректное значение для удаления.")
			}

			// Вывод дерева в порядке возрастания
			fmt.Println("Дерево в порядке возрастания:")
			tree.InOrder()
			fmt.Printf("Дерево корректно: %v\n", tree.IsValidBST())
		case "6":
			reader := bufio.NewReader(os.Stdin)

			// Ввод первой пары строк
			fmt.Println("Введите первую строку (str1):")
			str1, _ := reader.ReadString('\n')
			str1 = str1[:len(str1)-1] // Убираем символ новой строки

			fmt.Println("Введите вторую строку (str2):")
			str2, _ := reader.ReadString('\n')
			str2 = str2[:len(str2)-1] // Убираем символ новой строки

			// Проверка первой пары строк
			fmt.Printf(hashmap.IsIsomorphic(str1, str2))
		case "exit":
			fmt.Println("Остановка программы...")
			return
		}

	}
}
