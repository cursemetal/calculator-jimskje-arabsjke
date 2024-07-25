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
		"I":      1,
		"II":     2,
		"III":    3,
		"IV":     4,
		"V":      5,
		"VI":     6,
		"VII":    7,
		"VIII":   8,
		"IX":     9,
		"X":      10,
		"XI":     11,
		"XII":    12,
		"XIII":   13,
		"XIV":    14,
		"XV":     15,
		"XVI":    16,
		"XVII":   17,
		"XVIII":  18,
		"XIX":    19,
		"XX":     20,
		"XXI":    21,
		"XXIV":   24,
		"XXV":    25,
		"XXVII":  27,
		"XXVIII": 28,
		"XXX":    30,
		"XXXII":  32,
		"XXXV":   35,
		"XXXVI":  36,
		"XL":     40,
		"XLII":   42,
		"XLV":    45,
		"XLVIII": 48,
		"XLIX":   49,
		"L":      50,
		"LIV":    54,
		"LVI":    56,
		"LX":     60,
		"LXIII":  63,
		"LXIV":   64,
		"LXX":    70,
		"LXXII":  72,
		"LXXX":   80,
		"LXXXI":  81,
		"XC":     90,
		"C":      100,
	}

	return convertrimsk[roman]
}

func rimskstr(num int) string {
	convertarabsk := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		24:  "XXIV",
		25:  "XXV",
		27:  "XXVII",
		28:  "XXVIII",
		30:  "XXX",
		32:  "XXXII",
		35:  "XXXV",
		36:  "XXXVI",
		40:  "XL",
		42:  "XLII",
		45:  "XLV",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		54:  "LIV",
		56:  "LVI",
		60:  "LX",
		63:  "LXIII",
		64:  "LXIV",
		70:  "LXX",
		72:  "LXXII",
		80:  "LXXX",
		81:  "LXXXI",
		90:  "XC",
		100: "C",
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
