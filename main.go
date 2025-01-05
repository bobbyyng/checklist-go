package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
	colorBlue  = "\033[34m"
)

type Task struct {
	ID          int
	Name        string
	IsCompleted bool
}

func PrintWarning(message string) {
	fmt.Println(colorRed + message + colorReset)
}

func PrintSuccess(message string) {
	fmt.Println(colorGreen + message + colorReset)
}

func PrintTitle(message string) {
	fmt.Println(colorBlue + message + colorReset)
}

var tasks []Task
var currentID int = 1

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		PrintTitle("\n=== To-Do List ===")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Exit")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addTask(reader)
		case "2":
			listTasks()
		case "3":
			markTaskAsCompleted(reader)
		case "4":
			PrintWarning("Exiting the To-Do List...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task name: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)
	taskName = strings.TrimSpace(taskName)

	newTask := Task{
		ID:          currentID,
		Name:        taskName,
		IsCompleted: false,
	}
	tasks = append(tasks, newTask)
	currentID++

	PrintSuccess("Task added successfully!")
}

func listTasks() {
	if len(tasks) == 0 {
		PrintWarning("No tasks found.")
		return
	}

	PrintTitle("=== Tasks ===")
	for _, task := range tasks {
		status := "Incomplete"
		if task.IsCompleted {
			status = "Completed"
		}
		fmt.Printf("%d. %s (%s)\n", task.ID, task.Name, status)
	}
}

func markTaskAsCompleted(reader *bufio.Reader) {
	listTasks()
	fmt.Print("Enter the ID of the task to mark as completed: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	taskID, err := strconv.Atoi(input)
	if err != nil {
		PrintWarning("Invalid task ID.")
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].IsCompleted = true
			PrintSuccess("Task marked as completed!")
			return
		}
	}

	PrintWarning("Task not found.")
}
