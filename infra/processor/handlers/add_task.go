package handlers

import (
	"fmt"
	"os"

	"github.com/acmesquita/task_tracker/core/services"
	"github.com/acmesquita/task_tracker/infra/commands"
	"github.com/acmesquita/task_tracker/infra/processor/adapter"
)

func AddTask(service services.TaskService, request adapter.Request) {
	fmt.Println("Adding task")
	description := request.GetOptions()["description"]
	if description == "" {
		description = request.GetOptions()["d"]
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
