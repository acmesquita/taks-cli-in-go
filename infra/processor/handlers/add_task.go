package handlers

import (
	"fmt"
	"os"

	"github.com/acmesquita/task_tracker/core/services"
	"github.com/acmesquita/task_tracker/infra/commands"
)

func AddTask(service services.TaskService, options map[string]string) {
	fmt.Println("Adding task")
	description := options["description"]
	if description == "" {
		description = options["d"]
	}
	if description == "" {
		fmt.Println("Description is required")
		commands.HandleHelperMessage()
		os.Exit(1)
	}
	task := service.AddTask(description)
	fmt.Println(task.ID, task.Description, task.Status)
	fmt.Println("Task added successfully")
}
