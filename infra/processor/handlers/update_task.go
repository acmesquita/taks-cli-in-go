package handlers

import (
	"fmt"
	"os"

	"github.com/acmesquita/task_tracker/app/services"
	"github.com/acmesquita/task_tracker/infra/commands"
)

func UpdateTask(service services.TaskService, options map[string]string) {
	fmt.Println("Updating task")
	description := options["description"]
	if description == "" {
		description = options["d"]
	}
	if description == "" {
		fmt.Println("Description is required")
		commands.HandleHelperMessage()
		os.Exit(1)
	}
	task := service.UpdateTask(options["id"], description)
	if task == nil {
		fmt.Println("Task not found")
		os.Exit(1)
	}
	fmt.Println(task.ID, task.Description, task.Status)
	fmt.Println("Task updated successfully")
}
