package main

import (
	"fmt"
)

func ComputerNumber(number int) string {
	if number > 9 && number < 20 {
		return fmt.Sprintf("%d компьютеров", number)
	} else {
		switch number % 10 {
		case 1:
			return fmt.Sprintf("%d компьютер", number)
		case 2, 3, 4:
			return fmt.Sprintf("%d компьютера", number)
		default:
			return fmt.Sprintf("%d компьютеров", number)
		}
	}
}

func main() {
	var number int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	result := ComputerNumber(number)
	fmt.Println(result)
}
