package main

import (
	"bufio"
	"checklist-go/internal/task"
	"checklist-go/internal/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	taskManager := task.NewTaskManager()

	for {
		utils.PrintTitle("\n=== To-Do List ===")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Exit")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			taskManager.AddTask(reader)
		case "2":
			taskManager.ListTasks()
		case "3":
			taskManager.MarkTaskAsCompleted(reader)
		case "4":
			utils.PrintWarning("Exiting the To-Do List...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
