package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите целое число: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка ввода числа")
		return
	}

	if n > 0 {
		fmt.Println("Положительное число")
	} else if n < 0 {
		fmt.Println("Отрицательное число")
	} else {
		fmt.Println("Ноль")
	}

	if n%2 == 0 {
		fmt.Println("Чётное число")
	} else {
		fmt.Println("Нечётное число")
	}

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("Делится на 3 и 5")
	} else {
		fmt.Println("Не делится на 3 и 5")
	}
}
