package processor

import (
	"os"

	"github.com/acmesquita/task_tracker/core/services"
	"github.com/acmesquita/task_tracker/infra/commands"
	"github.com/acmesquita/task_tracker/infra/processor/handlers"
)

type Processor struct {
	service services.TaskService
}

func NewProcessor(service services.TaskService) *Processor {
	return &Processor{service: service}
}

func (p *Processor) Process(command string, options map[string]string) {
	switch command {
	case "add":
		handlers.AddTask(p.service, options)
	case "update":
		handlers.UpdateTask(p.service, options)
	case "delete":
		handlers.DeleteTask(p.service, options)
	case "list":
		handlers.ListTasks(p.service)
	case "get":
		handlers.FindTask(p.service, options)
	case "mark-done":
		handlers.MarkDoneTask(p.service, options)
	case "help":
		handleHelpCommand()
	}
}

func handleHelpCommand() {
	commands.HandleHelperMessage()
	os.Exit(0)
}
