package main

import (
	"fmt"
	"os"

	"github.com/acmesquita/task_tracker/core/repository"
	"github.com/acmesquita/task_tracker/core/services"
	"github.com/acmesquita/task_tracker/infra/commands"
	"github.com/acmesquita/task_tracker/infra/processor"
)

func main() {
	command, args := commands.ParseCommands(os.Args[1:])
	options := commands.ParseOptions(command, args)

	service := services.NewTaskService(&repository.JSONTaskRepository{
		FilePath: "tasks.json",
	})

	processor := processor.NewProcessor(*service)

	fmt.Println("Processing command", command, options)
	processor.Process(command, options)
	fmt.Println("Command processed")
}
