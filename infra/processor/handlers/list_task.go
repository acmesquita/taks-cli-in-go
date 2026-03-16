package handlers

import (
	"fmt"

	"github.com/acmesquita/task_tracker/app/services"
)

func ListTasks(service services.TaskService) {
	fmt.Println("Listing tasks")
	tasks := service.ListTasks()
	for _, task := range tasks {
		fmt.Println(task.ID, task.Description, task.Status)
	}
	fmt.Println("Tasks listed successfully")
}
