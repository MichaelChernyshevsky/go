package to_do

import (
	"bufio"
	"fmt"
	"strconv"
)

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("Нет задач.")
		return
	}

	fmt.Println("\nСписок задач:")
	for _, task := range tasks {
		status := " "
		if task.Complete {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Text)
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Введите текст задачи: ")
	scanner.Scan()
	text := scanner.Text()

	taskID++
	newTask := Task{ID: taskID, Text: text, Complete: false}
	tasks = append(tasks, newTask)

	fmt.Println("Задача добавлена.")
}

func completeTask(scanner *bufio.Scanner) {
	fmt.Print("Введите номер задачи для отметки как выполненной: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Неверный ввод. Введите число.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = true
			fmt.Println("Задача отмечена как выполненная.")
			return
		}
	}
	fmt.Println("Задача не найдена.")
}

func deleteTask(scanner *bufio.Scanner) {
	fmt.Print("Введите номер задачи для удаления: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Неверный ввод. Введите число.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Задача удалена.")
			return
		}
	}
	fmt.Println("Задача не найдена.")
}
