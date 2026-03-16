package processor_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/acmesquita/task_tracker/core/repository"
	"github.com/acmesquita/task_tracker/core/services"
	"github.com/acmesquita/task_tracker/infra/processor"
)

func newJSONBackedProcessor(t *testing.T) (*processor.Processor, *repository.JSONTaskRepository, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "tasks-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	repo := &repository.JSONTaskRepository{
		FilePath: tmpFile.Name(),
	}

	service := services.NewTaskService(repo)
	p := processor.NewProcessor(*service)

	cleanup := func() {
		_ = os.Remove(tmpFile.Name())
	}

	return p, repo, cleanup
}

func TestProcessor_AddAndListTasks_WithJSONRepository(t *testing.T) {
	p, repo, cleanup := newJSONBackedProcessor(t)
	defer cleanup()

	p.Process("add", map[string]string{"description": "Buy groceries"})
	p.Process("add", map[string]string{"description": "Learn Go"})

	tasks := repo.ListTasks()

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].Description != "Buy groceries" {
		t.Errorf("unexpected first task description: %s", tasks[0].Description)
	}

	if tasks[1].Description != "Learn Go" {
		t.Errorf("unexpected second task description: %s", tasks[1].Description)
	}
}

func TestProcessor_MarkDoneAndGetTask_WithJSONRepository(t *testing.T) {
	p, repo, cleanup := newJSONBackedProcessor(t)
	defer cleanup()

	p.Process("add", map[string]string{"description": "Finish report"})

	tasks := repo.ListTasks()
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task after add, got %d", len(tasks))
	}

	id := tasks[0].ID

	p.Process("mark-done", map[string]string{"id": id})

	task := repo.GetTask(id)
	if task == nil {
		t.Fatalf("expected task with id %s to exist", id)
	}

	if task.Status != "done" {
		t.Errorf("expected task status to be done, got %s", task.Status)
	}
}

func TestProcessor_MarkInProgressAndGetTask_WithJSONRepository(t *testing.T) {
	p, repo, cleanup := newJSONBackedProcessor(t)
	defer cleanup()

	p.Process("add", map[string]string{"description": "Finish report"})

	tasks := repo.ListTasks()
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task after add, got %d", len(tasks))
	}

	id := tasks[0].ID

	p.Process("mark-in-progress", map[string]string{"id": id})

	task := repo.GetTask(id)
	if task == nil {
		t.Fatalf("expected task with id %s to exist", id)
	}

	if task.Status != "in_progress" {
		t.Errorf("expected task status to be in progress, got %s", task.Status)
	}
}

func TestProcessor_DeleteTask_WithJSONRepository(t *testing.T) {
	p, repo, cleanup := newJSONBackedProcessor(t)
	defer cleanup()

	p.Process("add", map[string]string{"description": "Finish report"})

	tasks := repo.ListTasks()
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task after add, got %d", len(tasks))
	}

	id := tasks[0].ID

	p.Process("delete", map[string]string{"id": id})

	tasks = repo.ListTasks()
	if len(tasks) != 0 {
		t.Fatalf("expected 0 tasks after delete, got %d", len(tasks))
	}

	task := repo.GetTask(id)
	if task != nil {
		t.Errorf("expected task with id %s to be deleted", id)
	}
}

func TestProcessor_UpdateTask_WithJSONRepository(t *testing.T) {
	p, repo, cleanup := newJSONBackedProcessor(t)
	defer cleanup()

	p.Process("add", map[string]string{"description": "Finish report"})

	tasks := repo.ListTasks()
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task after add, got %d", len(tasks))
	}

	id := tasks[0].ID

	p.Process("update", map[string]string{"id": id, "description": "Finish report v2"})

	task := repo.GetTask(id)
	if task == nil {
		t.Fatalf("expected task with id %s to exist", id)
	}

	if task.Description != "Finish report v2" {
		t.Errorf("expected task description to be Finish report v2, got %s", task.Description)
	}
}

func TestProcessor_FindTask_WithJSONRepository(t *testing.T) {
	p, repo, cleanup := newJSONBackedProcessor(t)
	defer cleanup()

	p.Process("add", map[string]string{"description": "Finish report"})

	tasks := repo.ListTasks()
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task after add, got %d", len(tasks))
	}

	id := tasks[0].ID

	task := repo.GetTask(id)
	if task == nil {
		t.Fatalf("expected task with id %s to exist", id)
	}

	if task.Description != "Finish report" {
		t.Errorf("expected task description to be Finish report, got %s", task.Description)
	}

	if task.Status != "todo" {
		t.Errorf("expected task status to be todo, got %s", task.Status)
	}
}

func TestProcessor_HelpCommand_WithJSONRepository(t *testing.T) {
	p, _, cleanup := newJSONBackedProcessor(t)
	defer cleanup()

	err := func() error {
		defer func() error {
			if r := recover(); r != nil {
				return fmt.Errorf("panic: %v", r)
			}
			return nil
		}()
		p.Process("help", map[string]string{})
		return nil
	}()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestProcessor_InvalidCommand_WithJSONRepository(t *testing.T) {
	p, _, cleanup := newJSONBackedProcessor(t)
	defer cleanup()

	err := func() error {
		defer func() error {
			if r := recover(); r != nil {
				return fmt.Errorf("panic: %v", r)
			}
			return nil
		}()
		p.Process("invalid", map[string]string{})
		return nil
	}()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
