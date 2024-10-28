package stack

import (
	"fmt"
	"unicode"
)

// Функция для выполнения арифметических операций
func applyOperation(a, b int, op rune) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	default:
		panic("неизвестная операция")
	}
}

// Функция для определения приоритета операций
func precedence(op rune) int {
	if op == '+' || op == '-' {
		return 1
	}
	if op == '*' {
		return 2
	}
	return 0
}

// Функция для вычисления выражения
func Evaluate(expression string) (int, error) {
	var values Stack         // Стек для чисел
	var operators OperatorStack // Стек для операторов

	// Вспомогательная функция для выполнения операции на верхушке стека
	applyTopOperation := func() {
		if values.IsEmpty() || operators.IsEmpty() {
			return
		}
		b, _ := values.Pop()
		a, _ := values.Pop()
		op, _ := operators.Pop()
		values.Push(applyOperation(a, b, op))
	}

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])

		if unicode.IsSpace(ch) {
			continue // Пропускаем пробелы
		}

		// Если символ — цифра, читаем всё число
		if unicode.IsDigit(ch) {
			num := 0
			for i < len(expression) && unicode.IsDigit(rune(expression[i])) {
				num = num*10 + int(expression[i]-'0')
				i++
			}
			i-- // Сдвигаем индекс назад, т.к. цикл увеличил его на 1 лишний раз
			values.Push(num)
		} else if ch == '(' {
			operators.Push(ch)
		} else if ch == ')' {
			for !operators.IsEmpty() && func() bool { op, _ := operators.Top(); return op != '(' }() {
				applyTopOperation()
			}
			operators.Pop() // Убираем '('
		} else if ch == '+' || ch == '-' || ch == '*' {
			for !operators.IsEmpty() && precedence(func() rune { op, _ := operators.Top(); return op }()) >= precedence(ch) {
				applyTopOperation()
			}
			operators.Push(ch)
		} else {
			return 0, fmt.Errorf("недопустимый символ: %c", ch)
		}
	}

	// Выполняем оставшиеся операции
	for !operators.IsEmpty() {
		applyTopOperation()
	}

	if result, _ := values.Pop(); values.IsEmpty() {
		return result, nil
	} else {
		return 0, fmt.Errorf("ошибка при вычислении выражения")
	}
}
