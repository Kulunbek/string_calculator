package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const maxStringLength = 10
const maxResultLength = 40

// Функция для форматирования длинного результата
func formatResult(result string) string {
	if len(result) > maxResultLength {
		return result[:maxResultLength] + "..."
	}
	return result
}

// Функция для выполнения арифметической операции
func calculateStringExpression(str1, op, str2 string) (string, error) {
	switch op {
	case "+":
		return formatResult(str1 + str2), nil
	case "-":
		return formatResult(strings.Replace(str1, str2, "", 1)), nil
	default:
		return "", fmt.Errorf("unsupported operation between strings: %s", op)
	}
}

// Функция для выполнения операции со строкой и числом
func calculateStringNumberExpression(str, op string, num int) (string, error) {
	if num < 1 || num > 10 {
		return "", fmt.Errorf("number must be between 1 and 10")
	}
	switch op {
	case "*":
		return formatResult(strings.Repeat(str, num)), nil
	case "/":
		if len(str) < num {
			return "", fmt.Errorf("cannot divide string length by larger number")
		}
		return formatResult(str[:len(str)/num]), nil
	default:
		return "", fmt.Errorf("unsupported operation between string and number: %s", op)
	}
}

// Главная функция разбора и обработки выражения
func processExpression(input string) (string, error) {
	// Используем константу maxStringLength в регулярных выражениях
	stringOpRegex := regexp.MustCompile(fmt.Sprintf(`^"(.{1,%d})"\s*([+\-])\s*"(.{1,%d})"$`, maxStringLength, maxStringLength))
	stringNumOpRegex := regexp.MustCompile(fmt.Sprintf(`^"(.{1,%d})"\s*([*/])\s*([1-9]|10)$`, maxStringLength))

	// Проверка строки-выражения
	if matches := stringOpRegex.FindStringSubmatch(input); matches != nil {
		str1, op, str2 := matches[1], matches[2], matches[3]
		return calculateStringExpression(str1, op, str2)
	}

	// Проверка операции строки и числа
	if matches := stringNumOpRegex.FindStringSubmatch(input); matches != nil {
		str, op, numStr := matches[1], matches[2], matches[3]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return "", fmt.Errorf("invalid number")
		}
		return calculateStringNumberExpression(str, op, num)
	}

	// Если выражение не соответствует шаблону
	return "", fmt.Errorf("invalid expression")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строковое выражение:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Обработка выражения
	result, err := processExpression(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}

	// Вывод результата
	fmt.Println(result)
}
