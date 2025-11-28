package main

import (
	"fmt"
	"strconv"
)

func main() {
	var transactions []float64
	var input string

	fmt.Println("Введите транзакции (по одной). Чтобы закончить — введите 'n'")

	for {
		fmt.Print("Сумма: ")
		fmt.Scanf("%s\n", &input)

		// Проверяем выход
		if input == "n" || input == "N" {
			break
		}

		// Пробуем распарсить
		amount, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("  Ошибка ввода, попробуйте снова")
			continue
		}

		transactions = append(transactions, amount)
	}

	// Вывод транзакций
	fmt.Println("\n--- Транзакции ---")
	for i, t := range transactions {
		fmt.Printf("%d. %.2f\n", i+1, t)
	}

	// Подсчёт баланса
	var balance float64
	for _, t := range transactions {
		balance += t
	}
	fmt.Printf("\nБаланс: %.2f\n", balance)

	fmt.Print("\nНажмите Enter, чтобы выйти...")
	fmt.Scanln()

}
