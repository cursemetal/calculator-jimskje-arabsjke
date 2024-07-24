package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rimskint(roman string) int {
	convertrimsk := map[string]int{
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
	return convertrimsk[roman]
}

func rimskstr(num int) string {
	convertarabsk := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}
	return convertarabsk[num]
}
func main() {
	typer := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	expression, _ := typer.ReadString('\n')

	// Удалем пробелы
	expression = strings.TrimSpace(expression)

	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		fmt.Println("Неверный ввод")
		return
	}

	isRoman := false
	_, err := strconv.Atoi(parts[0])
	if err != nil {
		isRoman = true
	}

	_, err = strconv.Atoi(parts[2])
	if err != nil && !isRoman {
		fmt.Println("Неверный ввод")
		return
	}

	var num1, num2 int
	if isRoman {
		num1 = rimskint(parts[0])
		num2 = rimskint(parts[2])
	} else {
		num1, err = strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Неверное число")
			return
		}
		num2, err = strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Неверное число")
			return
		}
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		fmt.Println("Введите числа от 1 до 10")
		return
	}
	var result int
	switch parts[1] {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		// Целочисленное деление
		result = num1 / num2
	default:
		fmt.Println("Неподдерживаемая операция")
		return
	}

	if isRoman {
		// Проверка на <= 0
		if result <= 0 {
			fmt.Println("Ответ должен быть больше нуля")
			return
		}
		fmt.Println(rimskstr(result))
	} else {
		fmt.Println(result)
	}
}
