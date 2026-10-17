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
		fmt.Println("Ошибка: N должно быть больше 0")
		return
	}

	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	avg := float64(sum) / float64(n)

	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее:", avg)
}
