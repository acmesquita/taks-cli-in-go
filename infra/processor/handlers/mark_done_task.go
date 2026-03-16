package handlers

import (
	"fmt"
	"os"

	"github.com/acmesquita/task_tracker/core/services"
	"github.com/acmesquita/task_tracker/infra/commands"
)

func MarkDoneTask(service services.TaskService, options map[string]string) {
	fmt.Println("Marking task as done")
	id := options["id"]
	if id == "" {
		fmt.Println("ID is required")
		commands.HandleHelperMessage()
		os.Exit(1)
	}
	task := service.MarkTaskAsDone(id)
	if task == nil {
		fmt.Println("Task not found")
		os.Exit(1)
	}
	fmt.Println(task.ID, task.Description, task.Status)
	fmt.Println("Task marked as done successfully")
}
