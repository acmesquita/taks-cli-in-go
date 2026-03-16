package processor

import (
	"os"

	"github.com/acmesquita/task_tracker/core/services"
	"github.com/acmesquita/task_tracker/infra/commands"
	"github.com/acmesquita/task_tracker/infra/processor/adapter"
	"github.com/acmesquita/task_tracker/infra/processor/handlers"
)

type Processor struct {
	service services.TaskService
}

func NewProcessor(service services.TaskService) *Processor {
	return &Processor{service: service}
}

func (p *Processor) Process(command string, options map[string]string) {

	request := adapter.NewRequest(command, options)

	switch command {
	case "add":
		handlers.AddTask(p.service, *request)
	case "update":
		handlers.UpdateTask(p.service, *request)
	case "delete":
		handlers.DeleteTask(p.service, *request)
	case "list":
		handlers.ListTasks(p.service)
	case "get":
		handlers.FindTask(p.service, *request)
	case "mark-done":
		handlers.MarkDoneTask(p.service, *request)
	case "mark-in-progress":
		handlers.MarkInProgressTask(p.service, *request)
	case "help":
		handleHelpCommand()
	}
}

func handleHelpCommand() {
	commands.HandleHelperMessage()
	os.Exit(0)
}
