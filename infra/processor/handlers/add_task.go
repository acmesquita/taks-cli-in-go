package handlers

import (
	"fmt"

	"github.com/acmesquita/task_tracker/app/services"
)

func AddTask(service services.TaskService, options map[string]string) {
	fmt.Println("Adding task")
	task := service.AddTask(options["description"])
	fmt.Println(task.ID, task.Description, task.Status)
	fmt.Println("Task added successfully")
}
