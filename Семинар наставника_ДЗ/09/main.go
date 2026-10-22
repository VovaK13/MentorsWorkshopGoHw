package main

import "fmt"

func main() {
	var balance float64
	var action int

	fmt.Print("Введите баланс: ")
	fmt.Scan(&balance)
	fmt.Println("1. Посмотреть баланс")
	fmt.Println("2. Пополнить баланс")
	fmt.Println("3. Снять средства")
	fmt.Print("Выберите действие: ")
	fmt.Scan(&action)

	switch action {
	case 1:
		fmt.Println("Баланс:", balance)
	case 2:
		var amount float64
		fmt.Print("Сумма пополнения: ")
		fmt.Scan(&amount)
		balance += amount
		fmt.Println("Баланс:", balance)
	case 3:
		var amount float64
		fmt.Print("Сумма снятия: ")
		fmt.Scan(&amount)
		if amount > balance {
			fmt.Println("Недостаточно средств")
		} else {
			balance -= amount
			fmt.Println("Баланс:", balance)
		}
	default:
		fmt.Println("Неизвестное действие")
	}
}
