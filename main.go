package main

import (
	"bufio"
	"fmt"
	"go_projects/to_do"
	"os"
)

func main() {
	Interface_main()
}

func Interface_main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n1.To-Do List")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			to_do.Interface_to_do()
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}
