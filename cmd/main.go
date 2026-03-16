package main

import (
	"os"

	"github.com/acmesquita/task_tracker/app/repository"
	"github.com/acmesquita/task_tracker/app/services"
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
	processor.Process(command, options)
}
