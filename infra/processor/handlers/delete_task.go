package handlers

import (
	"fmt"
	"os"

	"github.com/acmesquita/task_tracker/core/services"
	"github.com/acmesquita/task_tracker/infra/commands"
	"github.com/acmesquita/task_tracker/infra/processor/adapter"
)

func DeleteTask(service services.TaskService, request adapter.Request) {
	fmt.Println("Deleting task")
	id := request.GetOptions()["id"]
	if id == "" {
		fmt.Println("ID is required")
		commands.HandleHelperMessage()
		os.Exit(1)
	}
	task := service.DeleteTask(id)
	if task == nil {
		fmt.Println("Task not found")
		os.Exit(1)
	}
	fmt.Println(task.ID, task.Description, task.Status)
	fmt.Println("Task deleted successfully")
}
