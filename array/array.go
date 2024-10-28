	package array

	import (
		"errors"
		"fmt"
		"os"
	)

	// MyArray - структура для хранения динамического массива
	type MyArray[T any] struct {
		data     []T
		capacity int
		length   int
	}

	// NewMyArray - конструктор для создания массива с начальной емкостью
	func NewMyArray[T any](size int) *MyArray[T] {
		if size < 1 {
			size = 1
		}
		return &MyArray[T]{
			data:     make([]T, 0, size),
			capacity: size,
			length:   0,
		}
	}

	// MPUSH добавляет элемент в конец массива
	func (arr *MyArray[T]) MPUSH(element T) {
		if arr.length == arr.capacity {
			arr.resize()
		}
		arr.data = append(arr.data, element)
		arr.length++
	}

	// MPUSHIndex добавляет элемент по указанному индексу
	func (arr *MyArray[T]) MPUSHIndex(index int, element T) error {
		if index < 0 || index > arr.length {
			return errors.New("index out of range")
		}
		if arr.length == arr.capacity {
			arr.resize()
		}
		arr.data = append(arr.data[:index], append([]T{element}, arr.data[index:]...)...)
		arr.length++
		return nil
	}

	// MDEL удаляет элемент по указанному индексу
	func (arr *MyArray[T]) MDEL(index int) error {
		if index < 0 || index >= arr.length {
			return errors.New("index out of range")
		}
		arr.data = append(arr.data[:index], arr.data[index+1:]...)
		arr.length--
		return nil
	}

	// MGET возвращает элемент по индексу
	func (arr *MyArray[T]) MGET(index int) (T, error) {
		if index < 0 || index >= arr.length {
			var zero T
			return zero, errors.New("index out of range")
		}
		return arr.data[index], nil
	}

	// MRESET заменяет элемент по индексу
	func (arr *MyArray[T]) MRESET(index int, element T) error {
		if index < 0 || index >= arr.length {
			return errors.New("index out of range")
		}
		arr.data[index] = element
		return nil
	}

	// resize увеличивает емкость массива в два раза
	func (arr *MyArray[T]) resize() {
		newCapacity := arr.capacity * 2
		newData := make([]T, arr.length, newCapacity)
		copy(newData, arr.data)
		arr.data = newData
		arr.capacity = newCapacity
	}

	// Length возвращает текущую длину массива
	func (arr *MyArray[T]) Length() int {
		return arr.length
	}

	// SaveToFile сохраняет массив в файл
	func (arr *MyArray[T]) SaveToFile(filename string) error {
		file, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("unable to open file for writing: %w", err)
		}
		defer file.Close()

		for _, value := range arr.data[:arr.length] {
			_, err := fmt.Fprintln(file, value)
			if err != nil {
				return fmt.Errorf("error writing to file: %w", err)
			}
		}
		return nil
	}

	// LoadFromFile загружает массив из файла
	func (arr *MyArray[T]) LoadFromFile(filename string) error {
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("unable to open file for reading: %w", err)
		}
		defer file.Close()

		arr.Clear() // Очистка массива перед загрузкой
		var value T
		for {
			_, err := fmt.Fscanln(file, &value)
			if err != nil {
				break
			}
			arr.MPUSH(value)
		}
		return nil
	}

	// Clear очищает массив
	func (arr *MyArray[T]) Clear() {
		arr.data = make([]T, 0, arr.capacity)
		arr.length = 0
	}

	// Print выводит все элементы массива
	func (arr *MyArray[T]) Print() {
		for i := 0; i < arr.length; i++ {
			fmt.Printf("%v ", arr.data[i])
		}
		fmt.Println()
	}
