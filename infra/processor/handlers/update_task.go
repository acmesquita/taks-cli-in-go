package handlers

import (
	"fmt"
	"os"

	"github.com/acmesquita/task_tracker/core/services"
	"github.com/acmesquita/task_tracker/infra/commands"
	"github.com/acmesquita/task_tracker/infra/processor/adapter"
)

func UpdateTask(service services.TaskService, request adapter.Request) {

	fmt.Println("Updating task")
	description := request.GetOptions()["description"]
	if description == "" {
		description = request.GetOptions()["d"]
	}
	if description == "" {
		fmt.Println("Description is required")
		commands.HandleHelperMessage()
		os.Exit(1)
	}
	task := service.UpdateTask(request.GetOptions()["id"], description)
	if task == nil {
		fmt.Println("Task not found")
		os.Exit(1)
	}
	fmt.Println(task.ID, task.Description, task.Status)
	fmt.Println("Task updated successfully")
}
