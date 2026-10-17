package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите целое число: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка ввода числа")
		return
	}

	result := "Нет"
	if isPrime(n) {
		result = "Да"
	}

	fmt.Println("Число простое?", result)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
