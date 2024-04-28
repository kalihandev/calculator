package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Римские цифры к арабским и наоборот
var romanToArabic = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

func romanToArabicFunc(roman string) int {
	romanSimbol := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	if _, ok := romanSimbol[roman]; ok { // сопоставление римских цифр с арабскими
		return romanSimbol[roman]
	} else {
		panic("формат математической операции не удовлетворяет заданию — " +
			"операнды от 1 до 10 включительно.")
	}
}

func arabicToRomanFunc(arabic int) string {
	// Определение соответствий между римскими цифрами и их значением
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string
	for i := 0; i < len(values); i++ {
		// Пока число больше или равно значению в списке values
		for arabic >= values[i] {
			result += numerals[i] // Добавить соответствующую римскую цифру в результат
			arabic -= values[i]   // Уменьшить число на значение римской цифры
		}
	}
	return result
}

func isRomanNumber(s string) bool {
	_, ok := romanToArabic[s]
	return ok
}

func parseAndCalculate(expression string) (string, error) {
	// Удаление пробелов и проверка строки
	expression = strings.ReplaceAll(expression, " ", "")
	re := regexp.MustCompile(`^([IVXLCDM]+|\d+)([-+*/])([IVXLCDM]+|\d+)$`) //нифига не работает
	matches := re.FindStringSubmatch(expression)

	operand1, operator, operand2 := matches[1], matches[2], matches[3]
	operand1 = strings.TrimSpace(operand1)
	operand2 = strings.TrimSpace(operand2)
	operator = strings.TrimSpace(operator)
	// Проверяем, какие у нас операнды, и выполняем вычисления
	isRoman1 := isRomanNumber(operand1)
	isRoman2 := isRomanNumber(operand2)

	if isRoman1 != isRoman2 {
		return "", fmt.Errorf("Операнды должны быть из одной системы счисления")
	}

	var num1, num2 int
	if isRoman1 {
		num1, num2 = romanToArabicFunc(operand1), romanToArabicFunc(operand2)
	} else {
		var err error
		num1, err = strconv.Atoi(operand1)
		if err != nil {
			return "", err
		}
		num2, err = strconv.Atoi(operand2)
		if err != nil {
			return "", err
		}
	}

	result := 0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		return "", fmt.Errorf("Неправильный оператор")
	}

	if result < 1 {
		return "", fmt.Errorf("Результат больше 10")
	}

	// Выводим результат в нужном формате
	if isRoman1 {
		return arabicToRomanFunc(result), nil
	} else {
		return strconv.Itoa(result), nil
	}
}

func main() {
	for {
		fmt.Println("Введите выражение (например, III + II или 3 + 2):")
		var input string
		fmt.Scanln(&input)

		result, err := parseAndCalculate(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}
