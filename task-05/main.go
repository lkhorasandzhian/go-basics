package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите количество чисел: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка ввода количества")
		return
	}

	if n <= 0 {
		fmt.Println("Кол-во должно быть положительным")
		return
	}

	var min, max int

	for i := 0; i < n; i++ {
		var number int

		fmt.Printf("Введите число %d: ", i+1)
		if _, err := fmt.Scan(&number); err != nil {
			fmt.Println("Ошибка ввода числа")
			return
		}

		if i == 0 {
			min = number
			max = number
			continue
		}

		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}

	fmt.Println("Минимум:", min)
	fmt.Println("Максимум:", max)
}
