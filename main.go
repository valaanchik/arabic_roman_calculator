package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var romToAr = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func operator(oper string, a, b int) int {
	var result int
	switch oper {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "/":
		result = a / b
	case "*":
		result = a * b
	}
	return result
}

func resultRome(resultInt int) string {
	var resultString string
	for resultInt > 0 {
		if resultInt >= 100 {
			resultString += "C"
			resultInt -= 100
		} else if resultInt >= 90 {
			resultString += "XC"
			resultInt -= 90
		} else if resultInt >= 50 {
			resultString += "L"
			resultInt -= 50
		} else if resultInt >= 40 {
			resultString += "XL"
			resultInt -= 40
		} else if resultInt >= 10 {
			resultString += "X"
			resultInt -= 10
		} else if resultInt >= 9 {
			resultString += "IX"
			resultInt -= 9
		} else if resultInt >= 5 {
			resultString += "V"
			resultInt -= 5
		} else if resultInt >= 4 {
			resultString += "IV"
			resultInt -= 4
		} else if resultInt >= 1 {
			resultString += "I"
			resultInt -= 1
		}
	}

	return resultString
}

func main() {
	var str string
	fmt.Scan(&str)
	str = strings.ReplaceAll(str, " ", "")

	re := regexp.MustCompile(`^(\d+|[IVX]+)([\+\-\*/])(\d+|[IVX]+)$`)
	part := re.FindStringSubmatch(str)

	if len(part) != 4 {
		if !strings.ContainsAny(str, "+-*/") {
			panic("Выдача паники, так как строка не является математической операцией.")
		} else {
			panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		}
	}

	if romToAr[part[1]] != 0 && romToAr[part[3]] != 0 {
		aInt := romToAr[part[1]]
		bInt := romToAr[part[3]]
		resultInt := operator(part[2], aInt, bInt)
		if resultInt < 1 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
		fmt.Println(resultRome(resultInt))
	}
	if romToAr[part[1]] != 0 && romToAr[part[3]] == 0 || romToAr[part[1]] == 0 && romToAr[part[3]] != 0 {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}
	if romToAr[part[1]] == 0 || romToAr[part[3]] == 0 {
		intA, _ := strconv.Atoi(part[1])
		intB, _ := strconv.Atoi(part[3])
		if intA <= 0 || intB <= 0 || intA >= 10 || intB >= 10 {
			panic("Выдача паники, так как входные данные не соответствуют требованиям")
		}
		fmt.Println(operator(part[2], intA, intB))
	}
}
