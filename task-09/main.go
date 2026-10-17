package main

import "fmt"

func main() {
	var balance float64
	var action int

	fmt.Print("Введите начальный баланс: ")
	if _, err := fmt.Scan(&balance); err != nil {
		fmt.Println("Ошибка ввода баланса")
		return
	}

	if balance < 0 {
		fmt.Println("Баланс должен быть неотрицательным")
		return
	}

	fmt.Println("1. Посмотреть баланс")
	fmt.Println("2. Пополнить баланс")
	fmt.Println("3. Снять средства")

	fmt.Print("Ваш выбор: ")
	if _, err := fmt.Scan(&action); err != nil {
		fmt.Println("Ошибка ввода действия")
		return
	}

	var value float64
	switch action {
	case 1:
		fmt.Println("Текущий баланс:", balance)
	case 2:
		fmt.Print("Введите сумму пополнения: ")
		if _, err := fmt.Scan(&value); err != nil {
			fmt.Println("Ошибка ввода суммы")
			return
		}

		if value <= 0 {
			fmt.Println("Сумма должна быть положительной")
			return
		}

		balance += value
		fmt.Println("Текущий баланс:", balance)
	case 3:
		fmt.Print("Введите сумму снятия: ")
		if _, err := fmt.Scan(&value); err != nil {
			fmt.Println("Ошибка ввода суммы")
			return
		}

		if value <= 0 {
			fmt.Println("Сумма должна быть положительной")
			return
		}

		if value > balance {
			fmt.Println("Сумма снятия не должна превышать баланс")
			return
		}

		balance -= value
		fmt.Println("Текущий баланс:", balance)
	default:
		fmt.Println("Неизвестное действие")
	}
}
