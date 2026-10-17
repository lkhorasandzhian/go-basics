package main

import "fmt"

func main() {
	var a, b float64
	var operation string

	fmt.Print("Введите первое число: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("Ошибка ввода первого числа")
		return
	}

	fmt.Print("Введите операцию: ")
	if _, err := fmt.Scan(&operation); err != nil {
		fmt.Println("Ошибка ввода операции")
		return
	}

	fmt.Print("Введите второе число: ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Ошибка ввода второго числа")
		return
	}

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
		fmt.Println("Неизвестная операция. Необходимо ввести операцию '+', '-', '*', '/'")
	}
}
