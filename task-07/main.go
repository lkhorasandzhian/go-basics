package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите N: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка ввода числа")
		return
	}

	if n <= 0 {
		fmt.Println("N должно быть положительным")
		return
	}

	sum := 0

	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}

	fmt.Println("Сумма:", sum)
}
