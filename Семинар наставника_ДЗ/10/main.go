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

	sum := x
	max := x
	min := x
	even := 0
	odd := 0

	if x%2 == 0 {
		even++
	} else {
		odd++
	}

	for i := 1; i < n; i++ {
		fmt.Print("Введите число: ")
		fmt.Scan(&x)

		sum += x
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
		if x%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	avg := float64(sum) / float64(n)
	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее:", avg)
	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)
	fmt.Println("Чётных:", even)
	fmt.Println("Нечётных:", odd)
}
