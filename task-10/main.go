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
		fmt.Println("Количество чисел должно быть положительным")
		return
	}

	sum := 0
	evenCount := 0
	oddCount := 0

	var min, max int

	for i := 0; i < n; i++ {
		var number int

		fmt.Printf("Введите число %d: ", i+1)
		if _, err := fmt.Scan(&number); err != nil {
			fmt.Println("Ошибка ввода числа")
			return
		}

		sum += number

		if i == 0 {
			min = number
			max = number
		} else {
			if number < min {
				min = number
			}
			if number > max {
				max = number
			}
		}

		if number%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	avg := float64(sum) / float64(n)

	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее:", avg)
	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)
	fmt.Println("Количество чётных:", evenCount)
	fmt.Println("Количество нечётных:", oddCount)
}
