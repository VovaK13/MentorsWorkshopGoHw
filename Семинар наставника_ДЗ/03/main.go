package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите число N: ")
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("N должно быть больше нуля")
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
