package main

import "fmt"

func main() {

	var num1, num2 int
	var operator string

	fmt.Print("Ввести первую цифру:")
	fmt.Scanln(&num1)

	fmt.Print("Ввести действие (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Ввести вторую цифру:")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("Результат: %d\n", num1+num2)
	case "-":
		fmt.Printf("Результат: %d\n", num1-num2)
	case "*":
		fmt.Printf("Результат: %d\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Результат: %d\n", num1/num2)
		} else {
			fmt.Println("Ошибка: деление на ноль.")
		}
	default:
		fmt.Println("Неправильное действие")
	}
}