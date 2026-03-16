package services

import (
	"testing"

	"github.com/acmesquita/task_tracker/app/model"
	"github.com/acmesquita/task_tracker/app/repository"
)

func setupTaskRepository() repository.TaskRepository {
	return &repository.TaskRepositoryImpl{
		Tasks: make([]*model.Task, 0),
	}
}

func TestAddTask(t *testing.T) {
	taskRepository := setupTaskRepository()
	taskService := NewTaskService(taskRepository)
	task := taskService.AddTask("Buy groceries")
	if task.ID == "" {
		t.Errorf("Task ID is empty")
	}

	if len(taskRepository.(*repository.TaskRepositoryImpl).Tasks) != 1 {
		t.Errorf("Task not added")
	}
}

func TestUpdateTask(t *testing.T) {
	taskRepository := setupTaskRepository()
	taskService := NewTaskService(taskRepository)
	task := taskService.AddTask("Buy groceries")
	taskService.UpdateTask(task.ID, "Buy groceries and cook dinner")
	if task.Description != "Buy groceries and cook dinner" {
		t.Errorf("Task description not updated")
	}
}

func TestDeleteTask(t *testing.T) {
	taskRepository := setupTaskRepository()
	taskService := NewTaskService(taskRepository)
	task := taskService.AddTask("Buy groceries")
	taskService.DeleteTask(task.ID)
	if len(taskRepository.(*repository.TaskRepositoryImpl).Tasks) != 0 {
		t.Errorf("Task not deleted")
	}
}

func TestListTasks(t *testing.T) {
	taskRepository := setupTaskRepository()
	taskService := NewTaskService(taskRepository)

	task1 := taskService.AddTask("Buy groceries")
	task2 := taskService.AddTask("Buy groceries and cook dinner")

	if task1.ID == "" || task2.ID == "" {
		t.Errorf("Tasks not added")
	}
	tasks := taskService.ListTasks()

	if len(tasks) != 2 {
		t.Errorf("Tasks not listed")
	}
	if tasks[0].ID != task1.ID {
		t.Errorf("Task 1 not listed")
	}
	if tasks[1].ID != task2.ID {
		t.Errorf("Task 2 not listed")
	}
}

func TestGetTask(t *testing.T) {
	taskRepository := setupTaskRepository()
	taskService := NewTaskService(taskRepository)
	task := taskService.AddTask("Buy groceries")
	taskFound := taskService.GetTask(task.ID)
	if taskFound == nil {
		t.Errorf("Task not found")
	}
	if taskFound.ID != task.ID {
		t.Errorf("Task ID not found")
	}
}

func TestMarkTaskAsDone(t *testing.T) {
	taskRepository := setupTaskRepository()
	taskService := NewTaskService(taskRepository)
	task := taskService.AddTask("Buy groceries")
	taskService.MarkTaskAsDone(task.ID)
	if task.Status != "done" {
		t.Errorf("Task not marked as done")
	}
}
