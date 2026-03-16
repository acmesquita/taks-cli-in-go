package handlers

import (
	"fmt"

	"github.com/acmesquita/task_tracker/core/services"
	"github.com/acmesquita/task_tracker/infra/processor/adapter"
)

func ListTasks(service services.TaskService, request adapter.Request) {
	fmt.Println("Listing tasks")
	status := request.GetOptions()["status"]
	if status == "" {
		status = "all"
	}

	tasks := service.ListTasksByStatus(status)
	for _, task := range tasks {
		fmt.Println(task.ID, task.Description, task.Status)
	}
	fmt.Println("Tasks listed successfully")
}
