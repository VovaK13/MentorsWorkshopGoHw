package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("Количество должно быть больше нуля")
		return
	}

	positive := 0
	negative := 0
	zeros := 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Print("Введите число: ")
		fmt.Scan(&x)

		if x > 0 {
			positive++
		} else if x < 0 {
			negative++
		} else {
			zeros++
		}
	}

	fmt.Println("Положительных:", positive)
	fmt.Println("Отрицательных:", negative)
	fmt.Println("Нулей:", zeros)
}
