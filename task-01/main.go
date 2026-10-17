package main

import "fmt"

func main() {
	var a, b float64
	var operation string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите операцию: ")
	fmt.Scan(&operation)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	switch operation {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}
