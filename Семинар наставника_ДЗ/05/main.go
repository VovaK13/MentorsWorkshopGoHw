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

	var x int
	fmt.Print("Введите число: ")
	fmt.Scan(&x)

	max := x
	min := x

	for i := 1; i < n; i++ {
		fmt.Print("Введите число: ")
		fmt.Scan(&x)
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)
}
