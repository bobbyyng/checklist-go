package task

import (
	"bufio"
	"checklist-go/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

type Task struct {
	ID          int
	Name        string
	IsCompleted bool
}

type TaskManager struct {
	tasks     []Task
	currentID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:     []Task{},
		currentID: 1,
	}
}

func (tm *TaskManager) AddTask(reader *bufio.Reader) {
	fmt.Print("Enter task name: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	newTask := Task{
		ID:          tm.currentID,
		Name:        taskName,
		IsCompleted: false,
	}
	tm.tasks = append(tm.tasks, newTask)
	tm.currentID++

	utils.PrintSuccess("Task added successfully!")
}

func (tm *TaskManager) ListTasks() {
	if len(tm.tasks) == 0 {
		utils.PrintWarning("No tasks found.")
		return
	}

	utils.PrintTitle("=== Tasks ===")
	for _, task := range tm.tasks {
		status := "Incomplete"
		if task.IsCompleted {
			status = "Completed"
		}
		fmt.Printf("%d. %s (%s)\n", task.ID, task.Name, status)
	}
}

func (tm *TaskManager) MarkTaskAsCompleted(reader *bufio.Reader) {
	tm.ListTasks()
	fmt.Print("Enter the ID of the task to mark as completed: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	taskID, err := strconv.Atoi(input)
	if err != nil {
		utils.PrintWarning("Invalid task ID.")
		return
	}

	for i, task := range tm.tasks {
		if task.ID == taskID {
			tm.tasks[i].IsCompleted = true
			utils.PrintSuccess("Task marked as completed!")
			return
		}
	}

	utils.PrintWarning("Task not found.")
}
