package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// the Task structure
type Task struct {
	ID   int
	Name string
}

var tasks []Task
var nextID int

func displayMenu() {
	fmt.Println("\n--- To-Do List CLI ---")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Remove Task")
	fmt.Println("4. Exit")
	fmt.Print("Choose an option: ")
}

func addTask() {
	fmt.Print("Enter task name: ")
	reader := bufio.NewReader(os.Stdin)
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	if taskName != "" {
		tasks = append(tasks, Task{ID: nextID, Name: taskName})
		nextID++
		fmt.Println("Task added successfully!")
	} else {
		fmt.Println("Task name cannot be empty!")
	}
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("\n--- Tasks ---")
	for _, task := range tasks {
		fmt.Printf("%d. %s\n", task.ID+1, task.Name)
	}
}

func removeTask() {
	fmt.Print("Enter task ID to remove: ")
	var taskID int
	_, err := fmt.Scanf("%d", &taskID)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid task ID.")
		return
	}

	var indexToRemove = -1
	for i, task := range tasks {
		if task.ID == taskID {
			indexToRemove = i
			break
		}
	}

	if indexToRemove != -1 {
		tasks = append(tasks[:indexToRemove], tasks[indexToRemove+1:]...)
		fmt.Println("Task removed successfully!")
	} else {
		fmt.Println("Task ID not found!")
	}
}

func main() {
	for {
		displayMenu()

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid choice, please try again.")
			bufio.NewReader(os.Stdin).ReadString('\n')
			continue
		}

		switch choice {
		case 1:
			addTask()
		case 2:
			listTasks()
		case 3:
			removeTask()
		case 4:
			fmt.Println("Exiting the To-Do List CLI.")
			return
		default:
			fmt.Println("Invalid option! Please choose again.")
		}
	}
}
