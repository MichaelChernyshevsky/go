package to_do

import (
	"bufio"
	"fmt"
	"os"
)

func Interface_to_do() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nTo-Do List")
		fmt.Println("1. Показать все задачи")
		fmt.Println("2. Добавить задачу")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Выйти")
		fmt.Print("Выберите действие: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			showTasks()
		case "2":
			addTask(scanner)
		case "3":
			completeTask(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}
